package selectstatement

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
)

type sqlJoinClauseVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}
type sqlJoinVisitorImpl struct {
	parser.SqlVisitor
}

func (s *sqlJoinClauseVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.JoinClauseContext:
		arr := val.Accept(s).([][]interface{})
		var joinClause []*NodeJoinClause
		for _, v := range arr {
			joinClause = append(joinClause, &NodeJoinClause{
				JoinType: v[0].(JoinType),
				TableSource: &NodeTableSource{
					Source: v[1].(*Source),
					Alias:  v[2].(string),
				},
				OnExpress: v[3].(*NodeCondition),
			})
		}
		return joinClause
	default:
		panic("not impl")
	}
}
func (s *sqlJoinClauseVisitorImpl) VisitJoinClause(ctx *parser.JoinClauseContext) interface{} {
	res := [][]interface{}{}
	for k := range ctx.AllJoin() {
		var joinType JoinType
		joinType = Join
		if ctx.AllJoin()[k].LEFT() != nil {
			joinType = LeftJoin
		}
		if ctx.AllJoin()[k].RIGHT() != nil {
			joinType = RightJoin
		}
		if ctx.AllJoin()[k].INNER() != nil {
			joinType = InnerJoin
		}
		var aliasName string
		if len(ctx.AllAlias()) > 0 {
			aliasName = ctx.AllAlias()[k].GetText()
		}
		tmp := []interface{}{
			joinType,
			new(sqlTableSourceVisitorImpl).Visit(ctx.AllTableSource()[k]),
			aliasName,
			new(SqlLogicExpressionVisitorImpl).Visit(ctx.AllLogicExpression()[k]),
		}
		res = append(res, tmp)
	}
	return res
}
func (s *sqlJoinVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.JoinContext:
		joinType := val.Accept(s).(JoinType)
		return joinType
	default:
		panic("not impl")
	}
}
func (s *sqlJoinVisitorImpl) VisitJoin(ctx *parser.JoinContext) interface{} {
	return ctx.JOIN().Accept(s)
}
