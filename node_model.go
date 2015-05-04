package goflow

import "strings"

//XML节点通用信息
type NodeModel struct {
	BaseModel
	Inputs  []*TransitionModel `xml:"-"`          //输入变迁集合
	Outputs []*TransitionModel `xml:"transition"` //输出变迁集合
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
func (p *NodeModel) RunOutTransition(execution *Execution) error {
	for _, tm := range p.Outputs {
		tm.Enabled = true
		err := tm.Execute(execution)
		if err != nil {
			return err
		}
	}
	return nil
}

//合并处理通用流程
func MergeHandle(execution *Execution, activeNodes []string) error {
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
	return nil
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
