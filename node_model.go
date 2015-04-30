package goflow

import "strings"

type NodeModel struct {
	BaseModel
	Inputs  []*TransitionModel `xml:"-"`          //输入变迁集合
	Outputs []*TransitionModel `xml:"transition"` //输出变迁集合
}

func (p *NodeModel) GetInputs() []*TransitionModel {
	return p.Inputs
}

func (p *NodeModel) GetOutputs() []*TransitionModel {
	return p.Outputs
}

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

func (p *NodeModel) MergeHandle(execution *Execution, activeNodes []string) error {
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
