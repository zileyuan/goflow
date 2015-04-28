package goflow

import (
	"time"
)

type TaskModel struct {
	NodeModel
	Assignee    string       `xml:"assignee,attr"`    //参与者变量名称
	PerformType PERFORM_TYPE `xml:"performType,attr"` //参与方式
	TaskType    TASK_TYPE    `xml:"taskType,attr"`    //任务类型
	AutoExecute bool         `xml:"autoExecute,attr"` //是否自动执行
	ExpireTime  time.Time    `xml:"-"`                //期望完成时间
}

func (p *TaskModel) CreateHandle(execution *Execution) error {
	tasks := execution.Engine.CreateTask(p, execution)
	execution.AddTasks(tasks)
	return nil
}

func (p *TaskModel) MergeActorHandle(execution *Execution) error {
	return nil
}

func (p *TaskModel) Execute(execution *Execution) error {
	//todo
	return nil
}
