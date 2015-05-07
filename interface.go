package goflow

//节点模型需要实现的接口
type INodeModel interface {
	GetName() string
	Exec(execution *Execution)
	Execute(execution *Execution)
	GetInputs() []*TransitionModel
	GetOutputs() []*TransitionModel
	AddInputs(tm ...*TransitionModel)
	BuildInterceptors(processService *ProcessService)
}

//拦截器需要实现的接口
type IInterceptor interface {
	GetName() string //Unique
	Intercept(execution *Execution)
	Clone() IInterceptor
}
