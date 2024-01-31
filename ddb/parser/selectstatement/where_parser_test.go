package selectstatement

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
	sql := `/*EXISTS(from e$自动化流程管理 on 名称=wholesql.名称 where )*/
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
	) wholesql) exists_inner0 where 名称=wholesql.名称))`
	sql = `((length(字段))) > 1`
	sql = `(case 字段 when '1' then 1 else 2 end) > 1`
	sql = `(select 数量 from 库存 where 名称 = '东西') = 2`
	sql = `(select 数量 from 库存 where 名称 = '东西') = (select 总数 from 仓库 where 产品 = '物品')`
	sql = `not (select 数量 from 库存 where 名称 = '东西') like 'q%'`
	// sql = `not (数量 > 2)`
	sql = `(select 数量 from 库存 where 名称 = '东西') in (1,2)`
	sql = `/*IN(名称 in e$sqlhis(名称) where )*/
	(名称 in (select 名称 from (select dbname,exetime,id,rowsaffected,sql,username,usedtime from (
	select
	  *
	from
	  sqlhis
	) wholesql) in_inner))`
	sql = `substr(总数,1,1) in (select 数量 from 库存 where 名称 = 'dsa')`
	sql = `(EXISTS(select name from 库存))`
	sql = `((select count(*) from (select name from 库存) cnt_inner0) = 2)`
	sql = `((select count(*) from (select dbname,exetime,id,rowsaffected,sql,username,usedtime from (
		select
		  *
		from
		  sqlhis
		) wholesql) cnt_inner0 where dbname=wholesql.名称) = 1)`
	node := ParserWhereNode(sql)

	spew.Dump(node)
	spew.Dump(node.ConditionLines(nil, "wholesql", nil, false))
	println("========================")
	println(node.WhereString(nil, "wholesql", nil, true, false))
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
	node := ParserWhereNode(`(
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
	spew.Dump(node.ConditionLines(nil, "wholesql", nil, false))
	println("========================")
	println(node.WhereString(nil, "wholesql", nil, true, false))
	// t.Log(tree.ToStringTree(nil, p))
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
func TestMultiTable(t *testing.T) {
	node := ParserWhereNode("select 1 from a where exists(select 1 from b where a.c=b.c)")
	spew.Dump(node)
	spew.Dump(node.ConditionLines(nil, "wholesql", nil, false))
	println("========================")
	println(node.WhereString(nil, "wholesql", nil, true, false))
}
