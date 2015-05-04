package goflow

//工作节点元素，目前是Task和SubProcess的基础
type WorkModel struct {
	NodeModel
	Action string `xml:"action,attr"`
}
