package tablesources

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

type SqlTableSourcesVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}
type statement interface {
	NewSqlSelectStatementVisitorImpl()
}

func (s *SqlTableSourcesVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.TableSourcesContext:
		m := val.Accept(s).(map[string]interface{})
		var tableSources []*model.NodeTableSource
		for k := range m {
			tableSources = append(tableSources, &model.NodeTableSource{
				Alias:  k,
				Source: m[k].(*model.Source),
			})
		}
		return tableSources
	default:
		panic("not impl")
	}
}
func (s *SqlTableSourcesVisitorImpl) VisitTableSources(ctx *parser.TableSourcesContext) interface{} {
	res := make(map[string]interface{})
	for k := range ctx.AllTableSource() {
		res[ctx.Alias(k).GetText()] = ctx.TableSource(k).Accept(s)
	}
	return res
}

func (s *SqlTableSourcesVisitorImpl) VisitTableSource(ctx *parser.TableSourceContext) interface{} {
	// visitor := new(selectstatement.SqlSelectStatementVisitorImpl)
	// var vars map[string]interface{}
	// visitor.vars = vars
	return &model.Source{
		TableName:       ctx.TableName().GetText(),
		SelectStatement: nil,
		// SelectStatement: visitor.Visit(ctx.SelectStatement()).(*model.NodeSelectStatement),
	}
}
