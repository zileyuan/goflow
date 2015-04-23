package goflow

type NodeModel struct {
	BaseModel
	Inputs  []*TransitionModel ``                 //输入变迁集合
	Outputs []*TransitionModel `xml:"transition"` //输出变迁集合
}

func (p *NodeModel) GetInputs() []*TransitionModel {
	return p.Inputs
}

func (p *NodeModel) GetOutputs() []*TransitionModel {
	return p.Outputs
}

func (p *NodeModel) RunOutTransition(execution *Execution) {
	for _, tm := range p.Outputs {
		tm.Enabled = true
		tm.Execute(execution)
	}
}
