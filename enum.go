package goflow

//特性的关键字
const (
	DEFAULT_SEPARATOR   = "."                                    //默认分割符
	DEFAULT_KEY_ACTOR   = "S_ACTOR"                              //指定的actor
	DEFAULT_START_ID    = "00000000-0000-0000-0000-000000000000" //start node
	ORDERNO_TIME_LAYOUT = "2006-01-02_15:04:05.000"              //OrderNo时间格式
	STD_TIME_LAYOUT     = "2006-01-02 15:04:05"                  //标准时间格式

)

//流程可用的开关,包含Process、Order、Task
type FLOW_STATUS int

const (
	FS_FINISH      FLOW_STATUS = iota //结束状态
	FS_ACTIVITY                       //活动状态
	FS_TERMINATION                    //终止状态
)

//委托代理的状态
type SURROGATE_STATUS int

const (
	SS_DISABLE SURROGATE_STATUS = iota //不可用
	SS_ENABLE                          //可用
)

//任务类型
type TASK_ORDER int

const (
	TO_MAJOR  TASK_ORDER = iota //主办任务
	TO_ASSIST                   //协办任务
)

//任务类型
type TASK_TYPE string

const (
	TT_MAJOR  TASK_TYPE = "MAJOR"  //主办任务
	TT_ASSIST TASK_TYPE = "ASSIST" //协办任务
)

//任务参与方式
type PERFORM_ORDER int

const (
	PO_ANY PERFORM_ORDER = iota //普通任务，即：任何一个参与者处理完即执行下一步
	PO_ALL                      //会签任务，即：所有参与者都完成，才可执行下一步
)

//任务参与方式
type PERFORM_TYPE string

const (
	PT_ANY PERFORM_TYPE = "ANY" //普通任务，即：任何一个参与者处理完即执行下一步
	PT_ALL PERFORM_TYPE = "ALL" //会签任务，即：所有参与者都完成，才可执行下一步
)

//执行任务的角色
type EXEC_ROLE string

const (
	ER_ADMIN   EXEC_ROLE = "ADMIN"   //管理员执行
	ER_AUTO    EXEC_ROLE = "AUTO"    //自动执行
	ER_ORDERNO EXEC_ROLE = "ORDERNO" //流程编号
)
