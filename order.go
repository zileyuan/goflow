package goflow

import "time"

//流程工作单实体类（一般称为流程实例）
type Order struct {
	Id             string    `xorm:"varchar(36) pk notnull"`    //主键ID
	Version        int       `xorm:"tinyint"`                   //版本
	ProcessId      string    `xorm:"varchar(36) notnull index"` //流程定义ID
	Creator        string    `xorm:"varchar(36)"`               //流程实例创建者ID
	CreateTime     time.Time `xorm:"datetime notnull"`          //流程实例创建时间
	ParentId       string    `xorm:"varchar(36) index"`         //流程实例为子流程时，该字段标识父流程实例ID
	ParentNodeName string    `xorm:"varchar(100)"`              //流程实例为子流程时，该字段标识父流程哪个节点模型启动的子流程
	ExpireTime     time.Time `xorm:"datetime"`                  //流程实例期望完成时间
	LastUpdateTime time.Time `xorm:"datetime"`                  //流程实例上一次更新时间
	LastUpdator    string    `xorm:"varchar(36)"`               //流程实例上一次更新人员ID
	Priority       int       `xorm:"tinyint"`                   //流程实例优先级
	OrderNo        string    `xorm:"varchar(100) index"`        //流程实例编号
	Variable       string    `xorm:"varchar(2000)"`             //流程实例附属变量
}

//根据ID得到Order
func (p *Order) GetOrderById(id string) bool {
	p.Id = id
	success, err := orm.Get(p)
	PanicIf(err, "fail to GetOrderById")
	return success
}

//得到活动的Order（通过SQL）
func GetActiveOrdersSQL(querystring string, args ...interface{}) []*Order {
	orders := make([]*Order, 0)
	err := orm.Where(querystring, args).Find(&orders)
	PanicIf(err, "fail to GetActiveOrdersSQL")
	return orders
}
