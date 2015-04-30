package goflow

//流程定义实体类
type SubProcessModel struct {
	WorkModel
	ProcessName string        `xml:"processName,attr"` //子流程名称
	Version     int           `xml:"version,attr"`     //子流程版本号
	SubProcess  *ProcessModel `xml:"-"`                //子流程Model对象
}

func (p *SubProcessModel) Execute(execution *Execution) error {
	return p.RunOutTransition(execution)
}
