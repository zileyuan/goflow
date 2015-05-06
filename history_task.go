package goflow

import "time"

//任务实体类
type HistoryTask struct {
	Id           string        `xorm:"varchar(36) pk notnull"`     //主键ID
	OrderId      string        `xorm:"varchar(36) notnull index"`  //流程实例ID
	TaskName     string        `xorm:"varchar(100) notnull index"` //任务名称
	DisplayName  string        `xorm:"varchar(200) notnull"`       //任务显示名称
	PerformType  PERFORM_ORDER `xorm:"tinyint"`                    //任务参与方式
	TaskType     TASK_ORDER    `xorm:"tinyint notnull"`            //任务类型
	Operator     string        `xorm:"varchar(36)"`                //任务处理者ID
	CreateTime   time.Time     `xorm:"datetime notnull"`           //任务创建时间
	FinishTime   time.Time     `xorm:"datetime"`                   //任务完成时间
	ExpireTime   time.Time     `xorm:"datetime"`                   //期望任务完成时间
	Action       string        `xorm:"varchar(200)"`               //任务关联的Action(WEB为表单URL)
	ParentTaskId string        `xorm:"varchar(36) index"`          //父任务ID
	Variable     string        `xorm:"varchar(2000)"`              //任务附属变量(json存储)
	TaskState    FLOW_STATUS   `xorm:"tinyint notnull"`            //任务状态
}

//根据ID得到HistoryTask
func (p *HistoryTask) GetHistoryTaskById(id string) bool {
	p.Id = id
	success, err := orm.Get(p)
	PanicIf(err, "fail to GetHistoryTaskById")
	return success
}

//通过HistoryTask生成Task
func (p *HistoryTask) Undo() *Task {
	task := &Task{
		Id:           p.Id,
		TaskName:     p.TaskName,
		DisplayName:  p.DisplayName,
		TaskType:     p.TaskType,
		ExpireTime:   p.ExpireTime,
		Action:       p.Action,
		ParentTaskId: p.ParentTaskId,
		Variable:     p.Variable,
		PerformType:  p.PerformType,
		Operator:     p.Operator,
	}
	return task
}
