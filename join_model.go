package goflow

type JoinModel struct {
	NodeModel
}

func (p *JoinModel) MergeBranchHandle(execution *Execution) error {
	activeNodes := FindActiveNodes(p)
	return MergeHandle(execution, activeNodes)
}

func (p *JoinModel) Execute(execution *Execution) error {
	p.MergeBranchHandle(execution)
	if execution.IsMerged {
		p.RunOutTransition(execution)
	}
	return nil
}

func FindForkTaskNames(node INodeModel) []string {
	ret := make([]string, 0)
	switch node.(type) {
	case *ForkModel:
	default:
		for _, tm := range node.GetInputs() {
			switch tm.Source.(type) {
			case *SubProcessModel:
				ret = append(ret, tm.Source.(*SubProcessModel).Name)
			case *TaskModel:
				ret = append(ret, tm.Source.(*TaskModel).Name)
			default:
				ret = append(ret, FindForkTaskNames(tm.Source)...)
			}
		}
	}
	return ret
}

func FindActiveNodes(node INodeModel) []string {
	return FindForkTaskNames(node)
}
