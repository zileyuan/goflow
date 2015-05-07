package goflow

import "github.com/Knetic/govaluate"

//XML决策节点
type DecisionModel struct {
	NodeModel
	Expr string `xml:"expr,attr"` //决策的条件表达式
}

//执行
func (p *DecisionModel) Exec(execution *Execution) {
	var next interface{}
	if p.Expr != "" {
		expression, _ := govaluate.NewEvaluableExpression(p.Expr)
		next, _ = expression.Evaluate(execution.Args)
	}

	for _, tm := range p.Outputs {
		if next == nil {
			expression, _ := govaluate.NewEvaluableExpression(tm.Expr)
			canflow, _ := expression.Evaluate(execution.Args)
			if canflow.(bool) {
				tm.Enabled = true
				tm.Execute(execution)
			}
		} else {
			if tm.Name == next.(string) {
				tm.Enabled = true
				tm.Execute(execution)
			}
		}
	}
}
