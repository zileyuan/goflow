package goflow

import (
	"time"

	"github.com/lunny/log"
)

//历史流程实例实体类
type HistoryOrder struct {
	Id         string      `xorm:"varchar(36) pk notnull"`    //主键ID
	ProcessId  string      `xorm:"varchar(36) notnull index"` //流程定义ID
	Creator    string      `xorm:"varchar(36)"`               //流程实例创建者ID
	CreateTime time.Time   `xorm:"datetime notnull"`          //流程实例创建时间
	ParentId   string      `xorm:"varchar(36) index"`         //流程实例为子流程时，该字段标识父流程实例ID
	ExpireTime time.Time   `xorm:"datetime"`                  //流程实例期望完成时间
	Priority   int         `xorm:"tinyint"`                   //流程实例优先级
	OrderNo    string      `xorm:"varchar(100) index"`        //流程实例编号
	Variable   string      `xorm:"varchar(2000)"`             //流程实例附属变量
	OrderState FLOW_STATUS `xorm:"tinyint notnull"`           //流程实例状态
	FinishTime time.Time   `xorm:"datetime"`                  //完成时间
}

func (p *HistoryOrder) DataFromOrder(order *Order) {
	p.Id = order.Id
	p.ProcessId = order.ProcessId
	p.CreateTime = order.CreateTime
	p.ExpireTime = order.ExpireTime
	p.Creator = order.Creator
	p.ParentId = order.ParentId
	p.Priority = order.Priority
	p.OrderNo = order.OrderNo
	p.Variable = order.Variable
}

func (p *HistoryOrder) Save() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(p)
	log.Infof("HistoryOrder %d inserted", p.Id)
	return err
}

func (p *HistoryOrder) Update() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Update(p)
	log.Infof("HistoryOrder %d updated", p.Id)
	return err
}

func (p *HistoryOrder) Delete() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Delete(p)
	log.Infof("HistoryOrder %d deleted", p.Id)
	return err
}

func (p *HistoryOrder) GetHistoryOrderById(id string) (bool, error) {
	p.Id = id
	success, err := orm.Get(p)
	return success, err
}

func (p *HistoryOrder) Undo() *Order {
	order := &Order{
		Id:             p.Id,
		ProcessId:      p.ProcessId,
		CreateTime:     p.CreateTime,
		ExpireTime:     p.ExpireTime,
		Creator:        p.Creator,
		LastUpdator:    p.Creator,
		LastUpdateTime: p.FinishTime,
		ParentId:       p.ParentId,
		Priority:       p.Priority,
		OrderNo:        p.OrderNo,
		Variable:       p.Variable,
		Version:        0,
	}
	return order
}
