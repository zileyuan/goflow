package goflow

//任务状态
type TASK_STATUS int

const (
	TS_CLOSED   TASK_STATUS = iota //关闭
	TS_ACTIVITY                    //活动
)

//流程可用的开关,包含Process、Order
type FLOW_STATUS int

const (
	FS_FINISH      FLOW_STATUS = iota //结束状态
	FS_ACTIVITY                       //活动状态
	FS_TERMINATION                    //终止状态
)

type SURROGATE_STATUS int

const (
	SS_DISABLE SURROGATE_STATUS = iota //不可用
	SS_ENABLE                          //可用
)

//任务类型
type TASK_TYPE int

const (
	TT_MAJOR  TASK_TYPE = iota //主办任务
	TT_ASSIST                  //协办任务
	TT_RECORD                  //仅仅记录
)

//任务参与方式
type PERFORM_TYPE int

const (
	PT_ANY PERFORM_TYPE = iota //普通任务，即：任何一个参与者处理完即执行下一步
	PT_ALL                     //会签任务，即：所有参与者都完成，才可执行下一步
)
