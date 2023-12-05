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
func TestSelectStatement(t *testing.T) {
	//函数表达式原样返回：substr(a.流水号,1,2)
	var sql string
	sql = `SELECT *,名称,a.*,b.*,/*注释*/substr(a.流水号,1,2),a.单位名称,b.state
	FROM (SELECT 流水号,单位名称 FROM 规上月度1) a
	LEFT JOIN 规上月度1_as b on a.流水号 = b.流水号
	INNER JOIN 规上月度2 c on a.流水号 = c.流水号
	WHERE (b.流水号 IS NULL OR b.流水号 = 'zz') AND a.流水号 LIKE 'qq%'/*注释*/
	AND /*EXISTS(from e$自动化流程管理 on 名称=wholesql.名称 where )*/
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

	sql = `SELECT substr(a.流水号,1,2),a.单位名称,b.state FROM 规上月度1 a LEFT JOIN 规上月度1_as b on a.流水号 = b.流水号
	INNER JOIN 规上月度2 c on a.流水号 = c.流水号
	WHERE (b.流水号 IS NULL OR b.流水号 = 'zz') AND a.流水号 LIKE 'qq%'`
	// sql = `SELECT *,a.*,期别,substr(a.流水号,1,2) FROM 月度1 a`
	// sql = `SELECT * FROM 月度1`
	node := ParserNode(sql)

	spew.Dump(node)
	println("========================")
	println(SelectStatementString(node, nil, "wholesql", nil, true))
}
func TestJoin(t *testing.T) {
	node := parserNodeJoin(`LEFT JOIN 规上月度1_as b on a.流水号 = b.流水号 INNER JOIN (select 流水号 from 学生表 cc) C on a.流水号 = C.流水号`)

	spew.Dump(node)
	println("========================")
	println(joinClauseString(node, nil, "wholesql", nil, true))
}
func TestTableSources(t *testing.T) {
	var sql = `规上月度1 a,规上月度1_as b`
	sql = `规上月度1 a,规上月度1_as b,(select 名称 from 学生表) c`
	nodes := parserNodeTableSources(sql)

	spew.Dump(nodes)
	println("========================")
	println(tableSourceString(nodes[0], nil, "wholesql", nil, true))
}
