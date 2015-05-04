package goflow

//历史任务参与者
type HistoryTaskActor struct {
	Id      string `xorm:"varchar(36) pk notnull"`    //主键ID
	TaskId  string `xorm:"varchar(36) notnull index"` //任务ID
	ActorId string `xorm:"varchar(36) notnull"`       //参与者ID
}
