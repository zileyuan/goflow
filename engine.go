package goflow

import (
	"time"

	"github.com/go-xorm/xorm"
)

//GFLOW数据流引擎
type Engine struct {
	ProcessService //流程定义业务类
}

//通过流程ID开始实例
func (p *Engine) StartInstanceById(id string, operator string, args map[string]interface{}) *Order {
	process := p.GetProcessById(id)
	return p.StartProcess(process, operator, args)
}

//通过流程NAME开始实例
func (p *Engine) StartInstanceByName(name string, version int, operator string, args map[string]interface{}) *Order {
	process := p.GetProcessByVersion(name, version)
	return p.StartProcess(process, operator, args)
}

//通过执行体Execution开始实例
func (p *Engine) StartInstanceByExecution(execution *Execution) *Order {
	process := execution.Process
	start := process.Model.GetStart()
	current := p.ExecuteByProcess(process, execution.Operator, execution.Args,
		execution.ParentOrder.Id, execution.ParentNodeName)
	start.Execute(current)
	return current.Order
}

//开始流程
func (p *Engine) StartProcess(process *Process, operator string, args map[string]interface{}) *Order {
	execution := p.ExecuteByProcess(process, operator, args, "", "")
	if process.Model != nil {
		start := process.Model.GetStart()
		start.Execute(execution)
	}
	return execution.Order
}

//执行流程
func (p *Engine) ExecuteByProcess(process *Process, operator string, args map[string]interface{},
	parentId string, parentNodeName string) *Execution {
	order := CreateOrder(process, operator, args, parentId, parentNodeName)
	execution := &Execution{
		Engine:   p,
		Process:  process,
		Order:    order,
		Operator: operator,
		Args:     args,
	}
	return execution
}

//通过任务ID，执行任务
func (p *Engine) GetExecutionByTaskId(id string, operator string, args map[string]interface{}) *Execution {
	task := CompleteTask(id, operator, args)

	order := &Order{}
	if order.GetOrderById(task.OrderId) {
		order.LastUpdator = operator
		order.LastUpdateTime = time.Now()
		if task.TaskType == TO_ASSIST { //协办任务完成不产生执行对象
			return nil
		} else {
			variable := JsonToMap(order.Variable)
			for k, v := range variable {
				if _, ok := args[k]; !ok { //判断 key 是否存在
					args[k] = v
				}
			}

			process := &Process{}
			if process.GetProcessById(order.ProcessId) {

				execution := &Execution{
					Engine:   p,
					Process:  process,
					Order:    order,
					Operator: operator,
					Task:     task,
					Args:     args,
				}
				return execution
			} else {
				return nil
			}
		}
	} else {
		return nil
	}
}

//执行并且跳到某个任务
func (p *Engine) ExecuteTask(id string, operator string, args map[string]interface{}) []*Task {
	execution := p.GetExecutionByTaskId(id, operator, args)
	if execution == nil {
		return nil
	}
	processModel := execution.Process.Model
	if processModel != nil {
		nodeModel := processModel.GetNode(execution.Task.TaskName)
		//将执行对象交给该任务对应的节点模型执行
		nodeModel.Execute(execution)
	}
	return execution.Tasks
}

//执行并且跳到某个任务
func (p *Engine) ExecuteAndJumpTask(id string, operator string, args map[string]interface{}, nodeName string) []*Task {
	execution := p.GetExecutionByTaskId(id, operator, args)
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

//根据流程实例ID，操作人ID，参数列表按照节点模型model创建新的自由任务
func (p *Engine) CreateFreeTask(orderId string, operator string, args map[string]interface{}, model *TaskModel) []*Task {
	order := &Order{}
	if order.GetOrderById(orderId) {
		order.LastUpdator = operator
		order.LastUpdateTime = time.Now()

		process := p.GetProcessById(order.Id)
		execution := &Execution{
			Engine:   p,
			Process:  process,
			Order:    order,
			Args:     args,
			Operator: operator,
		}
		return CreateTask(model, execution)
	}
	return nil
}

//新建引擎
func NewEngineByConfig(cfg string) *Engine {
	InitAccessByConfig(cfg)
	engine := &Engine{}
	engine.InitProcessService()
	return engine
}

//新建引擎
func NewEngineByXorm(orm *xorm.Engine) *Engine {
	InitAccessByXorm(orm)
	engine := &Engine{}
	engine.InitProcessService()
	return engine
}
