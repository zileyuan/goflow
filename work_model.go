package goflow

//工作节点元素，目前是Task和SubProcess的基础，同时也是业务数据的绑定点
type WorkModel struct {
	NodeModel
	Action string `xml:"action,attr"`
}
