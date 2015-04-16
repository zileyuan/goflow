package model

import (
	"time"
)

//流程定义process元素
type ProcessModel struct {
	BaseModel
	StartNodes      []StartModel    `xml:"start"`                //开始节点
	EndModels       []EndModel      `xml:"end"`                  //结束阶段
	TaskModels      []TaskModel     `xml:"task"`                 //任务节点
	DecisionModels  []DecisionModel `xml:"decision"`             //决策节点
	ForkModels      []ForkModel     `xml:"fork"`                 //分支节点
	JoinModels      []JoinModel     `xml:"join"`                 //合并节点
	InstanceUrl     string          `xml:"instanceUrl,attr"`     //流程实例启动url
	ExpireTime      time.Time       `xml:"expireTime,attr"`      //期望完成时间
	InstanceNoClass string          `xml:"instanceNoClass,attr"` //实例编号生成的class
}
