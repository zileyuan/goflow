package goflow

//XML分叉节点
type ForkModel struct {
	NodeModel
}

//执行
func (p *ForkModel) Execute(execution *Execution) {
	p.PrevIntercept(execution)
	p.RunOutTransition(execution)
	p.PostIntercept(execution)
}
