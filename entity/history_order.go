package entity

import (
	"goflow/define"
	"time"
)

//历史流程实例实体类
type HistoryOrder struct {
	Id         string              `xorm:"varchar(32) pk notnull"`                              //主键ID
	ProcessId  string              `xorm:"varchar(32) notnull index(IDX_HIST_ORDER_PROCESSID)"` //流程定义ID
	Creator    string              `xorm:"varchar(100)"`                                        //流程实例创建者ID
	CreateTime time.Time           `xorm:"datetime notnull"`                                    //流程实例创建时间
	ParentId   string              `xorm:"varchar(32) index(FK_HIST_ORDER_PARENTID)"`           //流程实例为子流程时，该字段标识父流程实例ID
	ExpireTime time.Time           `xorm:"datetime"`                                            //流程实例期望完成时间
	Priority   int                 `xorm:"tinyint"`                                             //流程实例优先级
	OrderNo    string              `xorm:"varchar(100) index(IDX_HIST_ORDER_NO)"`               //流程实例编号
	Variable   string              `xorm:"varchar(2000)"`                                       //流程实例附属变量
	OrderState define.ORDER_STATUS `xorm:"tinyint notnull"`                                     //流程实例状态
	FinishTime time.Time           `xorm:"datetime"`                                            //完成时间
}
