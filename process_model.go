package goflow

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/lunny/log"
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
	ExpireTime       string             `xml:"expireTime,attr"`     //期望完成时间
	Models           []INodeModel       `xml:"-"`                   //上面所有Node节点(Start\End\Task\Decision\Fork\Join\SubProcess)的集合
}

func (p *ProcessModel) BuildRelationship(content []byte, processService *ProcessService) {
	//解析xml内容
	err := xml.Unmarshal(content, p)
	if err != nil {
		log.Errorf("error to unmarshal xml %v", err)
		panic(fmt.Errorf("error to unmarshal xml!"))
	}

	//建立新的节点集合
	p.Models = make([]INodeModel, 0)
	for _, v := range p.StartNodes {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.EndModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.TaskModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.DecisionModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.ForkModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.JoinModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}
	for _, v := range p.SubProcessModels {
		v.INodeModel = v
		p.Models = append(p.Models, v)
	}

	//建立Node之间的关
	for _, v := range p.Models {
		for _, tm := range v.GetOutputs() {
			node := p.GetNode(tm.To)
			if node != nil {
				node.AddInputs(tm)
				tm.Target = node
			}
		}
		v.BuildInterceptors(processService)
	}
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
			if strings.ToUpper(node.Name) == strings.ToUpper(nodeName) {
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
			if strings.ToUpper(node.Name) == strings.ToUpper(nodeName) {
				return true
			}
		}
	}
	return false
}

//根据名称得到节点
func (p *ProcessModel) GetNode(name string) INodeModel {
	for _, v := range p.Models {
		if strings.ToUpper(v.GetName()) == strings.ToUpper(name) {
			return v
		}
	}
	return nil
}
