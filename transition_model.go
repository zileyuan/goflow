package goflow

type TransitionModel struct {
	BaseModel
	Source  INodeModel ``                //变迁的源节点引用
	Target  INodeModel ``                //变迁的目标节点引用
	To      string     `xml:"to,attr"`   //变迁的目的节点
	Expr    string     `xml:"expr,attr"` //变迁的条件表达式，用于decision
	Enabled bool       ``                //当前变迁路径是否可用
}

func (p *TransitionModel) Execute(execution *Execution) {
	if p.Enabled {
		switch p.Target.(type) {
		case *TaskModel:
			taskModel := p.Target.(*TaskModel)
			taskModel.CreateHandle(execution)
		default:
			p.Target.Execute(execution)
		}
	} else {
		return
	}
}
