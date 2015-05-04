package goflow

//XML节点基本模型
type BaseModel struct {
	Name        string `xml:"name,attr"`        //节点名称
	DisplayName string `xml:"displayName,attr"` //节点显示名称
}

func (p *BaseModel) GetName() string {
	return p.Name
}
