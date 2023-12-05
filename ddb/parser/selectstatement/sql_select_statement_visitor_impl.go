package selectstatement

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
	"github.com/linlexing/dbx/ddb/parser/model"
	"github.com/linlexing/dbx/ddb/parser/selectelements"
)

type sqlSelectStatementVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func (s *sqlSelectStatementVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.SelectStatementContext:
		node := val.Accept(s).(*model.NodeSelectStatement)
		return node
	default:
		panic("not impl")
	}
}
func (s *sqlSelectStatementVisitorImpl) VisitSelectStatement(ctx *parser.SelectStatementContext) interface{} {
	var nodeSelectElements *model.NodeSelectelements
	var nodeTableSources []*model.NodeTableSource
	var nodeJoinClause []*model.NodeJoinClause
	var nodeWhereClause *condition.Node
	if ctx.SelectElements() != nil {
		//SelectElements没注释
		nodeSelectElements = selectelements.ParseByContext(ctx.SelectElements())
	}
	if ctx.TableSources() != nil {
		nodeTableSources = parseByTableSourcesContext(ctx.TableSources(), s.vars)
	}
	if ctx.JoinClause() != nil {
		nodeJoinClause = parseByJoinContext(ctx.JoinClause(), s.vars)
	}
	if ctx.WhereClause() != nil {
		nodeWhereClause = condition.ParseByContext(ctx.WhereClause(), s.vars)
	}
	return &model.NodeSelectStatement{
		SelectElements: nodeSelectElements,
		TableSources:   nodeTableSources,
		JoinClause:     nodeJoinClause,
		WhereClause:    nodeWhereClause,
	}
}
