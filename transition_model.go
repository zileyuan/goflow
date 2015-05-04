package goflow

import "fmt"

//XML变迁节点元素
type TransitionModel struct {
	BaseModel
	Source  INodeModel `xml:"-"`         //变迁的源节点引用
	Target  INodeModel `xml:"-"`         //变迁的目标节点引用
	To      string     `xml:"to,attr"`   //变迁的目的节点
	Expr    string     `xml:"expr,attr"` //变迁的条件表达式，用于decision
	Enabled bool       `xml:"-"`         //当前变迁路径是否可用
}

//执行
func (p *TransitionModel) Execute(execution *Execution) error {
	if p.Enabled {
		switch p.Target.(type) {
		case *TaskModel:
			taskModel := p.Target.(*TaskModel)
			return CreateHandle(taskModel, execution)
		default:
			return p.Target.Execute(execution)
		}
	} else {
		return fmt.Errorf("Transition DisEnable!")
	}
}
