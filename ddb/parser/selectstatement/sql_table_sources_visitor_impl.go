package selectstatement

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
)

type sqlTableSourcesVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}
type sqlTableSourceVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func (s *sqlTableSourcesVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.TableSourcesContext:
		m := val.Accept(s).(map[interface{}]string)
		var tableSources []*NodeTableSource
		for k := range m {
			tableSources = append(tableSources, &NodeTableSource{
				Alias:  m[k],
				Source: k.(*Source),
			})
		}
		return tableSources
	default:
		panic("not impl")
	}
}
func (s *sqlTableSourcesVisitorImpl) VisitTableSources(ctx *parser.TableSourcesContext) interface{} {
	res := make(map[interface{}]string)
	for k := range ctx.AllTableSource() {
		visitor := new(sqlTableSourceVisitorImpl)
		str := ""
		var intf interface{}
		if ctx.TableSource(k) != nil {
			intf = ctx.TableSource(k).Accept(visitor)
		}
		if ctx.Alias(k) != nil {
			str = ctx.Alias(k).GetText()
		}
		res[intf] = str
	}
	return res
}

func (s *sqlTableSourceVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.TableSourceContext:
		node := val.Accept(s).(*Source)
		return node
	default:
		panic("not impl")
	}
}

func (s *sqlTableSourceVisitorImpl) VisitTableSource(ctx *parser.TableSourceContext) interface{} {
	visitor := new(sqlSelectStatementVisitorImpl)
	var comment, tableName string
	var nodeSS *NodeSelectStatement
	if ctx.COMMENT() != nil {
		comment = strings.TrimPrefix(ctx.COMMENT().GetText(), "/*")
		comment = strings.TrimSuffix(comment, "*/")
	}
	if ctx.TableName() != nil {
		tableName = ctx.TableName().GetText()
	}
	if ctx.SelectStatement() != nil {
		nodeSS = visitor.Visit(ctx.SelectStatement()).(*NodeSelectStatement)
	}
	return &Source{
		Comment:         comment,
		TableName:       tableName,
		SelectStatement: nodeSS,
	}
}
