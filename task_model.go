package goflow

//XML流程定义的任务节点
type TaskModel struct {
	WorkModel
	Assignee    string       `xml:"assignee,attr"`    //参与者变量名称
	PerformType PERFORM_TYPE `xml:"performType,attr"` //参与方式
	TaskType    TASK_TYPE    `xml:"taskType,attr"`    //任务类型
	AutoExecute bool         `xml:"autoExecute,attr"` //是否自动执行
	ExpireTime  string       `xml:"expireTime,attr"`  //期望完成时间
}

//执行
func (p *TaskModel) Exec(execution *Execution) {
	if ProcessPerformType(p.PerformType) == PO_ANY {
		p.RunOutTransition(execution)
	} else {
		MergeActorHandle(p, execution)
		if execution.IsMerged {
			p.RunOutTransition(execution)
		}
	}
}

//根据任务节点创建任务对象
func CreateTaskHandle(tm *TaskModel, execution *Execution) []*Task {
	tasks := CreateTask(tm, execution)
	execution.Tasks = append(execution.Tasks, tasks...)

	return tasks
}

//自动运行需要自动运行的任务
func AutoExecuteTask(tm *TaskModel, execution *Execution, tasks []*Task) {
	if tm.AutoExecute {
		for _, task := range tasks {
			execution.Engine.ExecuteTask(task.Id, string(ER_AUTO), execution.Args)
		}
	}
}

//合并任务角色的处理
func MergeActorHandle(tm *TaskModel, execution *Execution) {
	activeNodes := []string{tm.Name}
	MergeHandle(execution, activeNodes)
}
