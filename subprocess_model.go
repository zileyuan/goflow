package goflow

//流程定义实体类
type SubProcessModel struct {
	NodeModel
	processName string        `xml:"assignee,attr"` //子流程名称
	Version     int           `xml:"version,attr"`  //子流程版本号
	SubProcess  *ProcessModel `xml:"-"`             //子流程定义引用
}
