package goflow

import (
	"github.com/lunny/log"
	"goflow/core"
	"os"
)

type Engine struct {
	//configuration  cfg.Configuration   //配置对象
	processService core.ProcessService //流程定义业务类
	orderService   core.OrderService   //流程实例业务类
	taskService    core.TaskService    //任务业务类
	queryService   core.QueryService   //查询业务类
	managerService core.ManagerService //管理业务类
}

func init() {
	f, _ := os.Create("goflow.log")
	log.Std.SetOutput(f)
}
