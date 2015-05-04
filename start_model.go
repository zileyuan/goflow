package goflow

//XML开始节点元素
type StartModel struct {
	NodeModel
}

//执行
func (p *StartModel) Execute(execution *Execution) error {
	return p.RunOutTransition(execution)
}
