package goflow

type StartModel struct {
	NodeModel
}

func (p *StartModel) Execute(execution *Execution) {
	p.RunOutTransition(execution)
}
