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
	sql = `姓名,a.名字,a.字段1*a.字段2||a.字段3 AS 流水号,字段2 AS 名称号,字段3 别名3,a.姓名 AS 别名4,'b.姓名.信息' AS 别名5/*注释*/`
	// sql = `a.字段1 AS 流水号`
	// sql = `'啊啊' AS 流水号`
	// sql = `*,'啊啊' AS 流水号`
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

// func TestMain(t *testing.T) {
// 	t.Log("start")
// 	// node := ParserNode(`a=1 and /*PLAINTEXT*/
// 	// ((id=123) and
// 	// aaa=dd or
// 	// ccc=aa or
// 	// (
// 	// 	b=1 and
// 	// 	c=1 or
// 	// 	c in (select a.a from a
// 	// 		left join  b on a.a=b.a
// 	// 	)
// 	// )) and
// 	// /*COUNT(from ABC on abc=O.abc and where=O.aaa where a=1 and b=c exist(select 1 from aab where 1=2)) > 0*/
// 	// (COUNT(select 1 from ABC where abc=O.abc and where=O.aaa and a=1 and b=c and exist(select 1 from aab where 1=2)) > 0) AND
// 	// /*NOT IN(field in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2))*/
// 	// (field not in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2)) AND
// 	// /*PLAINTEXT*/
// 	// (not(EXISTS(select 1 from b
// 	// 	left join  b on a.a=b.a))) and
// 	// d=1 and
// 	// e=1 and
// 	// f=1`)
// 	node := ParserNode(`(
//         a = '1' AND
//         /*PLAINTEXT*/
//         ((id=123) and
//         aaa=dd or
//         ccc=aa or
//         (
// 			b=1 and
// 			c=1 or
// 			c in (select a.a from a	left join  b on a.a=b.a
// 			)
//         )) AND
//         /*NOT EXISTS(from ABC on abc=wholesql.abc and where=wholesql.aaa where a=1 and b=c exist(select 1 from aab where 1=2))*/
//         (NOT EXISTS(select 1 from ABC where abc=wholesql.abc and where=wholesql.aaa and a=1 and b=c exist(select 1 from aab where 1=2))) AND
//         /*COUNT(from ABC on abc=wholesql.abc and where=wholesql.aaa where a=1 and b=c exist(select 1 from aab where 1=2)) > 0*/
//         ((select count(*) from ABC where abc=wholesql.abc and
// where=wholesql.aaa) > 0) AND
//         /*NOT IN(field in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2)*/
//         (field not in (select 1 from ABC where where a=1 and b=c and exist(select 1 from aab where 1=2))) AND
//         /*PLAINTEXT*/
//         (not(EXISTS(select 1 from b left join  b on a.a=b.a))) AND
//         d = '1' AND
//         e = '1' AND
//         f = '1'
// )`)
// 	spew.Dump(node)
// 	spew.Dump(node.ConditionLines(nil, "wholesql", nil))
// 	println("========================")
// 	println(node.WhereString(nil, "wholesql", nil, true))
// 	// t.Log(tree.ToStringTree(nil, p))
// 	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
// }
