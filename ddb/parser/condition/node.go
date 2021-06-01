package condition

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
)

//NodeType 节点的类型
type NodeType string

const (
	//NodeAnd And节点，所有下属子节点之间用And连接
	NodeAnd NodeType = "AND"
	//NodeOr Or节点，所有下属子节点之间用Or连接
	NodeOr NodeType = "OR"
	//NodeCondition 条件节点，没有子节点
	NodeCondition NodeType = "CONDITION"
	//NodePlain 条件节点，没有子节点,原始条件定义
	NodePlain NodeType = "PLAIN"
	//NodeExists exists条件节点
	NodeExists NodeType = "EXISTS"
	//NodeInTable 在表列表条件
	NodeInTable NodeType = "INTABLE"
	//NodeCount 子查询统计
	NodeCount NodeType = "COUNT"
	//CommentPlainText 纯文本节点的注释
	CommentPlainText = "/*PLAINTEXT*/"
)

//encodeCSV 压缩一个字符串数组成csv数据，没有换行
func encodeCSV(val []string) string {
	bys := bytes.NewBuffer(nil)
	csvW := csv.NewWriter(bys)
	if err := csvW.Write(val); err != nil {
		log.Panic(err)
	}
	csvW.Flush()
	buf := bys.Bytes()
	//去掉尾部的回车
	return string(buf[0 : len(buf)-1])
}

//decodeCSV 解开一个csv
func decodeCSV(val string) []string {
	if len(val) == 0 {
		return nil
	}
	s, err := csv.NewReader(bytes.NewBufferString(val)).Read()
	if err != nil {
		log.Panic(err)
	}
	return s
}

//NodeLinkColumn 关联条件
type NodeLinkColumn struct {
	OuterColumn string //外层字段，即数据字段
	InnerColumn string //内层字段，即数据源字段
}

//Node 一个条件节点，可以有子节点，也可以是叶子
type Node struct {
	NodeType  NodeType
	Reverse   bool // 条件是否反转，即是否加上not，对于 and or节点无效
	Field     string
	Operate   pageselect.Operator
	Value     string
	Value2    string           //用于区间运算符，尾值
	PlainText string           //如果是关联查询，则是数据源附加的过滤条件
	From      string           //数据源，表名或视图名，仅exists intable有用
	Link      []NodeLinkColumn //关联条件，如果是in，则不起作用
	InColumn  string           //in的内层字段
	Children  []*Node
}

//NewLogicNode 分配一个逻辑节点，and、or
func NewLogicNode(nodeType NodeType, children []*Node) *Node {
	return &Node{
		NodeType: nodeType,
		Children: children,
	}
}

//NewConditionNode 分配一个条件节点
func NewConditionNode(field string, operate pageselect.Operator, value, value2 string) *Node {
	rev := &Node{
		NodeType: NodeCondition,
		Field:    field,
		Operate:  operate,
		Value:    value,
		Value2:   value2,
	}
	rev.Operate, rev.Reverse = rev.Operate.RemoveReverse()

	return rev
}

//NewPlainNode 分配一个文本条件节点
func NewPlainNode(text string) *Node {
	return &Node{
		NodeType:  NodePlain,
		PlainText: text,
	}
}

//NewInTableNode 分配一个InTable条件节点
func NewInTableNode(column, from, inColumn, where string, reverse bool) *Node {
	op := pageselect.OperatorIn
	if reverse {
		op = pageselect.OperatorNotIn
	}
	return &Node{
		NodeType:  NodeInTable,
		Field:     column,
		Operate:   op,
		From:      from,
		PlainText: where,
		InColumn:  inColumn,
	}
}

//NewExistsNode 分配一个Exists条件节点
func NewExistsNode(from string, link []NodeLinkColumn, where string, reverse bool) *Node {
	return &Node{
		NodeType:  NodeExists,
		From:      from,
		PlainText: where,
		Link:      link,
		Reverse:   reverse,
	}
}

