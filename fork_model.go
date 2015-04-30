package goflow

type ForkModel struct {
	NodeModel
}

func (p *ForkModel) Execute(execution *Execution) error {
	return p.RunOutTransition(execution)
}
