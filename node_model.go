package goflow

import "strings"

//XML节点通用信息
type NodeModel struct {
	BaseModel
	Inputs                  []*TransitionModel `xml:"-"`                     //输入变迁集合
	Outputs                 []*TransitionModel `xml:"transition"`            //输出变迁集合
	PrevInterceptorsSetting string             `xml:"prevInterceptors,attr"` //局部前置拦截器
	PostInterceptorsSetting string             `xml:"postInterceptors,attr"` //局部后置拦截器
	PrevInterceptors        []IInterceptor     `xml:"-"`                     //局部前置拦截器对象
	PostInterceptors        []IInterceptor     `xml:"-"`                     //局部后置拦截器对象
}

//建立拦截器
func (p *NodeModel) BuildInterceptors(processService *ProcessService) {
	p.PrevInterceptors = make([]IInterceptor, 0)
	p.PostInterceptors = make([]IInterceptor, 0)

	for _, v := range processService.InnerInterceptorCache {
		p.PrevInterceptors = append(p.PrevInterceptors, v.Clone())
	}

	pris := strings.Split(p.PrevInterceptorsSetting, ",")
	pois := strings.Split(p.PostInterceptorsSetting, ",")

	for _, s := range pris {
		if v := processService.GetCustomInterceptor(s); v != nil {
			p.PrevInterceptors = append(p.PrevInterceptors, v.Clone())
		}
	}

	for _, s := range pois {
		if v := processService.GetCustomInterceptor(s); v != nil {
			p.PostInterceptors = append(p.PostInterceptors, v.Clone())
		}
	}
}

func (p *NodeModel) AddInputs(tm ...*TransitionModel) {
	p.Inputs = append(p.Inputs, tm...)
}

//得到输入变迁
func (p *NodeModel) GetInputs() []*TransitionModel {
	return p.Inputs
}

//得到输出变迁
func (p *NodeModel) GetOutputs() []*TransitionModel {
	return p.Outputs
}

//运行变迁
func (p *NodeModel) RunOutTransition(execution *Execution) {
	for _, tm := range p.Outputs {
		tm.Enabled = true
		tm.Execute(execution)
	}
}

//执行拦截器
func (p *NodeModel) PrevIntercept(execution *Execution) {
	Intercept(p.PrevInterceptors, execution)
}

//执行拦截器
func (p *NodeModel) PostIntercept(execution *Execution) {
	Intercept(p.PostInterceptors, execution)
}

//执行拦截器
func Intercept(interceptors []IInterceptor, execution *Execution) {
	for _, interceptor := range interceptors {
		interceptor.Intercept(execution)
	}
}

//合并处理通用流程
func MergeHandle(execution *Execution, activeNodes []string) {
	processModel := execution.Process.Model

	isSubProcessMerged := false
	if processModel.ContainsSubProcessNodeNames(activeNodes...) {
		orders, _ := GetActiveOrdersSQL("ParentId = ? and Id <> ?", execution.Order.Id, execution.ChildOrderId)
		isSubProcessMerged = len(orders) == 0
	}

	isTaskMerged := false
	if isSubProcessMerged && processModel.ContainsTaskNodeNames(activeNodes...) {
		tasks, _ := GetActiveTasksSQL("OrderId = ? and Id <> ? and Name in (?)", execution.Order.Id, execution.Task.Id, strings.Join(activeNodes, ","))
		isTaskMerged = len(tasks) == 0
	}
	execution.IsMerged = isSubProcessMerged && isTaskMerged
}

//能否驳回
func CanRejected(currentNode INodeModel, parentNode INodeModel) bool {
	switch parentNode.(type) {
	case *TaskModel:
		if parentNode.(*TaskModel).PerformType == PT_ANY {
			return false
		}
	default:
	}
	result := false
	for _, tm := range currentNode.GetInputs() {
		source := tm.Source
		if source == parentNode {
			return true
		} else {
			switch source.(type) {
			case *ForkModel:
				continue
			case *JoinModel:
				continue
			case *SubProcessModel:
				continue
			case *StartModel:
				continue
			default:
			}
			result = result || CanRejected(source, parentNode)
		}
	}
	return result
}
