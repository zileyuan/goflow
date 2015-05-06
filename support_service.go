package goflow

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/lunny/log"
)

//根据OrderID得到活动流程
func GetActiveTasksByOrderId(orderId string) []*Task {
	task := &Task{}
	tasks, _ := task.GetActiveTasksByOrderId(orderId)
	return tasks
}

//得到任务的角色
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
		OrderId:     execution.Order.Id,
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

//保存任务
func SaveTask(task *Task, actors ...string) {
	task.Id = NewUUID()
	task.PerformType = PT_ANY
	Save(task, task.Id)
	AssignTask(task.Id, actors...)
}

//根据已有任务、任务类型、参与者创建新的任务，适用于转派，动态协办处理
func CreateNewTask(taskId string, taskType TASK_TYPE, actors ...string) {
	task := &Task{}
	task.GetTaskById(taskId)
	newTask := *task
	pNewTask := &newTask
	pNewTask.TaskType = taskType
	pNewTask.CreateTime = time.Now()
	pNewTask.ParentTaskId = taskId
	SaveTask(pNewTask, actors...)
}

//驳回任务
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

//撤销任务
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
		Delete(task, task.Id)
	}

	task := historyTask.Undo()
	task.Id = NewUUID()
	task.CreateTime = time.Now()
	Save(task, task.Id)
	AssignTask(task.Id, task.Operator)
	return task
}

//加任务角色
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

//删除任务角色
func RemoveTaskActor(taskId string, actors ...string) {
	task := &Task{}
	task.GetTaskById(taskId)
	if len(actors) > 0 && task.TaskType == TT_MAJOR {
		for _, actorId := range actors {
			taskActor := &TaskActor{
				TaskId:  taskId,
				ActorId: actorId,
			}
			DeleteObj(taskActor)
		}
		v := JsonToMap(task.Variable)
		oldActors := strings.Split(v[DEFAULT_KEY_ACTOR].(string), ",")
		for _, actor := range actors {
			for k, s := range oldActors {
				if s == actor {
					oldActors = StringsRemoveAtIndex(oldActors, k)
					break
				}
			}
		}
		v[DEFAULT_KEY_ACTOR] = oldActors
		task.Variable = MapToJson(v)
		Update(task, task.Id)
	}
}

//结束并且提取任务
func TakeTask(taskId string, operator string) *Task {
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

//是否被授权执行任务
func IsAllowed(task *Task, operator string) bool {
	if operator == string(ER_ADMIN) || operator == string(ER_AUTO) || (task.Operator != "" && task.Operator == operator) {
		return true
	} else {
		taskActors, _ := task.GetTaskActors()
		return len(taskActors) == 0
	}
}

//完成任务
func CompleteTask(taskId string, operator string, args map[string]interface{}) *Task {
	task := &Task{}
	task.GetTaskById(taskId)
	task.Variable = MapToJson(args)
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
		Delete(task, task.Id)

		taskActors, _ := GetTaskActorsByTaskId(historyTask.Id)
		for _, taskActor := range taskActors {
			historyTaskActor := &HistoryTaskActor{
				Id:      taskActor.Id,
				TaskId:  taskActor.TaskId,
				ActorId: taskActor.ActorId,
			}
			Save(historyTaskActor, historyTaskActor.Id)
			Delete(taskActor, taskActor.Id)
		}
	}
	return task
}

//创建Order
func CreateOrder(process *Process, operator string, args map[string]interface{},
	parentId string, parentNodeName string) *Order {
	now := time.Now()
	order := &Order{
		Id:             NewUUID(),
		ParentId:       parentId,
		ParentNodeName: parentNodeName,
		ProcessId:      process.Id,
		Creator:        operator,
		CreateTime:     now,
		LastUpdateTime: now,
		LastUpdator:    operator,
		Variable:       MapToJson(args),
		OrderNo:        GenerateNo(),
	}
	orderNo := args[string(ER_ORDERNO)]
	if orderNo != nil && orderNo.(string) != "" {
		order.OrderNo = orderNo.(string)
	}
	model := process.Model
	if model != nil {
		order.ExpireTime = model.ExpireTime
	}
	SaveOrder(order)
	return order
}

//生成OrderNo
func GenerateNo() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%s_%04d", FormatTime(time.Now(), TIME_LAYOUT), r.Intn(1000))
}

//保存Order
func SaveOrder(order *Order) {
	historyOrder := &HistoryOrder{}
	historyOrder.DataByOrder(order)

	historyOrder.OrderState = FS_ACTIVITY
	Save(order, order.Id)
	Save(historyOrder, historyOrder.Id)
}

//完成Order
func CompleteOrder(id string) {
	order := &Order{}
	order.GetOrderById(id)

	historyOrder := &HistoryOrder{}
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_FINISH

	Update(historyOrder, historyOrder.Id)
	Delete(order, order.Id)
}

//唤醒Order
func ResumeOrder(id string) {
	historyOrder := &HistoryOrder{}
	historyOrder.GetHistoryOrderById(id)
	historyOrder.OrderState = FS_ACTIVITY
	order := historyOrder.Undo()

	Save(order, order.Id)
	Save(historyOrder, historyOrder.Id)

}

//终止Order
func TerminateOrder(id string, operator string) {
	tasks := GetActiveTasksByOrderId(id)
	for _, task := range tasks {
		CompleteTask(task.Id, operator, nil)
	}

	order := &Order{}
	order.GetOrderById(id)
	historyOrder := &HistoryOrder{}
	historyOrder.DataByOrder(order)
	historyOrder.OrderState = FS_TERMINATION
	historyOrder.FinishTime = time.Now()

	Update(historyOrder, historyOrder.Id)
	Delete(order, order.Id)
}

func GetSurrogate(operator string, processName string) string {
	var result []string
	now := time.Now()
	surrogates, _ := GetSurrogateSQL("State = ? and StartTime =< ?  and EndTime >= ? and Operator in (?) and ProcessName in (?)", SS_ENABLE, now, now, operator, processName)
	for _, surrogate := range surrogates {
		result = append(result, surrogate.Surrogate)
	}
	return strings.Join(result, ",")
}
