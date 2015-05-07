package goflow

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
func (p *TransitionModel) Execute(execution *Execution) {
	if p.Enabled {
		switch p.Target.(type) {
		case *TaskModel:
			taskModel := p.Target.(*TaskModel)
			tasks := CreateTaskHandle(taskModel, execution)
			AutoExecuteTask(taskModel, execution, tasks)
		case *SubProcessModel:
			subProcessModel := p.Target.(*SubProcessModel)
			StartSubProcessHandle(subProcessModel, execution)
		default:
			p.Target.Execute(execution)
		}
	}
}
