package selectelements

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// type TreeShapeListener struct {
// 	*BaseSqlListener
// }

// func NewTreeShapeListener() *TreeShapeListener {
// 	return new(TreeShapeListener)
// }

//	func (*TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//		// fmt.Println(ctx.ToStringTree)
//	}
func TestEasy(t *testing.T) {
	//表达式原样返回
	var sql string
	// sql = `*`
	sql = `姓名,a.名字,a.字段1*a.字段2||a.字段3 AS 流水号/*注释*/,字段2 AS 名称号,字段3 别名3,a.姓名 AS 别名4,/*注释*/'b.姓名.信息' AS 别名5/*注释*/`
	// sql = `a.字段1 AS 流水号`
	// sql = `'啊啊' AS 流水号`
	sql = `*,a.*,'啊啊' AS 流水号`
	// sql = `a.*,b.*`
	// sql = `*`
	node := ParserNode(sql)
	spew.Dump(node)
	// spew.Dump(node.ColumnName (nil, "wholesql", nil))
	println("========================")
	println(SelectElementsString(node))

	// println("——————————————————————————")
	// is := antlr.NewInputStream("a.字段1||a.字段2 AS 流水号,substr(字段2||'附,加',1,2)/*ZZ,ZZ*/ AS 计算值")
	// lexer := parser.NewSqlLexer(is)
	// // Read all tokens
	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }

}
