package goflow

import (
	"os"

	"github.com/lunny/log"
)

type Engine struct {
	processService *ProcessService //流程定义业务类
	orderService   *OrderService   //流程实例业务类
	taskService    *TaskService    //任务业务类
	queryService   *QueryService   //查询业务类
	managerService *ManagerService //管理业务类
}

func (p *Engine) StartInstanceById(id string, operator string) *Order {
	process := new(Process)
	process.GetProcessById(id)
	return p.StartProcess(process, operator)
}

func (p *Engine) StartProcess(process *Process, operator string) *Order {
	execution := p.Execute(process, operator)
	if process.Model != nil {
		start := process.Model.GetStart()
		start.Execute(execution)
	}
	return execution.Order
}

func (p *Engine) Execute(process *Process, operator string) *Execution {
	order := p.orderService.CreateOrder(process, operator)
	execution := &Execution{
		Engine:   p,
		Process:  process,
		Order:    order,
		Operator: operator,
	}
	return execution
}

func NewEngine() {

}

func init() {
	f, _ := os.Create("goflow.log")
	log.Std.SetOutput(f)
}
