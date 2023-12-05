package logicexpression

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
)

func parserNode(val string) *condition.Node {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	val, vars = condition.ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.LogicExpression()
	visitor := new(SqlLogicExpressionVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*condition.Node)

}
func ParseByContext(ctx parser.ILogicExpressionContext, vars map[string]interface{}) *condition.Node {
	visitor := new(SqlLogicExpressionVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).(*condition.Node)
}
