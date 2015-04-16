package model

import (
	"goflow/define"

	"time"
)

type TaskModel struct {
	Assignee    string              `xml:"assignee,attr"`    //参与者变量名称
	PerformType define.PERFORM_TYPE `xml:"performType,attr"` //参与方式
	TaskType    define.TASK_TYPE    `xml:"taskType,attr"`    //任务类型
	AutoExecute bool                `xml:"autoExecute,attr"` //是否自动执行
	ExpireTime  time.Time           ``                       //期望完成时间
}
