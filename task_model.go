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

func (p *TaskModel) CreateHandle(execution *Execution) error {
	tasks := CreateTask(p, execution)
	execution.Tasks = append(execution.Tasks, tasks...)
	return nil
}

func (p *TaskModel) MergeActorHandle(execution *Execution) error {
	activeNodes := []string{p.Name}
	return p.MergeHandle(execution, activeNodes)
}

func (p *TaskModel) Execute(execution *Execution) error {
	if p.PerformType == PT_ANY {
		p.RunOutTransition(execution)
	} else {
		p.MergeActorHandle(execution)
		if execution.IsMerged {
			p.RunOutTransition(execution)
		}
	}
	return nil
}