//NewCountNode 分配一个Count条件节点
func NewCountNode(from string, link []NodeLinkColumn, where string,
	operate pageselect.Operator, value, value2 string) *Node {
	return &Node{
		NodeType:  NodeCount,
		From:      from,
		PlainText: where,
		Link:      link,
		Operate:   operate,
		Value:     value,
		Value2:    value2,
	}
}

//简化一个Node，将and/or尽量合并成一层
func (node *Node) reduction() {
	switch node.NodeType {
	case NodeAnd:
		nodes := []*Node{}
		for _, one := range node.Children {
			one.reduction()
			if one.NodeType == NodeAnd {
				for _, sub := range one.Children {
					nodes = append(nodes, sub)
				}
			} else {
				nodes = append(nodes, one)
			}
		}
		node.Children = nodes
	case NodeOr:
		nodes := []*Node{}
		for _, one := range node.Children {
			one.reduction()
			if one.NodeType == NodeOr {
				for _, sub := range one.Children {
					nodes = append(nodes, sub)
				}
			} else {
				nodes = append(nodes, one)
			}
		}
		node.Children = nodes
	}
}
func signString(str string) string {
	return "'" + strings.ReplaceAll(str, "'", "''") + "'"
}
func (node *Node) string(prev string, fields map[string]schema.DataType, outerTableName string, views map[string]string, buildComment bool) string {
	switch node.NodeType {
	case NodePlain:
		comment := ""
		if buildComment {
			comment = prev + commentPlainText + "\n"
		}
		if node.Reverse {
			return comment + prev + "(not(" + node.PlainText + "))"
		}
		return comment + prev + "(" + node.PlainText + ")"
	case NodeCount:
		from := node.From
		if len(views) > 0 {
			if viewDef, ok := views[from]; ok {
				from = fmt.Sprintf("(%s) cnt_inner", viewDef)
				if len(node.PlainText) > 0 {
					from = fmt.Sprintf("(select * from (%s) cnt_inner0 where %s) cnt_inner", viewDef, node.PlainText)
				}
			}
		}
		iwhere := []string{}
		for _, one := range node.Link {
			iwhere = append(iwhere, fmt.Sprintf("%s=%s.%s", one.InnerColumn, outerTableName, one.OuterColumn))
		}
		countText := fmt.Sprintf("(select count(*) from %s where %s)", from,
			strings.Join(iwhere, " and\n"))
		valStr := node.Value
		op := node.Operate
		if node.Reverse {
			op = op.Reverse()
		}
		if op == pageselect.OperatorBetween || op == pageselect.OperatorNotBetween {
			valStr = node.Value + "," + node.Value2
		}
		comment := ""
		if buildComment {
			comment = prev + fmt.Sprintf("/*COUNT(from %s on %s where %s) %s %s*/\n",
				node.From, strings.Join(iwhere, " and "), node.PlainText, op.String(), valStr)
		}
		//数值支持少数几种运算符
		switch op {
		case pageselect.OperatorEqu, pageselect.OperatorGreaterThan,
			pageselect.OperatorGreaterThanOrEqu, pageselect.OperatorLessThan,
			pageselect.OperatorLessThanOrEqu:
			return comment + prev +
				fmt.Sprintf("(%s %s %s)", countText, op.String(), node.Value)
		case pageselect.OperatorNotEqu:
			return comment + prev +
				fmt.Sprintf("(%s <> %s)", countText, node.Value)
		case pageselect.OperatorBetween:
			return comment + prev +
				fmt.Sprintf("(%s between %s and %s)", countText, node.Value, node.Value2)
		case pageselect.OperatorNotBetween:
			return comment + prev +
				fmt.Sprintf("(%s not between %s and %s)", countText, node.Value, node.Value2)
		default:
			panic("invalid op:" + op.String() + " at linkcount")
		}
	case NodeExists:
		from := node.From
		if viewDef, ok := views[from]; ok {
			from = fmt.Sprintf("(%s) exists_inner", viewDef)
			if len(node.PlainText) > 0 {
				from = fmt.Sprintf("(select * from (%s) exists_inner0 where %s) exists_inner", viewDef, node.PlainText)
			}
		} else {
			if len(node.PlainText) > 0 {
				from = fmt.Sprintf("(select * from %s exists_inner0 where %s) exists_inner", from, node.PlainText)
			}
		}
		iwhere := []string{}
		for _, one := range node.Link {
			iwhere = append(iwhere, fmt.Sprintf("%s=%s.%s", one.InnerColumn, outerTableName, one.OuterColumn))
		}
		innerText := fmt.Sprintf("(select 1 from %s where %s)", from,
			strings.Join(iwhere, " and\n"))
		cPrev := "EXISTS"
		if node.Reverse {
			cPrev = "NOT EXISTS"
		}
		comment := ""
		if buildComment {
			comment = prev + fmt.Sprintf("/*%s(from %s on %s where %s)*/\n",
				cPrev, node.From, strings.Join(iwhere, " and "), node.PlainText)
		}
		return comment + prev + "(" + cPrev + "(" + innerText + "))"
	case NodeInTable:
		from := node.From
		if viewDef, ok := views[from]; ok {
			from = fmt.Sprintf("(%s) in_inner", viewDef)
		}
		iwhere := ""
		if len(node.PlainText) > 0 {
			iwhere = fmt.Sprintf("where %s", node.PlainText)
		}

		innerText := fmt.Sprintf("(select 1 from %s where %s)", from, iwhere)
		cPrev := commentIn
		cop := "in"
		if node.Reverse {
			cPrev = commentNotIn
			cop = "not in"
		}
		comment := ""
		if buildComment {
			comment = prev + cPrev + fmt.Sprintf("%s in %s(%s) where %s",
				node.Field, node.From, node.InColumn, node.PlainText) + "*/\n"
		}

		return comment + prev + "(" + node.Field + " " + cop + " " + innerText + ")"
	case NodeAnd:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev, fields, outerTableName, views, buildComment))
		}
		return prev + "(\n" + strings.Join(list, " AND\n") + "\n" + prev + ")"
	case NodeOr:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev, fields, outerTableName, views, buildComment))
		}
		// list[0] = strings.Repeat("\t", level+1) + "(" + strings.TrimSpace(list[0])
		// list[len(list)-1] = list[len(list)-1] + ")"
		return prev + "(\n" + strings.Join(list, " OR\n") + "\n" + prev + ")"
	case NodeCondition:
		op := node.Operate
		//计算反转
		if node.Reverse {
			op = op.Reverse()
		}
		switch op {
		case pageselect.OperatorEqu, pageselect.OperatorGreaterThan,
			pageselect.OperatorGreaterThanOrEqu, pageselect.OperatorLessThan,
			pageselect.OperatorLessThanOrEqu, pageselect.OperatorRegexp, pageselect.OperatorNotRegexp:
			var v string
			if fields[node.Field] == schema.TypeInt ||
				fields[node.Field] == schema.TypeFloat {
				v = node.Value
			} else {
				v = signString(node.Value)
			}
			return prev +
				fmt.Sprintf("%s %s %s", node.Field, op.String(), v)
		case pageselect.OperatorNotEqu:
			var v string
			if fields[node.Field] == schema.TypeInt ||
				fields[node.Field] == schema.TypeFloat {
				v = node.Value
			} else {
				v = signString(node.Value)
			}
			return prev +
				fmt.Sprintf("%s <> %s", node.Field, v)
		//OperatorLike 包含
		case pageselect.OperatorLike:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.Field, signString("%"+node.Value+"%"))
		//OperatorNotLike 不包含
		case pageselect.OperatorNotLike:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.Field, signString("%"+node.Value+"%"))
			//OperatorPrefix 前缀
		case pageselect.OperatorPrefix:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.Field, signString(node.Value+"%"))
			//OperatorNotPrefix 非前缀
		case pageselect.OperatorNotPrefix:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.Field, signString(node.Value+"%"))
			//OperatorSuffix 后缀
		case pageselect.OperatorSuffix:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.Field, signString("%"+node.Value))
			//OperatorNotSuffix 非后缀
		case pageselect.OperatorNotSuffix:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.Field, signString("%"+node.Value))
			//OperatorIn 在列表
		case pageselect.OperatorIn:
			list := []string{}
			for _, one := range decodeCSV(node.Value) {
				var v string
				if fields[node.Field] == schema.TypeInt ||
					fields[node.Field] == schema.TypeFloat {
					v = one
				} else {
					v = signString(one)
				}
				list = append(list, v)
			}
			return prev +
				fmt.Sprintf("%s IN (%s)", node.Field, encodeCSV(list))
			//OperatorNotIn 不在列表
		case pageselect.OperatorNotIn:
			list := []string{}
			for _, one := range decodeCSV(node.Value) {
				var v string
				if fields[node.Field] == schema.TypeInt ||
					fields[node.Field] == schema.TypeFloat {
					v = one
				} else {
					v = signString(one)
				}
				list = append(list, v)
			}
			return prev +
				fmt.Sprintf("%s NOT IN (%s)", node.Field, encodeCSV(list))
			//OperatorIsNull 为空
		case pageselect.OperatorIsNull:
			return prev +
				fmt.Sprintf("%s IS NULL", node.Field)
			//OperatorIsNotNull is not null
		case pageselect.OperatorIsNotNull:
			return prev +
				fmt.Sprintf("%s IS NOT NULL", node.Field)

			//OperatorLengthEqu 长度等于
		case pageselect.OperatorLengthEqu:
			return prev +
				fmt.Sprintf("LENGTH(%s) = %s", node.Field, node.Value)

			//OperatorLengthNotEqu 长度不等于
		case pageselect.OperatorLengthNotEqu:
			return prev +
				fmt.Sprintf("LENGTH(%s) <> %s", node.Field, node.Value)
			//OperatorLengthGreaterThan 长度大于
		case pageselect.OperatorLengthGreaterThan:
			return prev +
				fmt.Sprintf("LENGTH(%s) > %s", node.Field, node.Value)
			//OperatorLengthGreaterThanOrEqu 长度 >=
		case pageselect.OperatorLengthGreaterThanOrEqu:
			return prev +
				fmt.Sprintf("LENGTH(%s) >= %s", node.Field, node.Value)
			//OperatorLengthLessThan 长度 <
		case pageselect.OperatorLengthLessThan:
			return prev +
				fmt.Sprintf("LENGTH(%s) < %s", node.Field, node.Value)
			//OperatorLengthLessThanOrEqu 长度<=
		case pageselect.OperatorLengthLessThanOrEqu:
			return prev +
				fmt.Sprintf("LENGTH(%s) <= %s", node.Field, node.Value)
		case pageselect.OperatorBetween:
			var v, v2 string
			if fields[node.Field] == schema.TypeInt ||
				fields[node.Field] == schema.TypeFloat {
				v, v2 = node.Value, node.Value2
			} else {
				v, v2 = signString(node.Value), signString(node.Value2)
			}
			return prev +
				fmt.Sprintf("%s between %s and %s", node.Field, v, v2)
		case pageselect.OperatorNotBetween:
			var v, v2 string
			if fields[node.Field] == schema.TypeInt ||
				fields[node.Field] == schema.TypeFloat {
				v, v2 = node.Value, node.Value2
			} else {
				v, v2 = signString(node.Value), signString(node.Value2)
			}
			return prev +
				fmt.Sprintf("%s not between %s and %s", node.Field, v, v2)
		default:
			panic("not impl " + node.Operate.String())
		}

	default:
		panic("not impl")
	}
}

