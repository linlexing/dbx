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
	stream := antlr.NewInputStream(`where not(a is not null or a BETWEEN 2 and 3) and  
	1='exists(select count(*) from ''bb'' where bbc.aa=a)' and a=2 and (b=c or 1=2) or 
	test||23 like '%aa' and c like 'aa%'`)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)

	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.WhereClause()
	visitor := new(SqlWhereVisitorImpl)
	spew.Dump(visitor.Visit(tree))
	// t.Log(tree.ToStringTree(nil, p))
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
