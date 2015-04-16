package entity

//历史任务参与者
type HistoryTaskActor struct {
	TaskId  string `xorm:"varchar(32) notnull index(IDX_HIST_TASKACTOR_TASK)"` //任务ID
	ActorId string `xorm:"varchar(50) notnull"`                                //参与者ID
}
