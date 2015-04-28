package goflow

type StartModel struct {
	NodeModel
}

func (p *StartModel) Execute(execution *Execution) error {
	return p.RunOutTransition(execution)
}
