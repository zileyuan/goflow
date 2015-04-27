package goflow

import (
	"time"

	"github.com/lunny/log"
)

//流程定义实体类
type Process struct {
	Id          string        `xorm:"varchar(32) pk notnull"`               //主键ID
	Version     int           `xorm:"tinyint"`                              //版本
	Name        string        `xorm:"varchar(100) index(IDX_PROCESS_NAME)"` //流程定义名称
	DisplayName string        `xorm:"varchar(200)"`                         //流程定义显示名称
	InstanceUrl string        `xorm:"varchar(200)"`                         //当前流程的实例URL(一般为流程第一步的URL),该字段可以直接打开流程申请的表单
	State       FLOW_STATUS   `xorm:"tinyint"`                              //是否可用的开关
	CreateTime  time.Time     `xorm:"datetime"`                             //创建时间
	Creator     string        `xorm:"varchar(50)"`                          //创建人
	Content     []byte        `xorm:"text"`                                 //流程定义XML
	Model       *ProcessModel `xorm:"-"`                                    //Model对象
}

func (p *Process) Save() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.InsertOne(p)
	log.Infof("Process %d inserted", p.Id)
	return err
}

func (p *Process) Update() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Update(p)
	log.Infof("Process %d updated", p.Id)
	return err
}

func (p *Process) Delete() error {
	session := orm.NewSession()
	defer session.Close()
	_, err := session.Id(p.Id).Delete(p)
	log.Infof("Process %d deleted", p.Id)
	return err
}

func (p *Process) GetProcessById(id string) (bool, error) {
	p.Id = id
	success, err := orm.Get(p)
	return success, err
}

func (p *Process) GetProcessByVersion(name string, version int) (bool, error) {
	//todo
	return true, nil
}

func (p *Process) SetModel(model *ProcessModel) {
	p.Model = model
	p.Name = model.Name
	p.DisplayName = model.DisplayName
	p.InstanceUrl = model.InstanceUrl
}
