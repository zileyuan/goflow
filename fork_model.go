package goflow

//XML分叉节点
type ForkModel struct {
	NodeModel
}

//执行
func (p *ForkModel) Execute(execution *Execution) error {
	return p.RunOutTransition(execution)
}
