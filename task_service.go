package goflow

import (
	"fmt"
	"time"

	"github.com/lunny/log"
)

type TaskService struct {
}

func (p *TaskService) CreateTask(taskModel *TaskModel, execution *Execution) []*Task {
	//todo
	return nil
}

func (p *TaskService) CreateNewTask(taskId string, taskType TASK_TYPE, actors ...string) []*Task {
	//todo
	return nil
}

func (p *TaskService) RejectTask(processModel *ProcessModel, task *Task) *Task {
	//todo
	return nil
}

func (p *TaskService) CompleteTask(taskId string, operator string, args map[string]interface{}) *Task {
	//todo
	task := new(Task)
	task.GetTaskById(taskId)
	if p.IsAllowed(task, operator) {
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
		historyTask.Save()
		task.Delete()
	}
	return task
}

func (p *TaskService) IsAllowed(task *Task, operator string) bool {
	//todo
	return true
}

func (p *TaskService) Take(taskId string, operator string) *Task {
	task := &Task{}
	success, err := task.GetTaskById(taskId)
	if err != nil {
		log.Errorf("error to get task by id %v", err)
		panic(fmt.Errorf("error to get task by id!"))
	}

	if success {
		if !p.IsAllowed(task, operator) {
			return nil
		}
		task.Operator = operator
		task.FinishTime = time.Now()
		task.Update()
		return task
	} else {
		log.Infof("fail to get task by id %v", err)
		return nil
	}
}
