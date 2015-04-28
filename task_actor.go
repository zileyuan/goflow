package goflow

//任务参与者
type TaskActor struct {
	TaskId  string `xorm:"varchar(48) notnull index(IDX_TASKACTOR_TASK)"` //任务ID
	ActorId string `xorm:"varchar(48) notnull"`                           //参与者ID
}
