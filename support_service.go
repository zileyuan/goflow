package goflow

import (
	"fmt"
	"strings"
	"time"

	"github.com/lunny/log"
)

func GetActiveTasksByOrderId(orderId string) []*Task {
	task := &Task{}
	tasks, _ := task.GetActiveTasksByOrderId(orderId)
	return tasks
}

func GetTaskActors(taskModel *TaskModel, execution *Execution) []string {
	assignee := taskModel.Assignee
	if assignee != "" {
		assigneeInf := execution.Args[taskModel.Assignee]
		if assigneeInf == nil {
			assigneeInf = taskModel.Assignee
		}
		switch assigneeInf.(type) {
		case string:
			return strings.Split(assigneeInf.(string), ",")
		case []string:
			return assigneeInf.([]string)
		case int:
			return []string{IntToStr(assigneeInf.(int))}
		default:
		}
	}
	return nil
}

//创建task，并根据model类型决定是否分配参与者
func CreateTask(taskModel *TaskModel, execution *Execution) []*Task {
	actors := GetTaskActors(taskModel, execution)
	args := execution.Args
	args[DEFAULT_KEY_ACTOR] = actors

	task := &Task{
		Id:          execution.Order.Id,
		TaskName:    taskModel.Name,
		DisplayName: taskModel.DisplayName,
		CreateTime:  time.Now(),
		TaskType:    taskModel.TaskType,
		Model:       taskModel,
		ExpireTime:  taskModel.ExpireTime,
		Variable:    MapToJson(args),
	}
	if execution.Task == nil {
		task.ParentTaskId = DEFAULT_START_ID
	} else {
		task.ParentTaskId = execution.Task.Id
	}

	action := args[taskModel.Action]
	if action == nil {
		task.Action = taskModel.Action
	} else {
		task.Action = action.(string)
	}

	tasks := make([]*Task, 0)
	actors = GetTaskActors(taskModel, execution)

	if taskModel.PerformType == PT_ANY {
		SaveTask(task, actors...)
		tasks = append(tasks, task)
	} else {
		for _, actor := range actors {
			singleTask := *task
			pSingleTask := &singleTask
			SaveTask(pSingleTask, actor)
			tasks = append(tasks, pSingleTask)
		}
	}
	return tasks
}

func SaveTask(task *Task, actors ...string) {
	task.Id = NewUUID()
	task.PerformType = PT_ANY
	Save(task, task.Id)
	AssignTask(task.Id, actors...)
}

//根据已有任务、任务类型、参与者创建新的任务，适用于转派，动态协办处理
func (p *TaskService) CreateNewTask(taskId string, taskType TASK_TYPE, actors ...string) {
	task := &Task{}
	task.GetTaskById(taskId)
	newTask := *task
	pNewTask := &newTask
	pNewTask.TaskType = taskType
	pNewTask.CreateTime = time.Now()
	pNewTask.ParentTaskId = taskId
	SaveTask(pNewTask, actors...)
}

func RejectTask(processModel *ProcessModel, currTask *Task) *Task {
	parentTaskId := currTask.ParentTaskId
	if parentTaskId == "" || parentTaskId == DEFAULT_START_ID {
		return nil
	}
	currentNode := processModel.GetNode(currTask.TaskName)
	historyTask := &HistoryTask{}
	historyTask.GetHistoryTaskById(parentTaskId)
	parentNode := processModel.GetNode(historyTask.TaskName)
	if CanRejected(currentNode, parentNode) {
		task := historyTask.Undo()
		task.Id = NewUUID()
		task.CreateTime = time.Now()
		Save(task, task.Id)
		AssignTask(task.Id, task.Operator)
		return task
	}
	return nil
}

func WithdrawTask(taskId string, operator string) *Task {
	historyTask := &HistoryTask{}
	historyTask.GetHistoryTaskById(taskId)
	var tasks []*Task
	if historyTask.PerformType == PT_ANY {
		tasks, _ = GetNextAnyActiveTasks(historyTask.Id)
	} else {
		tasks, _ = GetNextAllActiveTasks(historyTask.OrderId, historyTask.TaskName, historyTask.ParentTaskId)
	}
	for _, task := range tasks {
		DeleteById(task, task.Id)
	}

	task := historyTask.Undo()
	task.Id = NewUUID()
	task.CreateTime = time.Now()
	Save(task, task.Id)
	AssignTask(task.Id, task.Operator)
	return task
}

