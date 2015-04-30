package goflow

import (
	"time"
)

type TaskModel struct {
	WorkModel
	Assignee    string       `xml:"assignee,attr"`    //参与者变量名称
	PerformType PERFORM_TYPE `xml:"performType,attr"` //参与方式
	TaskType    TASK_TYPE    `xml:"taskType,attr"`    //任务类型
	AutoExecute bool         `xml:"autoExecute,attr"` //是否自动执行
	ExpireTime  time.Time    `xml:"-"`                //期望完成时间
}

func (p *TaskModel) Execute(execution *Execution) error {
	if p.PerformType == PT_ANY {
		p.RunOutTransition(execution)
	} else {
		MergeActorHandle(p, execution)
		if execution.IsMerged {
			p.RunOutTransition(execution)
		}
	}
	return nil
}

func CreateHandle(tm *TaskModel, execution *Execution) error {
	tasks := CreateTask(tm, execution)
	execution.Tasks = append(execution.Tasks, tasks...)
	return nil
}

func MergeActorHandle(tm *TaskModel, execution *Execution) error {
	activeNodes := []string{tm.Name}
	return MergeHandle(execution, activeNodes)
}