//WhereString 返回规范化的where条件,传入视图列表，用于关联表查询的语句
func (node *Node) WhereString(fields map[string]schema.DataType, outerTableName string, views map[string]string, buildComment bool) string {
	return node.string("", fields, outerTableName, views, buildComment)
}

//ReferToColumns 条件中涉及到的列
func (node *Node) ReferToColumns() []string {
	rev := []string{}
	switch node.NodeType {
	case NodeAnd, NodeOr:
		for _, one := range node.Children {
			rev = append(rev, one.ReferToColumns()...)
		}
	case NodeCondition:
		rev = append(rev, node.Field)
	case NodePlain:
	}
	return rev
}

//ParserNode 根据一个where条件，返回node
func ParserNode(val string) *Node {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	val, vars = processComment(val)
	stream := antlr.NewInputStream(`where ` + val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.WhereClause()
	visitor := new(SqlWhereVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*Node)

}
func findBracketExpr(val string) (left string, right string) {
	iBracket := 0
	bSignQuoted := false
	for _, c := range val {
		left += string(c)
		switch c {
		case '(':
			if !bSignQuoted {
				iBracket++
			}
		case ')':
			if !bSignQuoted {
				iBracket--
			}
		case '\'':
			bSignQuoted = !bSignQuoted
		}
		if iBracket == 0 {
			break
		}
	}
	right = val[len(left):]
	return
}
func processPlainText(define string) *Node {
	if len(define) > 0 {
		txt := define[1 : len(define)-1] //去除括号
		reverse := false
		if strings.HasPrefix(txt, "not(") {
			reverse = true
			txt = txt[4 : len(txt)-1] //去除not及括号
		}
		node := NewPlainNode(txt)
		node.Reverse = reverse
		return node
	}
	return nil
}
func processCount(comment string) *Node {
	txt := comment[2 : len(comment)-2]
	if ps := regCount.FindStringSubmatchIndex(txt); len(ps) == 12 {
		from := txt[ps[2]:ps[3]]
		on := txt[ps[4]:ps[5]]
		where := txt[ps[6]:ps[7]]
		opStr := txt[ps[8]:ps[9]]
		valueStr := txt[ps[10]:ps[11]]
		link := []NodeLinkColumn{}
		for _, one := range strings.Split(on, "and") {
			one = strings.TrimSpace(one)
			arrs := strings.Split(one, "=")
			if len(arrs) != 2 {
				panic("invalid link:" + one)
			}
			//去除.号前面的表名
			link = append(link, NodeLinkColumn{
				InnerColumn: arrs[0],
				OuterColumn: strings.Split(arrs[1], ".")[1],
			})
		}
		op, err := pageselect.ParseOperatorFromString(opStr)
		if err != nil {
			panic(err)
		}
		value := valueStr
		value2 := ""
		if op == pageselect.OperatorBetween || op == pageselect.OperatorNotBetween {
			vals := strings.Split(valueStr, ",")
			value, value2 = vals[0], vals[1]
		}
		node := NewCountNode(from, link, where, op, value, value2)
		return node
	}
	panic("invalid count format:" + txt)
}
func processExists(comment string) *Node {
	txt := comment[2 : len(comment)-2]
	reverse := false
	if strings.HasPrefix(comment, commentNotExists) {
		//去除 not 前缀
		txt = txt[4:]
		reverse = true
	}
	//去除exists前缀和括号
	txt = txt[7 : len(txt)-1]
	if ps := regExists.FindStringSubmatchIndex(txt); len(ps) == 8 {
		from := txt[ps[2]:ps[3]]
		on := txt[ps[4]:ps[5]]
		where := txt[ps[6]:ps[7]]
		link := []NodeLinkColumn{}
		for _, one := range strings.Split(on, "and") {
			one = strings.TrimSpace(one)
			arrs := strings.Split(one, "=")
			if len(arrs) != 2 {
				panic("invalid link:" + one)
			}
			//去除.号前面的表名
			link = append(link, NodeLinkColumn{
				InnerColumn: arrs[0],
				OuterColumn: strings.Split(arrs[1], ".")[1],
			})
		}
		node := NewExistsNode(from, link, where, reverse)
		return node
	}
	panic("invalid exists format:" + txt)
}
func processIn(comment string) *Node {
	var txt string
	reverse := false
	if strings.HasPrefix(comment, commentNotIn) {
		//去除 /*not in( 前缀
		txt = comment[9 : len(comment)-3]
		reverse = true
	} else {
		txt = comment[6 : len(comment)-3]
	}

	if ps := regInTable.FindStringSubmatchIndex(txt); len(ps) == 10 {
		field := txt[ps[2]:ps[3]]
		from := txt[ps[4]:ps[5]]
		inField := txt[ps[6]:ps[7]]
		where := txt[ps[8]:ps[9]]
		node := NewInTableNode(field, from, inField, where, reverse)
		return node
	}
	panic("invalid intable format:" + txt)
}

//处理注释，识别关联查询并生成node列表
func processComment(define string) (rev string, vars map[string]interface{}) {
	wait := define
	vars = map[string]interface{}{}
	iDynamic := 0
	addDynamicNode := func(node *Node) string {
		iDynamic++
		id := fmt.Sprintf("dyna%d", iDynamic)
		vars[id] = node
		return fmt.Sprintf("%s(id=%s)", commentDynamicNode, id)
	}
	for positions := regComment.FindStringSubmatchIndex(
		wait); len(positions) == 4; positions = regComment.FindStringSubmatchIndex(wait) {

		comment := wait[positions[2]:positions[3]]
		rev += wait[:positions[2]] //注释之前的截断到返回值中
		if comment == commentPlainText {
			afterText := strings.TrimSpace(wait[positions[3]:])
			left, right := findBracketExpr(afterText)
			node := processPlainText(left)
			if node != nil {
				rev += addDynamicNode(node)
			}
			//光有注释，没有后面内容的，注释被清除
			wait = right
			continue
		}
		if strings.HasPrefix(comment, commentCount) {
			_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
			rev += addDynamicNode(processCount(comment))
			continue
		}
		if strings.HasPrefix(comment, commentExists) || strings.HasPrefix(comment, commentNotExists) {
			_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
			rev += addDynamicNode(processExists(comment))
			continue
		}
		if strings.HasPrefix(comment, commentIn) || strings.HasPrefix(comment, commentNotIn) {
			_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
			rev += addDynamicNode(processIn(comment))
			continue
		}
		//普通的注释，不做处理，纳入结果中
		rev += comment
		wait = wait[positions[3]:]
	}
	rev += wait
	return
}

//ConditionLines 遍历树，返回条件数组
func (node *Node) ConditionLines(fields map[string]schema.DataType, outerTableName string, views map[string]string) []*pageselect.ConditionLine {
	rev := []*pageselect.ConditionLine{}
	switch node.NodeType {
	case NodeAnd:
		for _, one := range node.Children {
			subConts := one.ConditionLines(fields, outerTableName, views)
			//如果已经有条件，且子条件是多行，则需要加上括号和and
			if len(rev) > 0 && len(subConts) > 1 {
				subConts[0].LeftBrackets += "("
				subConts[len(subConts)-1].RightBrackets += ")"
			}
			if len(rev) > 0 {
				rev[len(rev)-1].Logic = pageselect.AND
			}
			rev = append(rev, subConts...)
		}
	case NodeOr:
		for _, one := range node.Children {
			subConts := one.ConditionLines(fields, outerTableName, views)
			//如果已经有条件，且子条件是多行，则需要加上括号和and
			if len(rev) > 0 && len(subConts) > 1 {
				subConts[0].LeftBrackets += "("
				subConts[len(subConts)-1].RightBrackets += ")"
			}
			if len(rev) > 0 {
				rev[len(rev)-1].Logic = pageselect.OR
			}
			rev = append(rev, subConts...)
		}
		//OR最后需要加上括号
		if len(rev) > 1 {
			rev[0].LeftBrackets += "("
			rev[len(rev)-1].RightBrackets += ")"
		}

	case NodeCondition:
		op := node.Operate
		if node.Reverse {
			op = op.Reverse()
		}
		rev = append(rev, &pageselect.ConditionLine{
			ColumnName: node.Field,
			Operators:  op,
			Value:      node.Value,
			Value2:     node.Value2,
		})
	case NodePlain:
		txt := node.PlainText
		if node.Reverse {
			txt = "not(" + node.PlainText + ")"
		}
		rev = append(rev, &pageselect.ConditionLine{
			PlainText: txt,
		})
	// 关联
	case NodeCount, NodeExists, NodeInTable:
		rev = append(rev, &pageselect.ConditionLine{
			PlainText: node.WhereString(fields, outerTableName, views, false),
		})

	}
	return rev
}
