package logicexpression

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
)

type SqlLogicExpressionVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func NewSqlLogicExpressionVisitorImpl() *SqlLogicExpressionVisitorImpl {
	return &SqlLogicExpressionVisitorImpl{
		SqlVisitor: &parser.BaseSqlVisitor{},
		vars:       make(map[string]interface{}),
	}
}

func (s *SqlLogicExpressionVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.LogicExpressionContext:
		node := val.Accept(s).(*condition.Node)
		node.Reduction()
		return node
	default:
		panic("not impl")
	}
}
func (s *SqlLogicExpressionVisitorImpl) VisitLogicExpression(ctx *parser.LogicExpressionContext) interface{} {
	return condition.ParseLogicExpression(s, ctx, s.vars)
}
