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

	sql = `SELECT 期别,专业 FROM 月度1 UNION ALL SELECT 期别,专业 FROM 月度2
	UNION ALL SELECT 期别,专业 FROM 月度3 UNION SELECT 期别,专业 FROM 月度4`
	sql = `SELECT *,a.*,期别,substr(a.流水号,1,2) FROM 月度1 a`
	sql = `SELECT
	'法人单位' AS 单位类别,
	a.组织机构代码,
	a.详细名称,
	b.门类代码 AS 行业门类代码,
	a.行业代码,
	a.行政区划代码,
	nvl(a.产业活动单位数, 0) AS 产业活动单位数,
	a.登记注册类型,
	a.开业年份,
	a.开业月份,
	a.成立年份,
	a.成立月份,
	a.企业营业状态,
	a.机构类型,
	a.企业控股情况,
	a.执行会计制度类别,
	a.状态,
	a.统计局代码,
	a.期末从业人数,
	a.女性人数,
	a.全年营业收入,
	a.主营收入,
	a.资产总计,
	a.非企业支出合计,
	a.年末资产,
	(case when 行业代码 >= '06' and 行业代码 <= '46' and
			   nvl(主营收入,0) >= 5000
			then 'B10'
		  when 行业代码 >= '47' and 行业代码 <= '50' and
			   nvl(建筑业资质等级,'0000') > 'A000' and
			   nvl(建筑业资质等级,'0000') <= 'C999'
			then 'C10'
		  when 行业代码 = '7210'
			then 'H10'
		  when (行业代码 like '63%' and nvl(主营收入,0) >=20000) or
			   (行业代码 like '65%' and nvl(主营收入,0) >=5000)
			then 'E11'
		  when (行业代码 like '66%' or 行业代码 like '67%') and
			   nvl(主营收入,0) >=2000
			then 'E21'
		  else null
	end) as 专业标识,
	a.最后修改日期,
	a.入库日期,
	a.变更类别,
	a.统计局代码 as 法人_统计局代码,
	a.工业年报,
	a.建筑业年报,
	a.房地产业年报,
	a.批零贸易业年报,
	a.住宿餐饮业年报,
	a.确认级别,
	a.确认专业,
	a.剔除原因,
	a.附加选项1,
	a.附加选项2,
	a.附加选项3,
	a.附加选项4,
	a.附加选项5,
	a.附加选项6,
	a.附加标记1,
	a.附加标记2,
	a.附加标记3,
	a.附加标记4,
	a.附加标记5,
	a.附加标记6
  from
	t$法人单位表 a left join 行业代码表 b on a.行业代码=b.代码
  where
	(a.状态 is null or a.状态<> '剔除')
  UNION ALL
  SELECT
	'产业活动单位' AS 单位类别,
	a.组织机构代码,
	a.详细名称,
	b.门类代码 AS 行业门类代码,
	a.行业代码,
	a.行政区划代码,
	0 AS 产业活动单位数,
	a.登记注册类型,
	a.开业年份,
	a.开业月份,
	a.成立年份,
	a.成立月份,
	a.企业营业状态,
	a.机构类型,
	null as 企业控股情况,
	null as 执行会计制度类别,
	a.状态,
	a.统计局代码,
	a.期末从业人数,
	0 as 女性人数,
	0 as 全年营业收入,
	a.经营性单位收入 as 主营收入,
	0 as 资产总计,
	a.非经营性单位支出 as 非企业支出合计,
	0 as 年末资产,
	(case  when (行业代码 like '63%' and nvl(经营性单位收入,0) >=20000) or
			   (行业代码 like '65%' and nvl(经营性单位收入,0) >=5000)
			then 'E11'
		   when (行业代码 like '66%' or 行业代码 like '67%') and
			   nvl(经营性单位收入,0) >=2000
			then 'E21'
		  else null
	end) as 专业标识,
	a.最后修改日期,
	a.入库日期,
	a.变更类别,
	(select min(统计局代码) from t$法人单位表 d where (d.状态 is null or d.状态 <> '剔除') and a.归属法人组织机构代码 = d.组织机构代码) as 法人_统计局代码,
	null as 工业年报,
	null as 建筑业年报,
	null as 房地产业年报,
	null as 批零贸易业年报,
	null as 住宿餐饮业年报,
	a.确认级别,
	a.确认专业,
	a.剔除原因,
	a.附加选项1,
	a.附加选项2,
	a.附加选项3,
	a.附加选项4,
	a.附加选项5,
	a.附加选项6,
	a.附加标记1,
	a.附加标记2,
	a.附加标记3,
	a.附加标记4,
	a.附加标记5,
	a.附加标记6
  from
	t$产业活动单位表 a LEFT JOIN 行业代码表 b ON a.行业代码 = b.代码
  where
	a.状态 is null or a.状态<> '剔除'`
	// sql = `select * from 表 where 字段1 > .123e+10`
	sql = `select * from /*注释1*/表1,/*注释2*/表2`
	sql = `SELECT *,名称,a.*,b.*,/*注释*/substr(a.流水号,1,2),a.单位名称,b.state
	FROM /*注释规上月度1*/ (SELECT 流水号,单位名称 FROM 规上月度1) a
	LEFT JOIN /*注释规上月度1as*/规上月度1_as b on a.流水号 = b.流水号
	INNER JOIN /*注释规上月度2*/ 规上月度2 c on a.流水号 = c.流水号
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
	//sql = `select * from 表1 a LEFT JOIN /*TableSources注释测试*/ 规上月度1_as b on a.流水号 = b.流水号`
	node := ParserSelectNode(sql)

	spew.Dump(node)
	println("========================")
	println(node.SelectStatementString(nil))
}
func TestJoin(t *testing.T) {
	node := parserNodeJoin(`LEFT JOIN /*TableSources注释测试*/ 规上月度1_as b on a.流水号 = b.流水号 INNER JOIN /*TableSources注释测试*/ (select 流水号 from 学生表 cc) C on a.流水号 = C.流水号`)

	spew.Dump(node)
	println("========================")
	println(joinClauseString(node, nil))
}
func TestTableSources(t *testing.T) {
	var sql = `/*TableSources注释测试*/ 规上月度1 a,/*TableSources注释测试*/规上月度1_as b`
	// sql = `e$自动化流程管理,规上月度1 a,规上月度1_as b,(select 名称 from 学生表) c`
	sql = `/*TableSources注释测试*/ e$自动化流程管理 a`
	nodes := parserNodeTableSources(sql)

	spew.Dump(nodes)
	println("========================")
	for _, v := range nodes {
		println(tableSourceString(v, nil))
	}
}

func TestSelectElements(t *testing.T) {
	//表达式原样返回
	var sql string
	// sql = `*`
	sql = `姓名,a.名字,a.字段1*a.字段2||a.字段3 AS 流水号/*注释*/,字段2 AS 名称号,字段3 别名3,a.姓名 AS 别名4,/*注释*/'b.姓名.信息' AS 别名5/*注释*/`
	// sql = `a.字段1 AS 流水号`
	// sql = `'啊啊' AS 流水号`
	sql = `*,a.*,'啊啊' AS 流水号`
	// sql = `a.*,b.*`
	sql = `(case when 行业代码 >= '06' and 行业代码 <= '46' and
	nvl(主营收入,0) >= 5000
 then 'B10'
when 行业代码 >= '47' and 行业代码 <= '50' and
	nvl(建筑业资质等级,'0000') > 'A000' and
	nvl(建筑业资质等级,'0000') <= 'C999'
 then 'C10'
when 行业代码 = '7210'
 then 'H10'
when (行业代码 like '63%' and nvl(主营收入,0) >=20000) or
	(行业代码 like '65%' and nvl(主营收入,0) >=5000)
 then 'E11'
when (行业代码 like '66%' or 行业代码 like '67%') and
	nvl(主营收入,0) >=2000
 then 'E21'
else NULL
end) as 专业标识`
	sql = `(只带了1) as 字段1`
	sql = `CASE 行业代码 WHEN '06' THEN 'B10'
	WHEN '47' THEN 'C10'
	ELSE ''
	END as 专业标识`
	sql = `(CASE WHEN 行业代码 = '06' THEN 'B10'
	WHEN 行业代码 = '47' THEN 'C10'
	ELSE NULL
	END) as 专业标识,'1' as 数值,a.字段 as 字段a`
	sql = `(select min(统计局代码) from t$法人单位表 d where (d.状态 is null or d.状态 <> '剔除') and a.归属法人组织机构代码 = d.组织机构代码) as 法人_统计局代码`
	sql = `substr(名字,1,2) as 字段`
	sql = `(select 1 from 表a) as 字段`
	node := parserNodeSelectelements(sql)
	spew.Dump(node)
	// spew.Dump(node.ColumnName (nil, "wholesql", nil))
	println("========================")
	println(selectElementsString(node, nil))

}
