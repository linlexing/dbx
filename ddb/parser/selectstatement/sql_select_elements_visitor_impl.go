package selectstatement

import (
	"fmt"
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
)

var (
	//按.分隔，单双引号、括号内的不算
	regPoint = regexp.MustCompile(`(\([^)]*\)|'[^']*'|"[^"]*"|[^.()'"]+)+`)
	//括号括起的
	reBracket = regexp.MustCompile(`^\(.+\)$`)
)

type sqlSelectelementsVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func (s *sqlSelectelementsVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.SelectElementsContext:
		arr := val.Accept(s).([]interface{})
		var eles []*Element
		for _, v := range arr {
			eles = append(eles, v.(*Element))
		}
		return &NodeSelectelements{Elements: eles}
	default:
		panic("not impl")
	}
}
func (s *sqlSelectelementsVisitorImpl) VisitSelectElements(ctx *parser.SelectElementsContext) interface{} {
	var res []interface{}
	for k := range ctx.AllSelectElement() {
		res = append(res, ctx.SelectElement(k).Accept(s))
	}
	return res
}

func (s *sqlSelectelementsVisitorImpl) VisitSelectElement(ctx *parser.SelectElementContext) interface{} {
	var tableAlias, columnName, exprStr, asStr, aliaStr string
	var subquery *NodeSelectStatement
	expr, as, alias := ctx.Expr(), ctx.AS(), ctx.Alias()
	//判断有无括号
	bracket := reBracket.MatchString(expr.GetText())
	if bracket {
		expr = expr.Expr(0)
	}
	if expr != nil {
		//分隔a.xx，识别表别名
		if expr.ColumnName() != nil {
			col := expr.ColumnName().GetText()
			result := regPoint.FindAllString(col, -1)
			if len(result) == 1 {
				columnName = result[0]
			}
			if len(result) == 2 {
				tableAlias = result[0]
				columnName = result[1]
			}
		} else {
			exprStr = expr.GetText()
			if expr.CASE() != nil {
				exprStr = ""
				if len(expr.AllLogicExpression()) > 0 {
					exprStr = "CASE "
					for k := range expr.AllExpr() {
						var addStr string
						if len(expr.AllLogicExpression()) > k {
							addStr = fmt.Sprintf("WHEN %s THEN %s ",
								//这里不深入视图列表、关联表查询了
								expr.AllLogicExpression()[k].Accept(new(SqlLogicExpressionVisitorImpl)).(*NodeCondition).WhereString(nil, "wholesql", nil, true),
								expr.Expr(k).GetText())
						}
						//有ELSE最后一个给ELSE
						if expr.ELSE() != nil && k == len(expr.AllExpr())-1 {
							addStr = fmt.Sprintf("ELSE %s ", expr.Expr(k).GetText())
						}
						exprStr += addStr
					}
					exprStr += "END"
				} else {
					for k := range expr.AllExpr() {
						var addStr string
						if k == 0 {
							addStr = fmt.Sprintf("(CASE %s ", expr.Expr(0).GetText())
						}
						if k > 0 {
							if (k & 1) == 1 { //奇
								addStr = fmt.Sprintf("WHEN %s ", expr.Expr(k).GetText())
							} else { //偶
								addStr = fmt.Sprintf("THEN %s ", expr.Expr(k).GetText())
							}
							//有ELSE最后一个给ELSE
							if expr.ELSE() != nil && k == len(expr.AllExpr())-1 {
								addStr = fmt.Sprintf("ELSE %s ", expr.Expr(k).GetText())
							}
						}
						exprStr += addStr
					}
					exprStr += "END"
				}
			}
			if expr.SelectStatement() != nil {
				subquery = parseBySelectStatementContext(expr.SelectStatement(), s.vars)
			}
		}
	}
	if bracket && len(exprStr) > 0 {
		exprStr = "(" + exprStr + ")"
	}
	if as != nil {
		asStr = as.GetText()
	}
	if alias != nil {
		aliaStr = alias.GetText()
	}
	return &Element{
		TableAlias: tableAlias,
		ColumnName: columnName,
		Express:    exprStr,
		Subquery:   subquery,
		As:         asStr,
		Alias:      aliaStr,
	}
}
