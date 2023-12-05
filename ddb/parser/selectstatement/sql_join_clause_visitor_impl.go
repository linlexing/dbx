package selectstatement

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
	"github.com/linlexing/dbx/ddb/parser/logicexpression"
	"github.com/linlexing/dbx/ddb/parser/model"
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
		var joinClause []*model.NodeJoinClause
		for _, v := range arr {
			joinClause = append(joinClause, &model.NodeJoinClause{
				JoinType: v[0].(model.JoinType),
				TableSource: &model.NodeTableSource{
					Source: v[1].(*model.Source),
					Alias:  v[2].(string),
				},
				OnExpress: v[3].(*condition.Node),
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
		var joinType model.JoinType
		joinType = model.Join
		if ctx.AllJoin()[k].LEFT() != nil {
			joinType = model.LeftJoin
		}
		if ctx.AllJoin()[k].RIGHT() != nil {
			joinType = model.RightJoin
		}
		if ctx.AllJoin()[k].INNER() != nil {
			joinType = model.InnerJoin
		}
		tmp := []interface{}{
			joinType,
			new(sqlTableSourceVisitorImpl).Visit(ctx.AllTableSource()[k]),
			ctx.AllAlias()[k].GetText(),
			new(logicexpression.SqlLogicExpressionVisitorImpl).Visit(ctx.AllLogicExpression()[k]),
		}
		res = append(res, tmp)
	}
	return res
}
func (s *sqlJoinVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.JoinContext:
		joinType := val.Accept(s).(model.JoinType)
		return joinType
	default:
		panic("not impl")
	}
}
func (s *sqlJoinVisitorImpl) VisitJoin(ctx *parser.JoinContext) interface{} {
	return ctx.JOIN().Accept(s)
}
