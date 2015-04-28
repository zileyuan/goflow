package goflow

import (
	"time"
)

//流程定义process元素
type ProcessModel struct {
	BaseModel
	StartNodes       []*StartModel      `xml:"start"`               //开始节点
	EndModels        []*EndModel        `xml:"end"`                 //结束阶段
	TaskModels       []*TaskModel       `xml:"task"`                //任务节点
	DecisionModels   []*DecisionModel   `xml:"decision"`            //决策节点
	ForkModels       []*ForkModel       `xml:"fork"`                //分支节点
	JoinModels       []*JoinModel       `xml:"join"`                //合并节点
	SubProcessModels []*SubProcessModel `xml:"subprocess"`          //子流程节点
	InstanceAction   string             `xml:"instanceAction,attr"` //流程实例启动Action(web为url)
	ExpireTime       time.Time          `xml:"expireTime,attr"`     //期望完成时间
}

func (p *ProcessModel) GetStart() *StartModel {
	if len(p.StartNodes) > 0 {
		return p.StartNodes[0]
	}
	return nil
}

func (p *ProcessModel) GetNode(name string) INodeModel {
	//todo
	return nil
}
