package condition

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linlexing/dbx/ddb/parser"
)

// type TreeShapeListener struct {
// 	*BaseSqlListener
// }

// func NewTreeShapeListener() *TreeShapeListener {
// 	return new(TreeShapeListener)
// }

// func (*TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	// fmt.Println(ctx.ToStringTree)
// }
func TestMain(t *testing.T) {
	t.Log("start")
	stream := antlr.NewInputStream(`where 
	a=1 and 
	b=1 and
	c=1 or
	d=1 and
	e=1 and
	f=1`)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)

	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.WhereClause()
	visitor := new(SqlWhereVisitorImpl)
	node := visitor.Visit(tree)
	spew.Dump(node)
	println("========================")
	println(node.(*Node).WhereString())
	// t.Log(tree.ToStringTree(nil, p))
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
