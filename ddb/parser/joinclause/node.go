package joinclause

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

func ParserNode(val string) []*model.NodeJoinClause {
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
	tree := p.JoinClause()
	visitor := new(SqlJoinClauseVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).([]*model.NodeJoinClause)
}

func ParseByContext(ctx parser.IJoinClauseContext) []*model.NodeJoinClause {
	visitor := new(SqlJoinClauseVisitorImpl)
	return visitor.Visit(ctx).([]*model.NodeJoinClause)
}
