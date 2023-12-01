package selectstatement

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

func ParserNode(val string) *model.NodeSelectStatement {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	// val, vars = processComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectStatement()
	visitor := new(SqlSelectStatementVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*model.NodeSelectStatement)
}

func ParseByContext(ctx parser.ISelectStatementContext) *model.NodeSelectStatement {
	visitor := new(SqlSelectStatementVisitorImpl)
	return visitor.Visit(ctx).(*model.NodeSelectStatement)
}
