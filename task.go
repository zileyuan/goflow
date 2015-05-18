package goflow

import (
	"strings"
	"time"
)

//任务实体类
type Task struct {
	Id           string        `xorm:"varchar(36) pk notnull"`     //主键ID
	Version      int           `xorm:"tinyint"`                    //版本
	OrderId      string        `xorm:"varchar(36) notnull index"`  //流程实例ID
	TaskName     string        `xorm:"varchar(100) notnull index"` //任务名称
	DisplayName  string        `xorm:"varchar(200) notnull"`       //任务显示名称
	PerformType  PERFORM_ORDER `xorm:"tinyint"`                    //任务参与方式
	TaskType     TASK_ORDER    `xorm:"tinyint notnull"`            //任务类型
	Operator     string        `xorm:"varchar(36)"`                //任务处理者ID
	CreateTime   time.Time     `xorm:"datetime"`                   //任务创建时间
	FinishTime   time.Time     `xorm:"datetime"`                   //任务完成时间
	ExpireTime   time.Time     `xorm:"datetime"`                   //期望任务完成时间
	RemindTime   time.Time     `xorm:"datetime"`                   //提醒时间
	Action       string        `xorm:"varchar(200)"`               //任务关联的Action(WEB为表单URL)
	ParentTaskId string        `xorm:"varchar(36) index"`          //父任务ID
	Variable     string        `xorm:"varchar(2000)"`              //任务附属变量(json存储)
	Model        *TaskModel    `xorm:"-"`                          //Model对象
}

//根据ID得到任务
func (p *Task) GetTaskById(id string) bool {
	p.Id = id
	success, err := orm.Get(p)
	PanicIf(err, "fail to GetTaskById")
	return success
}

//得到活动任务
func (p *Task) GetActiveTasks() []*Task {
	tasks := make([]*Task, 0)
	err := orm.Find(&tasks, p)
	PanicIf(err, "fail to GetActiveTasks")
	return tasks
}

//根据OrderID得到活动任务
func (p *Task) GetActiveTasksByOrderId(orderId string) []*Task {
	p.OrderId = orderId
	tasks := make([]*Task, 0)
	err := orm.Find(&tasks, p)
	PanicIf(err, "fail to GetActiveTasksByOrderId")
	return tasks
}

//得到任务角色
func (p *Task) GetTaskActors() []*TaskActor {
	taskActors := make([]*TaskActor, 0)
	taskActor := &TaskActor{
		TaskId: p.Id,
	}
	err := orm.Find(&taskActors, taskActor)
	PanicIf(err, "fail to GetTaskActors")
	return taskActors
}

//得到下一个ANY类型的任务
func GetNextAnyActiveTasks(parentTaskId string) []*Task {
	task := &Task{
		ParentTaskId: parentTaskId,
	}
	tasks := make([]*Task, 0)
	err := orm.Find(&tasks, task)
	PanicIf(err, "fail to GetNextAnyActiveTasks")
	return tasks
}

//得到下一个ALL类型的任务
func GetNextAllActiveTasks(orderId string, taskName string, parentTaskId string) []*Task {
	historyTask := &HistoryTask{
		OrderId:      orderId,
		TaskName:     taskName,
		ParentTaskId: parentTaskId,
	}
	historyTasks := make([]*HistoryTask, 0)
	err := orm.Find(&historyTasks, historyTask)
	PanicIf(err, "fail to GetNextAllActiveTasks One")

	ids := make([]string, 0)
	for _, historyTask = range historyTasks {
		ids = append(ids, historyTask.Id)
	}
	tasks := make([]*Task, 0)
	err = orm.Where(`"ParentTaskId" in (?)`, strings.Join(ids, ",")).Find(&tasks)
	PanicIf(err, "fail to GetNextAllActiveTasks Two")

	return tasks
}

//得到活动的任务（通过SQL）
func GetActiveTasksSQL(querystring string, args ...interface{}) ([]*Task, error) {
	tasks := make([]*Task, 0)
	err := orm.Where(querystring, args).Find(&tasks)
	return tasks, err
}
