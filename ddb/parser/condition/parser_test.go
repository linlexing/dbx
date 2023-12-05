package condition

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
	// node := ParserNode(`length(abc)/*注释*/>0/*注释*/`)
	node := ParserNode(`/*EXISTS(from e$自动化流程管理 on 名称=wholesql.名称 where )*/
	(EXISTS(select 1 from (select 名称,类别,归属方式,创建用户,归属部门,最后修改时间,流水号 from (
	select
	  id as 流水号,
	  name as 名称,
	  category as 类别,
	  ownerby as 归属方式,
	  username as 创建用户,
	  dept as 归属部门,
	  lasttime as 最后修改时间
	from
	  dataflow
	where
	  (
		ownerby = 'd'
		and dept = {{P .User.Dept.Code}}
	  )
	  or (
		ownerby = 'u'
		and username = {{P .User.Name}}
	  )
	) wholesql) exists_inner0 where 名称=wholesql.名称))`)

	spew.Dump(node)
	spew.Dump(node.ConditionLines(nil, "wholesql", nil))
	println("========================")
	println(node.WhereString(nil, "wholesql", nil, true))
}
func TestMain(t *testing.T) {
	t.Log("start")
	// node := ParserNode(`a=1 and /*PLAINTEXT*/
	// ((id=123) and
	// aaa=dd or
	// ccc=aa or
	// (
	// 	b=1 and
	// 	c=1 or
	// 	c in (select a.a from a
	// 		left join  b on a.a=b.a
	// 	)
	// )) and
	// /*COUNT(from ABC on abc=O.abc and where=O.aaa where a=1 and b=c exist(select 1 from aab where 1=2)) > 0*/
	// (COUNT(select 1 from ABC where abc=O.abc and where=O.aaa and a=1 and b=c and exist(select 1 from aab where 1=2)) > 0) AND
	// /*NOT IN(field in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2))*/
	// (field not in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2)) AND
	// /*PLAINTEXT*/
	// (not(EXISTS(select 1 from b
	// 	left join  b on a.a=b.a))) and
	// d=1 and
	// e=1 and
	// f=1`)
	node := ParserNode(`(
        a = '1' AND
        /*PLAINTEXT*/
        ((id=123) and
        aaa=dd or
        ccc=aa or
        (
			b=1 and
			c=1 or
			c in (select a.a from a	left join  b on a.a=b.a
			)
        )) AND
        /*NOT EXISTS(from ABC on abc=wholesql.abc and where=wholesql.aaa where a=1 and b=c exist(select 1 from aab where 1=2))*/
        (NOT EXISTS(select 1 from ABC where abc=wholesql.abc and where=wholesql.aaa and a=1 and b=c exist(select 1 from aab where 1=2))) AND
        /*COUNT(from ABC on abc=wholesql.abc and where=wholesql.aaa where a=1 and b=c exist(select 1 from aab where 1=2)) > 0*/
        ((select count(*) from ABC where abc=wholesql.abc and
where=wholesql.aaa) > 0) AND
        /*NOT IN(field in ABC(acd) where a=1 and b=c and exist(select 1 from aab where 1=2)*/
        (field not in (select 1 from ABC where where a=1 and b=c and exist(select 1 from aab where 1=2))) AND
        /*PLAINTEXT*/
        (not(EXISTS(select 1 from b left join  b on a.a=b.a))) AND
        d = '1' AND
        e = '1' AND
        f = '1'
)`)
	spew.Dump(node)
	spew.Dump(node.ConditionLines(nil, "wholesql", nil))
	println("========================")
	println(node.WhereString(nil, "wholesql", nil, true))
	// t.Log(tree.ToStringTree(nil, p))
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
