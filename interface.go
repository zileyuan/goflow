package goflow

//节点模型需要实现的接口
type INodeModel interface {
	GetName() string
	Execute(execution *Execution) error
	GetInputs() []*TransitionModel
	GetOutputs() []*TransitionModel
	AddInputs(tm ...*TransitionModel)
}
