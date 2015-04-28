package goflow

//历史任务参与者
type HistoryTaskActor struct {
	TaskId  string `xorm:"varchar(36) notnull index"` //任务ID
	ActorId string `xorm:"varchar(36) notnull"`       //参与者ID
}
