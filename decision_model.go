package goflow

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

//XML决策节点
type DecisionModel struct {
	NodeModel
	Expr string `xml:"expr,attr"` //决策的条件表达式
}

//执行
func (p *DecisionModel) Execute(execution *Execution) error {
	expression, _ := govaluate.NewEvaluableExpression(p.Expr)
	next, _ := expression.Evaluate(execution.Args)

	isfound := false
	for _, tm := range p.Outputs {
		if next == nil {
			expression, _ := govaluate.NewEvaluableExpression(tm.Expr)
			canflow, _ := expression.Evaluate(execution.Args)
			if canflow.(bool) {
				tm.Enabled = true
				tm.Execute(execution)
				isfound = true
			}
		} else {
			if tm.Name == next {
				tm.Enabled = true
				tm.Execute(execution)
				isfound = true
			}
		}
	}
	if isfound {
		return fmt.Errorf("decision no next step!")
	} else {
		return nil
	}
}
