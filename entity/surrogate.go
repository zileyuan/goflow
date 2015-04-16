package entity

import (
	"goflow/define"
	"time"
)

//委托代理
type Surrogate struct {
	Id          string                  `xorm:"varchar(32) pk notnull"`                            //主键ID
	ProcessName string                  `xorm:"varchar(32) notnull"`                               //流程名称
	Operator    string                  `xorm:"varchar(50) notnull index(IDX_SURROGATE_OPERATOR)"` //授权人
	Surrogate   string                  `xorm:"varchar(50)"`                                       //代理人
	OpTime      time.Time               `xorm:"datetime"`                                          //操作时间
	StartTime   time.Time               `xorm:"datetime"`                                          //开始时间
	EndTime     time.Time               `xorm:"datetime"`                                          //结束时间
	State       define.SURROGATE_STATUS `xorm:"tinyint"`                                           //状态
}
