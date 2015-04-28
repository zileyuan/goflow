package goflow

import (
	"time"

	"github.com/lunny/log"
)

//任务实体类
type HistoryTask struct {
	Id           string       `xorm:"varchar(48) pk notnull"`     //主键ID
	OrderId      string       `xorm:"varchar(48) notnull index"`  //流程实例ID
	TaskName     string       `xorm:"varchar(100) notnull index"` //任务名称
	DisplayName  string       `xorm:"varchar(200) notnull"`       //任务显示名称
	PerformType  PERFORM_TYPE `xorm:"varchar(16)"`                //任务参与方式
	TaskType     TASK_TYPE    `xorm:"varchar(16) notnull"`        //任务类型
	Operator     string       `xorm:"varchar(48)"`                //任务处理者ID
	CreateTime   time.Time    `xorm:"datetime notnull"`           //任务创建时间
	FinishTime   time.Time    `xorm:"datetime"`                   //任务完成时间
	ExpireTime   int          `xorm:"datetime"`                   //期望任务完成时间
	Action       string       `xorm:"varchar(200)"`               //任务关联的Action(WEB为表单URL)
	ParentTaskId string       `xorm:"varchar(48) index"`          //父任务ID
	Variable     string       `xorm:"varchar(2000)"`              //任务附属变量(json存储)
	TaskState    FLOW_STATUS  `xorm:"tinyint notnull"`            //任务状态
}

func (p *HistoryTask) Update() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Update(p)
	log.Infof("HistoryTask %d updated", p.Id)
	return err
}

func (p *HistoryTask) Save() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(p)
	log.Infof("HistoryTask %d inserted", p.Id)
	return err
}