func AddTaskActor(taskId string, performType PERFORM_TYPE, actors ...string) {
	task := &Task{}
	task.GetTaskById(taskId)
	if performType == PT_ANY {
		AssignTask(taskId, actors...)
		v := JsonToMap(task.Variable)
		oldActor := v[DEFAULT_KEY_ACTOR].(string)
		v[DEFAULT_KEY_ACTOR] = oldActor + "," + strings.Join(actors, ",")
		task.Variable = MapToJson(v)
		Update(task, task.Id)
	} else {
		for _, actor := range actors {
			newTask := *task
			pNewTask := &newTask
			pNewTask.Id = NewUUID()
			pNewTask.CreateTime = time.Now()
			pNewTask.Operator = actor
			v := JsonToMap(task.Variable)
			v[DEFAULT_KEY_ACTOR] = actor
			task.Variable = MapToJson(v)
			Save(pNewTask, pNewTask.Id)
			AssignTask(pNewTask.Id, actor)
		}
	}
}

func Take(taskId string, operator string) *Task {
	task := &Task{}
	success, err := task.GetTaskById(taskId)
	if err != nil {
		log.Errorf("error to get task by id %v", err)
		panic(fmt.Errorf("error to get task by id!"))
	}

	if success {
		if !IsAllowed(task, operator) {
			return nil
		}
		task.Operator = operator
		task.FinishTime = time.Now()
		Update(task, task.Id)
		return task
	} else {
		log.Infof("fail to get task by id %v", err)
		return nil
	}
}

//对指定的任务分配参与者。参与者可以为用户、部门、角色
func AssignTask(taskId string, actors ...string) {
	if len(actors) == 0 {
		return
	} else {
		for _, actorId := range actors {
			if actorId != "" {
				taskActor := &TaskActor{
					Id:      NewUUID(),
					TaskId:  taskId,
					ActorId: actorId,
				}
				Save(taskActor, taskActor.Id)
			}
		}
	}
}

func IsAllowed(task *Task, operator string) bool {
	if operator == string(ER_ADMIN) || operator == string(ER_AUTO) || task.Operator == operator {
		return true
	} else {
		taskActors, _ := task.GetTaskActors()
		return len(taskActors) == 0
	}
}

func CompleteTask(taskId string, operator string, args map[string]interface{}) *Task {
	task := new(Task)
	task.GetTaskById(taskId)
	if IsAllowed(task, operator) {
		historyTask := &HistoryTask{
			Id:           task.Id,
			OrderId:      task.OrderId,
			CreateTime:   task.CreateTime,
			DisplayName:  task.DisplayName,
			TaskName:     task.TaskName,
			TaskType:     task.TaskType,
			ExpireTime:   task.ExpireTime,
			Action:       task.Action,
			ParentTaskId: task.ParentTaskId,
			Variable:     task.Variable,
			PerformType:  task.PerformType,
			FinishTime:   time.Now(),
			Operator:     operator,
			TaskState:    FS_FINISH,
		}
		Save(historyTask, historyTask.Id)
		DeleteById(task, task.Id)
	}
	return task
}

func CreateOrder(process *Process, operator string) *Order {
	//model := process.Model
	order := &Order{
		Id:         NewUUID(),
		ProcessId:  process.Id,
		Creator:    operator,
		CreateTime: time.Now(),
	}

	return order
}

func SaveOrder(order *Order) {
	historyOrder := new(HistoryOrder)
	historyOrder.DataFromOrder(order)

	historyOrder.OrderState = FS_ACTIVITY
	Save(order, order.Id)
	Save(historyOrder, historyOrder.Id)
}

func CompleteOrder(id string) {
	order := new(Order)
	order.GetOrderById(id)

	historyOrder := new(HistoryOrder)
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_FINISH

	Update(historyOrder, historyOrder.Id)
	DeleteById(order, order.Id)
}

func ResumeOrder(id string) {
	historyOrder := new(HistoryOrder)
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_ACTIVITY
	order := historyOrder.Undo()

	Save(order, order.Id)
	Save(historyOrder, historyOrder.Id)

}

func TerminateOrder(id string, operator string) {
	//todo
}
