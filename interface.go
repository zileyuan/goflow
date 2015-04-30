package goflow

type INodeModel interface {
	Execute(execution *Execution) error
	GetInputs() []*TransitionModel
	GetOutputs() []*TransitionModel
}
