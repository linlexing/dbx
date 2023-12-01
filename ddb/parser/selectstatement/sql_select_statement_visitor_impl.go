package selectstatement

import (
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
	"github.com/linlexing/dbx/ddb/parser/joinclause"
	"github.com/linlexing/dbx/ddb/parser/model"
	"github.com/linlexing/dbx/ddb/parser/selectelements"
	"github.com/linlexing/dbx/ddb/parser/tablesources"
)

var (
	regComment = regexp.MustCompile(`(?:[^']|'[^']*')*?(/\*[^*]*\*+(?:[^/*][^*]*\*+)*/)`)
)

type SqlSelectStatementVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func (s *SqlSelectStatementVisitorImpl) NewSqlSelectStatementVisitorImpl() *SqlSelectStatementVisitorImpl {
	return s
}

func (s *SqlSelectStatementVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.SelectStatementContext:
		node := val.Accept(s).(*model.NodeSelectStatement)
		return node
	default:
		panic("not impl")
	}
}
func (s *SqlSelectStatementVisitorImpl) VisitSelectStatement(ctx *parser.SelectStatementContext) interface{} {
	nodeSelectElements := selectelements.ParseByContext(ctx.SelectElements())
	nodeTableSources := tablesources.ParseByContext(ctx.TableSources())
	nodeJoinClause := joinclause.ParseByContext(ctx.JoinClause())
	nodeWhereClause := condition.ParseByContext(ctx.WhereClause())
	println(&model.NodeSelectStatement{
		SelectElements: nodeSelectElements,
		TableSources:   nodeTableSources,
		JoinClause:     nodeJoinClause,
		WhereClause:    nodeWhereClause,
	})
	return &model.NodeSelectStatement{
		SelectElements: nodeSelectElements,
		TableSources:   nodeTableSources,
		JoinClause:     nodeJoinClause,
		WhereClause:    nodeWhereClause,
	}

	// println(ctx.SELECT().GetText())
	// println(ctx.SelectElements().GetText())
	// println(ctx.FROM().GetText())
	// println(ctx.TableSources().GetText())
	// println(ctx.JoinClause().GetText())
	// println(ctx.JoinClause().AllJoin()[0].GetText())
	// println(ctx.JoinClause().AllTableSources()[0].GetText())
	// println(ctx.JoinClause().AllLogicExpression()[0].GetText())
	// println(ctx.JoinClause().AllJoin()[1].GetText())
	// println(ctx.JoinClause().AllTableSources()[1].GetText())
	// println(ctx.JoinClause().AllLogicExpression()[1].GetText())
	// // println(ctx.Join().GetText())
	// // println(ctx.ON().GetText())
	// // println(ctx.LogicExpression().GetText())
	// println(ctx.WhereClause().GetText())
	// println(ctx.GroupByClause().GetText())
	// println(ctx.HavingClause().GetText())
	// println(ctx.OrderByClause().GetText())
	// println(ctx.LimitClause().GetText())
	// println(ctx.AllSelectStatement())
	// println(ctx.SelectStatement(0).GetText()) //ctx.SelectStatement(i)
	// println(ctx.Union().GetText())
	// return nil
}

// func (s *SqlSelectStatementVisitorImpl) VisitSelectElement(ctx *parser.SelectElementContext) interface{} {
// 	//sql_parser里搜*SelectElementContext)，用那些方法，
// 	//（仿照where_visitor_impl的133）对应逻辑写
// 	//QQ AS WW
// 	// if ctx.GetText() == "*"{
// 	// 	return NewElement(exprStr, asStr, aliaStr)
// 	// }
// 	var tableAlias, columnName, exprStr, asStr, aliaStr string
// 	expr, as, alias := ctx.Expr(), ctx.AS(), ctx.Alias()
// 	if expr != nil {
// 		//分隔a.xx，识别表别名
// 		if expr.ColumnName() != nil {
// 			col := expr.ColumnName().GetText()
// 			result := regPoint.FindAllString(col, -1)
// 			if len(result) == 1 {
// 				columnName = result[0]
// 			}
// 			if len(result) == 2 {
// 				tableAlias = result[0]
// 				columnName = result[1]
// 			}
// 		} else {
// 			exprStr = expr.GetText()
// 		}
// 	}
// 	if as != nil {
// 		asStr = as.GetText()
// 	}
// 	if alias != nil {
// 		aliaStr = alias.GetText()
// 	}
// 	return NewElement(tableAlias, columnName, exprStr, asStr, aliaStr)
// }
