package goflow

import (
	"time"

	"github.com/lunny/log"
)

//任务实体类
type Task struct {
	Id           string       `xorm:"varchar(48) pk notnull"`     //主键ID
	Version      int          `xorm:"tinyint"`                    //版本
	OrderId      string       `xorm:"varchar(48) notnull index"`  //流程实例ID
	TaskName     string       `xorm:"varchar(100) notnull index"` //任务名称
	DisplayName  string       `xorm:"varchar(200) notnull"`       //任务显示名称
	PerformType  PERFORM_TYPE `xorm:"varchar(16)"`                //任务参与方式
	TaskType     TASK_TYPE    `xorm:"varchar(16) notnull"`        //任务类型
	Operator     string       `xorm:"varchar(48)"`                //任务处理者ID
	CreateTime   time.Time    `xorm:"datetime"`                   //任务创建时间
	FinishTime   time.Time    `xorm:"datetime"`                   //任务完成时间
	ExpireTime   int          `xorm:"datetime"`                   //期望任务完成时间
	Action       string       `xorm:"varchar(200)"`               //任务关联的Action(WEB为表单URL)
	ParentTaskId string       `xorm:"varchar(48) index"`          //父任务ID
	Variable     string       `xorm:"varchar(2000)"`              //任务附属变量(json存储)
}

func (p *Task) GetTaskById(id string) (bool, error) {
	p.Id = id
	success, err := orm.Get(p)
	return success, err
}

func (p *Task) Update() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Update(p)
	log.Infof("Task %d updated", p.Id)
	return err
}

func (p *Task) Save() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(p)
	log.Infof("Task %d inserted", p.Id)
	return err
}

func (p *Task) Delete() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Delete(p)
	log.Infof("Task %d deleted", p.Id)
	return err
}

func (p *Task) GetActiveTasks(orderId string) ([]*Task, error) {
	p.OrderId = orderId
	tasks := make([]*Task, 0)
	err := orm.Find(&tasks, p)
	return tasks, err
}
