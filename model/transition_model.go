package model

type TransitionModel struct {
	BaseModel
	Source  NodeModeler ``                //变迁的源节点引用
	Target  NodeModeler ``                //变迁的目标节点引用
	To      string      `xml:"to,attr"`   //变迁的目的节点
	Expr    string      `xml:"expr,attr"` //变迁的条件表达式，用于decision
	Enabled bool        ``                //当前变迁路径是否可用
}
