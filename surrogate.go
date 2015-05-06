package goflow

import (
	"time"
)

//委托代理
type Surrogate struct {
	Id          string           `xorm:"varchar(36) pk notnull"`    //主键ID
	ProcessName string           `xorm:"varchar(36) notnull"`       //流程名称
	Operator    string           `xorm:"varchar(36) notnull index"` //授权人
	Surrogate   string           `xorm:"varchar(36)"`               //代理人
	OpTime      time.Time        `xorm:"datetime"`                  //操作时间
	StartTime   time.Time        `xorm:"datetime"`                  //开始时间
	EndTime     time.Time        `xorm:"datetime"`                  //结束时间
	State       SURROGATE_STATUS `xorm:"tinyint"`                   //状态
}

//得到代理人（通过SQL）
func GetSurrogateSQL(querystring string, args ...interface{}) []*Surrogate {
	surrogates := make([]*Surrogate, 0)
	err := orm.Where(querystring, args).Find(&surrogates)
	PanicIf(err, "fail to GetSurrogateSQL")
	return surrogates
}
