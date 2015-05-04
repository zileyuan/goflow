package goflow

import (
	"time"
)

//XML定义process节点元素
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

//得到开始节点
func (p *ProcessModel) GetStart() *StartModel {
	if len(p.StartNodes) > 0 {
		return p.StartNodes[0]
	}
	return nil
}

//根据任务名，是否包含子流程
func (p *ProcessModel) ContainsSubProcessNodeNames(nodeNames ...string) bool {
	for _, node := range p.SubProcessModels {
		for _, nodeName := range nodeNames {
			if node.Name == nodeName {
				return true
			}
		}
	}
	return false
}

//根据任务名，是否包含任务
func (p *ProcessModel) ContainsTaskNodeNames(nodeNames ...string) bool {
	for _, node := range p.TaskModels {
		for _, nodeName := range nodeNames {
			if node.Name == nodeName {
				return true
			}
		}
	}
	return false
}

//根据名称得到节点
func (p *ProcessModel) GetNode(name string) INodeModel {
	for _, v := range p.StartNodes {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.EndModels {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.TaskModels {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.DecisionModels {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.ForkModels {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.JoinModels {
		if v.Name == name {
			return v
		}
	}

	for _, v := range p.SubProcessModels {
		if v.Name == name {
			return v
		}
	}
	return nil
}
