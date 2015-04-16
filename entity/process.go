package entity

import (
	"github.com/lunny/log"
	"goflow/define"
	"goflow/model"
	"time"
)

//流程定义实体类
type Process struct {
	Id          string                `xorm:"varchar(32) pk notnull"`               //主键ID
	Version     int                   `xorm:"tinyint"`                              //版本
	Name        string                `xorm:"varchar(100) index(IDX_PROCESS_NAME)"` //流程定义名称
	DisplayName string                `xorm:"varchar(200)"`                         //流程定义显示名称
	InstanceUrl string                `xorm:"varchar(200)"`                         //当前流程的实例URL(一般为流程第一步的URL),该字段可以直接打开流程申请的表单
	State       define.PROCESS_STATUS `xorm:"tinyint"`                              //是否可用的开关
	CreateTime  time.Time             `xorm:"datetime"`                             //创建时间
	Creator     string                `xorm:"varchar(50)"`                          //创建人
	Content     string                `xorm:"text"`                                 //流程定义XML

	Model model.ProcessModel //Model对象
}

func (p *Process) Save() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(p)
	log.Infof("Process %d inserted", p.Id)
	return err
}
