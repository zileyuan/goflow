package goflow

import (
	"os"
	"time"

	"github.com/lunny/log"
)

type Engine struct {
	ProcessService //流程定义业务类
	TaskService    //任务业务类
}

func (p *Engine) StartInstanceById(id string, operator string, args map[string]interface{}) *Order {
	process := new(Process)
	process.GetProcessById(id)
	return p.StartProcess(process, operator, args)
}

func (p *Engine) StartInstanceByName(name string, version int, operator string, args map[string]interface{}) *Order {
	process := p.GetProcessByVersion(name, version)
	return p.StartProcess(process, operator, args)
}

func (p *Engine) StartProcess(process *Process, operator string, args map[string]interface{}) *Order {
	execution := p.ExecuteByProcess(process, operator, args)
	if process.Model != nil {
		start := process.Model.GetStart()
		start.Execute(execution)
	}
	return execution.Order
}

func (p *Engine) ExecuteByProcess(process *Process, operator string, args map[string]interface{}) *Execution {
	order := CreateOrder(process, operator)
	execution := &Execution{
		Engine:   p,
		Process:  process,
		Order:    order,
		Operator: operator,
	}
	return execution
}

func (p *Engine) ExecuteByTaskId(id string, operator string, args map[string]interface{}) *Execution {
	task := CompleteTask(id, operator, args)

	order := &Order{}
	order.GetOrderById(task.OrderId)
	order.LastUpdator = operator
	order.LastUpdateTime = time.Now()
	if task.TaskType != TT_MAJOR {
		return nil
	} else {
		process := new(Process)
		process.GetProcessById(order.ProcessId)

		execution := &Execution{
			Engine:   p,
			Process:  process,
			Order:    order,
			Operator: operator,
			Task:     task,
		}
		return execution
	}
}

func (p *Engine) ExecuteAndJumpTask(id string, operator string, args map[string]interface{}, nodeName string) []*Task {
	execution := p.ExecuteByTaskId(id, operator, args)
	if execution != nil {
		model := execution.Process.Model
		if nodeName == "" {
			task := RejectTask(model, execution.Task)
			execution.Tasks = append(execution.Tasks, task)
		} else {
			nodeModel := model.GetNode(nodeName)
			tm := &TransitionModel{
				Target:  nodeModel,
				Enabled: true,
			}
			tm.Execute(execution)
		}
		return execution.Tasks
	}
	return []*Task{}
}

func NewEngine() *Engine {
	engine := &Engine{}
	engine.InitProcessService()
	return engine
}

func init() {
	f, _ := os.Create("goflow.log")
	log.Std.SetOutput(f)

	InitAccess()
}
