package condition

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/pageselect"
)

//SqlWhereVisitorImpl 完成条件串的转换，只支持简单的 字段名 运算符 值 条件
type SqlWhereVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

//去除单引号，如果有的话
func decodeSignStringIf(str string) string {
	if len(str) < 2 {
		return str
	}
	if str[0] == '\'' && str[len(str)-1] == '\'' {
		return strings.ReplaceAll(str[1:len(str)-1], "''", "'")
	}
	return str
}
func NewSqlWhereVisitorImpl() *SqlWhereVisitorImpl {
	return &SqlWhereVisitorImpl{
		SqlVisitor: &parser.BaseSqlVisitor{},
		vars:       make(map[string]interface{}),
	}
}

func (s *SqlWhereVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.WhereClauseContext:
		node := val.Accept(s).(*Node)
		node.reduction()
		return node
	default:
		panic("not impl")
	}
}
func (s *SqlWhereVisitorImpl) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return ctx.LogicExpression().Accept(s)
}
func isColumn(expr parser.IExprContext) bool {
	return expr.(*parser.ExprContext).ColumnName() != nil
}
func (s *SqlWhereVisitorImpl) VisitLogicExpression(ctx *parser.LogicExpressionContext) interface{} {
	//逻辑关系隔开的条件
	if logicExpression1, logicalOperator, logicExpression2 :=
		ctx.LogicExpression(0), ctx.GetLogicalOperator(), ctx.LogicExpression(1); logicExpression1 != nil && logicalOperator != nil && logicExpression2 != nil {
		var nodeType NodeType
		if ctx.AND() != nil {
			nodeType = NodeAnd
		} else {
			nodeType = NodeOr
		}
		return NewLogicNode(nodeType, []*Node{logicExpression1.Accept(s).(*Node),
			logicExpression2.Accept(s).(*Node)})
	}
	//运算符隔开的单个条件
	if expr1, operate, expr2 := ctx.Expr(0), ctx.ComparisonOperator(), ctx.Expr(1); expr1 != nil && operate != nil && expr2 != nil {
		if funcCall := expr1.(*parser.ExprContext).FunctionCall(); funcCall != nil {

			switch tv := funcCall.(*parser.FunctionCallContext).CommonFunction().(type) {
			case *parser.CommonFunctionContext:
				// tv := fc.(*parser.CommonFunctionContext)
				if strings.ToUpper(tv.FunctionName().GetText()) == "LENGTH" {
					if exprList := tv.FunctionArg().(*parser.FunctionArgContext).AllExpr(); len(exprList) == 1 &&
						isColumn(exprList[0]) {
						var ope pageselect.Operator
						switch operate.GetText() {
						case "=":
							ope = pageselect.OperatorLengthEqu
						case ">":
							ope = pageselect.OperatorLengthGreaterThan
						case "<":
							ope = pageselect.OperatorLengthLessThan
						case "<=":
							ope = pageselect.OperatorLengthLessThanOrEqu
						case ">=":
							ope = pageselect.OperatorLengthGreaterThanOrEqu
						case "<>":
							ope = pageselect.OperatorLengthNotEqu
						default:
							panic("invalid length opereate " + operate.GetText())
						}
						return NewConditionNode(tv.FunctionArg().GetText(), ope, decodeSignStringIf(expr2.GetText()))
					}
				}
				return NewPlainNode(getText(ctx))

			}
		}
		if !isColumn(expr1) {
			return NewPlainNode(getText(ctx))
		}
		var ope pageselect.Operator
		switch operate.GetText() {
		case "=":
			ope = pageselect.OperatorEqu
		case ">":
			ope = pageselect.OperatorGreaterThan
		case "<":
			ope = pageselect.OperatorLessThan
		case "<=":
			ope = pageselect.OperatorLessThanOrEqu
		case ">=":
			ope = pageselect.OperatorGreaterThanOrEqu
		case "<>":
			ope = pageselect.OperatorNotEqu
		case "~":
			ope = pageselect.OperatorRegexp
		case "!~":
			ope = pageselect.OperatorNotRegexp
		default:
			panic("invalid opereate " + operate.GetText())
		}
		return NewConditionNode(expr1.GetText(), ope, decodeSignStringIf(expr2.GetText()))

	}
	//BETWEEN
	if expr1, between, expr2, expr3 :=
		ctx.Expr(0), ctx.BETWEEN(), ctx.Expr(1), ctx.Expr(2); expr1 != nil &&
		between != nil && expr2 != nil && expr3 != nil {

		return NewPlainNode(getText(ctx))
	}
	//IN/NOT IN
	if not, in, expr := ctx.NOT(), ctx.IN(), ctx.AllExpr(); in != nil && len(expr) > 2 {
		if !isColumn(expr[0]) {
			return NewPlainNode(getText(ctx))
		}
		var ope pageselect.Operator
		if not != nil {
			ope = pageselect.OperatorNotIn
		} else {
			ope = pageselect.OperatorIn
		}
		strs := []string{}
		//需要将in后面的表达式进行字面量的转换，去掉引号
		for _, one := range expr[1:] {
			strs = append(strs, decodeSignStringIf(one.GetText()))
		}
		bys := bytes.NewBufferString("")
		r := csv.NewWriter(bys)
		if err := r.WriteAll([][]string{strs}); err != nil {
			panic(err)
		}
		return NewConditionNode(expr[0].GetText(), ope, bys.String())

	}
	//LIKE/NOT LIKE
	if not, like, field, val :=
		ctx.NOT(), ctx.LIKE(), ctx.Expr(0), ctx.Expr(1); like != nil &&
		field != nil && val != nil {
		if !isColumn(field) {
			return NewPlainNode(getText(ctx))
		}
		str := decodeSignStringIf(val.GetText())
		var first, last byte
		if len(str) > 0 {
			first = str[0]
		}
		if len(str) > 1 {
			last = str[len(str)-1]
		}

		var ope pageselect.Operator
		var valStr string
		if not != nil {
			if first == '%' && last == '%' {
				ope = pageselect.OperatorNotLike
				valStr = str[1 : len(str)-1]
			} else if first == '%' {
				ope = pageselect.OperatorNotSuffix
				valStr = str[1:]
			} else if last == '%' {
				ope = pageselect.OperatorNotPrefix
				valStr = str[:len(str)-1]
			} else {
				ope = pageselect.OperatorNotEqu
				valStr = str
			}
		} else {
			if first == '%' && last == '%' {
				ope = pageselect.OperatorLike
				valStr = str[1 : len(str)-1]
			} else if first == '%' {
				ope = pageselect.OperatorSuffix
				valStr = str[1:]
			} else if last == '%' {
				ope = pageselect.OperatorPrefix
				valStr = str[:len(str)-1]
			} else {
				ope = pageselect.OperatorEqu
				valStr = str
			}
		}
		return NewConditionNode(field.GetText(), ope, valStr)
	}
	//IS NULL/IS NOT NULL
	if is, not, null, field :=
		ctx.IS(), ctx.NOT(), ctx.NULL(), ctx.Expr(0); is != nil && null != nil && field != nil {
		if !isColumn(field) {
			return NewPlainNode(getText(ctx))
		}
		if not != nil {
			return NewConditionNode(field.GetText(), pageselect.OperatorIsNotNull, "")
		}
		return NewConditionNode(field.GetText(), pageselect.OperatorIsNull, "")

	}
	//'(' logicExpression ')'
	if left, right, logicExpr := ctx.GetLeftBracket(), ctx.GetRightBracket(), ctx.LogicExpression(0); left != nil &&
		right != nil && logicExpr != nil {
		return logicExpr.Accept(s)
	}
	return NewPlainNode(getText(ctx))
	// panic("不支持的表达式" + ctx.GetText())
}
func getText(node antlr.ParseTree) string {
	if node.GetChildCount() == 0 {
		return node.GetText()
	}
	list := []string{}
	for _, one := range node.GetChildren() {
		list = append(list, getText(one.(antlr.ParseTree)))
	}
	return strings.Join(list, " ")
}
