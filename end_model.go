package goflow

//XML的结束节点
type EndModel struct {
	NodeModel
}

func (p *EndModel) Exec(execution *Execution) {
	engine := execution.Engine
	order := execution.Order
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		CompleteTask(task.Id, string(ER_AUTO), nil)
	}
	CompleteOrder(order.Id)

	if order.ParentId == "" {
		parentOrder := &Order{}
		parentOrder.GetOrderById(order.ParentId)

		process := engine.GetProcessById(parentOrder.ProcessId)

		processModel := process.Model
		spm := processModel.GetNode(order.ParentNodeName).(*SubProcessModel)

		newExecution := &Execution{
			Engine:       engine,
			Process:      process,
			Order:        parentOrder,
			Args:         execution.Args,
			ChildOrderId: order.Id,
			Task:         execution.Task,
		}
		spm.Execute(newExecution)
		execution.Tasks = append(execution.Tasks, newExecution.Tasks...)
	}
}

//执行
func (p *EndModel) Execute(execution *Execution) {
	p.PrevIntercept(execution)
	p.Exec(execution)
	p.PostIntercept(execution)
}
