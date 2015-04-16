package model

type DecisionModel struct {
	NodeModel
	Expr string `xml:"expr,attr"` //变迁的条件表达式
}
