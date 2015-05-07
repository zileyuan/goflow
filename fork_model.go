package goflow

//XML分叉节点
type ForkModel struct {
	NodeModel
}

//执行
func (p *ForkModel) Exec(execution *Execution) {
	p.RunOutTransition(execution)
}
