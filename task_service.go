package goflow

import (
	"time"
)

type TaskService struct {
}

func (p *TaskService) CreateTask(taskModel *TaskModel, execution *Execution) []*Task {
	//todo
	return []*Task{}
}

func (p *TaskService) RejectTask(processModel *ProcessModel, task *Task) *Task {
	//todo
	return &Task{}
}

func (p *TaskService) Complete(taskId string, operator string, args map[string]interface{}) *Task {
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
			ActionUrl:    task.ActionUrl,
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
