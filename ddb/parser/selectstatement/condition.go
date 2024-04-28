package selectstatement

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
)

// NodeType 节点的类型
type NodeType string

// GetUserConditionViewDefineFunc 获取视图定义的函数
type GetUserConditionViewDefineFunc func(string) (string, error)

const (
	//ConditionNodeAnd And节点，所有下属子节点之间用And连接
	ConditionNodeAnd NodeType = "AND"
	//ConditionNodeOr Or节点，所有下属子节点之间用Or连接
	ConditionNodeOr NodeType = "OR"
	//ConditionNodeCondition 条件节点，没有子节点
	ConditionNodeCondition NodeType = "CONDITION"
	//ConditionNodePlain 条件节点，没有子节点,原始条件定义
	ConditionNodePlain NodeType = "PLAIN"
	//ConditionNodeExists exists条件节点
	ConditionNodeExists NodeType = "EXISTS"
	//ConditionNodeInTable 在表列表条件
	ConditionNodeInTable NodeType = "INTABLE"
	//ConditionNodeCount 子查询统计
	ConditionNodeCount NodeType = "COUNT"
	//ConditionCommentPlainText 纯文本节点的注释
	ConditionCommentPlainText = "/*PLAINTEXT*/"
)

// encodeCSV 压缩一个字符串数组成csv数据，没有换行
// func encodeCSV(val []string) string {
// 	bys := bytes.NewBuffer(nil)
// 	csvW := csv.NewWriter(bys)
// 	if err := csvW.Write(val); err != nil {
// 		log.Panic(err)
// 	}
// 	csvW.Flush()
// 	buf := bys.Bytes()
// 	//去掉尾部的回车
// 	return string(buf[0 : len(buf)-1])
// }

// decodeCSV 解开一个csv
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

// NodeLinkColumn 关联条件
type NodeLinkColumn struct {
	OuterColumn string //外层字段，即数据字段
	InnerColumn string //内层字段，即数据源字段
}

// NodeCondition 一个条件节点，可以有子节点，也可以是叶子
type NodeCondition struct {
	NodeType  NodeType
	Reverse   bool //条件是否反转，即是否加上not，对于 and or节点无效
	Field     string
	Type      schema.DataType //字段类型
	Func      string          //针对字段处理的函数,第一个参数是字段名
	Args      []string        //函数的参数
	Operate   pageselect.Operator
	Value     string
	Value2    string           //用于区间运算符，尾值
	PlainText string           //如果是关联查询，则是数据源附加的过滤条件
	From      string           //数据源，表名或视图名，仅exists intable count有用
	Link      []NodeLinkColumn //关联条件，如果是in，则不起作用
	InColumn  string           //in的内层字段

	SubSelect *NodeSelectStatement //如果是关联查询，nodetype=exists intable count，这里存放子查询，如果是从注释中获取，这里应该是nil，上面4个属性和这个是二选一

	Children []*NodeCondition
}

// NewLogicNode 分配一个逻辑节点，and、or
func NewLogicNode(nodeType NodeType, children []*NodeCondition) *NodeCondition {
	return &NodeCondition{
		NodeType: nodeType,
		Children: children,
	}
}

// NewConditionNode 分配一个条件节点
func NewFuncNode(field, funcName string, args []string) *NodeCondition {
	rev := &NodeCondition{
		NodeType: ConditionNodeCondition,
		Field:    field,
		Func:     funcName,
		Args:     args,
	}

	return rev
}

func NewSubSelectNode(subSelect *NodeSelectStatement) *NodeCondition {
	rev := &NodeCondition{
		NodeType:  ConditionNodeCondition,
		SubSelect: subSelect,
	}
	return rev
}

// NewConditionNode 分配一个条件节点
func NewConditionNode(field string, operate pageselect.Operator, value, value2 string) *NodeCondition {
	rev := &NodeCondition{
		NodeType: ConditionNodeCondition,
		Field:    field,
		Operate:  operate,
		Value:    value,
		Value2:   value2,
	}
	rev.Operate, rev.Reverse = rev.Operate.RemoveReverse()

	return rev
}

// NewPlainNode 分配一个文本条件节点
func NewPlainNode(text string) *NodeCondition {
	return &NodeCondition{
		NodeType:  ConditionNodePlain,
		PlainText: text,
	}
}

// NewInTableNode 分配一个InTable条件节点
func NewInTableNode(column, from, inColumn, where string, reverse bool) *NodeCondition {

	return &NodeCondition{
		NodeType:  ConditionNodeInTable,
		Field:     column,
		Reverse:   reverse,
		Operate:   pageselect.OperatorIn,
		From:      from,
		PlainText: where,
		InColumn:  inColumn,
	}
}

// NewInSubSelectNode 分配一个InTable子查询条件节点
func NewInSubSelectNode(field string, subSelect *NodeSelectStatement, reverse bool) *NodeCondition {
	return &NodeCondition{
		NodeType:  ConditionNodeInTable,
		Field:     field,
		Reverse:   reverse,
		SubSelect: subSelect,
	}
}

// NewExistsNode 分配一个Exists条件节点
func NewExistsNode(from string, link []NodeLinkColumn, where string, reverse bool) *NodeCondition {
	return &NodeCondition{
		NodeType:  ConditionNodeExists,
		From:      from,
		PlainText: where,
		Link:      link,
		Reverse:   reverse,
	}
}

// NewExistsSubSelectNode 分配一个Exists子查询条件节点
func NewExistsSubSelectNode(subSelect *NodeSelectStatement, reverse bool) *NodeCondition {
	return &NodeCondition{
		NodeType:  ConditionNodeExists,
		Reverse:   reverse,
		SubSelect: subSelect,
	}
}

// NewCountNode 分配一个Count条件节点
func NewCountNode(from string, link []NodeLinkColumn, where string,
	operate pageselect.Operator, value, value2 string) *NodeCondition {
	return &NodeCondition{
		NodeType:  ConditionNodeCount,
		From:      from,
		PlainText: where,
		Link:      link,
		Operate:   operate,
		Value:     value,
		Value2:    value2,
	}
}

// 简化一个Node，将and/or尽量合并成一层
func (node *NodeCondition) Reduction() {
	switch node.NodeType {
	case ConditionNodeAnd:
		nodes := []*NodeCondition{}
		for _, one := range node.Children {
			one.Reduction()
			if one.NodeType == ConditionNodeAnd {
				for _, sub := range one.Children {
					nodes = append(nodes, sub)
				}
			} else {
				nodes = append(nodes, one)
			}
		}
		node.Children = nodes
	case ConditionNodeOr:
		nodes := []*NodeCondition{}
		for _, one := range node.Children {
			one.Reduction()
			if one.NodeType == ConditionNodeOr {

				nodes = append(nodes, one.Children...)

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
func ifExpr(v string) bool {
	return len(v) >= 2 && v[0] == '`' && v[len(v)-1] == '`'
}
func decodeExpr(v string) string {
	if len(v) < 2 {
		return v
	}
	return v[1 : len(v)-1]
}
func ifels[T string](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}
func (node *NodeCondition) fieldName(getview GetUserConditionViewDefineFunc) string {
	fieldName := node.Field
	if len(node.Func) > 0 {
		args := ""
		if len(node.Args) > 0 {
			args = "," + strings.Join(node.Args, ",")
		}
		fieldName = fmt.Sprintf("%s(%s%s)", node.Func, node.Field, args)
	}
	if node.SubSelect != nil {
		fieldName = fmt.Sprintf("(%s)", node.SubSelect.SelectStatementString(getview))
	}
	return fieldName
}
func (node *NodeCondition) string(prev, outerTableName string, getview GetUserConditionViewDefineFunc,
	buildComment bool) string {
	if node == nil {
		return ""
	}
	switch node.NodeType {
	case ConditionNodePlain:
		comment := ""
		if buildComment {
			comment = prev + commentPlainText + "\n"
		}
		if node.Reverse {
			return comment + prev + "(not(" + node.PlainText + "))"
		}
		return comment + prev + "(" + node.PlainText + ")"
	case ConditionNodeCount:
		from := node.From
		if getview != nil {
			if viewDef, err := getview(from); err == nil {
				from = fmt.Sprintf("(%s) cnt_inner0", viewDef)
				if len(node.PlainText) > 0 {
					from = fmt.Sprintf("(select * from %s where %s) cnt_inner", from, node.PlainText)
				}
			} else {
				log.Println("getview error", err)
			}
		}
		iwhere := []string{}
		for _, one := range node.Link {
			if len(one.InnerColumn) > 0 && len(one.OuterColumn) > 0 {
				iwhere = append(iwhere, fmt.Sprintf("%s=%s.%s", one.InnerColumn, outerTableName, one.OuterColumn))
			}
		}
		strWhere := ""
		if len(iwhere) > 0 {
			strWhere = " where " + strings.Join(iwhere, " and\n")
		}
		countText := fmt.Sprintf("(select count(*) from %s%s)", from, strWhere)
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
	case ConditionNodeExists:
		if node.SubSelect != nil {
			cop := "EXISTS"
			if node.Reverse {
				cop = "NOT EXISTS"
			}
			return fmt.Sprintf("(%s(%s))", cop,
				node.SubSelect.SelectStatementString(getview))
		}
		from := node.From
		if getview != nil {
			if viewDef, err := getview(from); err == nil {
				from = fmt.Sprintf("(%s) exists_inner0", viewDef)
			} else {
				log.Println("getview error", err)
			}
		}
		if len(node.PlainText) > 0 {
			from = fmt.Sprintf("(select * from %s where %s) exists_inner", from, node.PlainText)
		}

		iwhere := []string{}
		for _, one := range node.Link {
			if len(one.InnerColumn) > 0 && len(one.OuterColumn) > 0 {
				iwhere = append(iwhere, fmt.Sprintf("%s=%s.%s", one.InnerColumn, outerTableName, one.OuterColumn))
			}
		}
		strWhere := ""
		if len(iwhere) > 0 {
			strWhere = " where " + strings.Join(iwhere, " and\n")
		}
		innerText := fmt.Sprintf("(select 1 from %s%s)", from, strWhere)
		cPrev := "EXISTS"
		if node.Reverse {
			cPrev = "NOT EXISTS"
		}
		comment := ""
		if buildComment {
			comment = prev + fmt.Sprintf("/*%s(from %s on %s where %s)*/\n",
				cPrev, node.From, strings.Join(iwhere, " and "), node.PlainText)
		}
		return comment + prev + "(" + cPrev + innerText + ")"
	case ConditionNodeInTable:
		if node.SubSelect != nil {
			cop := "in"
			if node.Reverse {
				cop = "not in"
			}
			return fmt.Sprintf("(%s %s (%s))", node.Field, cop,
				node.SubSelect.SelectStatementString(getview))
		}
		from := node.From
		if getview != nil {
			if viewDef, err := getview(from); err == nil {
				from = fmt.Sprintf("(%s) in_inner", viewDef)
			} else {
				log.Println("getview error", err)
			}
		}
		iwhere := ""
		if len(node.PlainText) > 0 {
			iwhere = fmt.Sprintf(" where %s", node.PlainText)
		}
		innerText := fmt.Sprintf("(select %s from %s%s)", node.InColumn, from, iwhere)
		cPrev := commentIn
		cop := "in"
		if node.Reverse {
			cPrev = commentNotIn
			cop = "not in"
		}
		comment := ""
		if buildComment {
			comment = prev + cPrev + fmt.Sprintf("%s in %s(%s) where %s)",
				node.Field, node.From, node.InColumn, node.PlainText) + "*/\n"
		}

		return comment + prev + "(" + node.Field + " " + cop + " " + innerText + ")"
	case ConditionNodeAnd:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev, outerTableName, getview, buildComment))
		}
		return prev + "(\n" + strings.Join(list, " AND\n") + "\n" + prev + ")"
	case ConditionNodeOr:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev, outerTableName, getview, buildComment))
		}
		// list[0] = strings.Repeat("\t", level+1) + "(" + strings.TrimSpace(list[0])
		// list[len(list)-1] = list[len(list)-1] + ")"
		return prev + "(\n" + strings.Join(list, " OR\n") + "\n" + prev + ")"
	case ConditionNodeCondition:

		op := node.Operate
		//计算反转
		if node.Reverse {
			op = op.Reverse()
		}
		//等于，不等于空值的，转换为is null 和 is not null
		if op == pageselect.OperatorEqu && len(node.Value) == 0 {
			op = pageselect.OperatorIsNull
		} else if op == pageselect.OperatorNotEqu && len(node.Value) == 0 {
			op = pageselect.OperatorIsNotNull
		}
		switch op {
		case pageselect.OperatorEqu, pageselect.OperatorGreaterThan,
			pageselect.OperatorGreaterThanOrEqu, pageselect.OperatorLessThan,
			pageselect.OperatorLessThanOrEqu, pageselect.OperatorRegexp,
			pageselect.OperatorNotRegexp, pageselect.OperatorNotEqu,
			pageselect.OperatorLikeArray, pageselect.OperatorNotLikeArray:
			v := node.Value
			if ifExpr(v) {
				v = v[1 : len(v)-1]
			} else {
				if node.isNumberField(node.Type) {
					//数值型的，空值自动转换成0
					if len(v) == 0 {
						v = "0"
					}
				} else {
					v = signString(v)
				}
			}

			return prev +
				fmt.Sprintf("%s %s %s", node.fieldName(getview),
					ifels(op == pageselect.OperatorNotEqu, "<>", op.String()), v)

		//OperatorLike 包含
		case pageselect.OperatorLike:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString("%"+node.Value+"%")))
		//OperatorNotLike 不包含
		case pageselect.OperatorNotLike:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString("%"+node.Value+"%")))

			//OperatorPrefix 前缀
		case pageselect.OperatorPrefix:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString(node.Value+"%")))
			//OperatorNotPrefix 非前缀
		case pageselect.OperatorNotPrefix:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString(node.Value+"%")))
			//OperatorSuffix 后缀
		case pageselect.OperatorSuffix:
			return prev +
				fmt.Sprintf("%s LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString("%"+node.Value)))
			//OperatorNotSuffix 非后缀
		case pageselect.OperatorNotSuffix:
			return prev +
				fmt.Sprintf("%s NOT LIKE %s", node.fieldName(getview), ifels(ifExpr(node.Value),
					decodeExpr(node.Value), signString("%"+node.Value)))
			//OperatorIn 在列表
		case pageselect.OperatorIn:
			list := []string{}
			for _, one := range decodeCSV(node.Value) {
				var v string
				if node.isNumberField(node.Type) {
					v = one
				} else {
					v = signString(one)
				}
				list = append(list, v)
			}
			return prev +
				fmt.Sprintf("%s IN (%s)", node.fieldName(getview), strings.Join(list, ","))
			//OperatorNotIn 不在列表
		case pageselect.OperatorNotIn:
			list := []string{}
			for _, one := range decodeCSV(node.Value) {
				var v string
				if node.isNumberField(node.Type) {
					v = one
				} else {
					v = signString(one)
				}
				list = append(list, v)
			}
			return prev +
				fmt.Sprintf("%s NOT IN (%s)", node.fieldName(getview), strings.Join(list, ","))
		//OperatorIsNull 为空
		case pageselect.OperatorIsNull:
			return prev +
				fmt.Sprintf("%s IS NULL", node.fieldName(getview))
			//OperatorIsNotNull is not null
		case pageselect.OperatorIsNotNull:
			return prev +
				fmt.Sprintf("%s IS NOT NULL", node.fieldName(getview))

			//OperatorLengthEqu 长度等于
		case pageselect.OperatorLengthEqu:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}
			return prev +
				fmt.Sprintf("LENGTH(%s) = %s", node.fieldName(getview), ifels(ifExpr(v), decodeExpr(v), v))

			//OperatorLengthNotEqu 长度不等于
		case pageselect.OperatorLengthNotEqu:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}

			return prev +
				fmt.Sprintf("LENGTH(%s) <> %s", node.Field, ifels(ifExpr(v), decodeExpr(v), v))
			//OperatorLengthGreaterThan 长度大于
		case pageselect.OperatorLengthGreaterThan:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}

			return prev +
				fmt.Sprintf("LENGTH(%s) > %s", node.Field, ifels(ifExpr(v), decodeExpr(v), v))
			//OperatorLengthGreaterThanOrEqu 长度 >=
		case pageselect.OperatorLengthGreaterThanOrEqu:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}

			return prev +
				fmt.Sprintf("LENGTH(%s) >= %s", node.Field, ifels(ifExpr(v), decodeExpr(v), v))
			//OperatorLengthLessThan 长度 <
		case pageselect.OperatorLengthLessThan:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}

			return prev +
				fmt.Sprintf("LENGTH(%s) < %s", node.Field, ifels(ifExpr(v), decodeExpr(v), v))
			//OperatorLengthLessThanOrEqu 长度<=
		case pageselect.OperatorLengthLessThanOrEqu:
			//空转换成0
			v := node.Value
			if len(v) == 0 {
				v = "0"
			}

			return prev +
				fmt.Sprintf("LENGTH(%s) <= %s", node.Field, ifels(ifExpr(v), decodeExpr(v), v))
		case pageselect.OperatorBetween:
			var v, v2 string
			v, v2 = node.Value, node.Value2
			if len(v) == 0 {
				v = "0"
			}
			if len(v2) == 0 {
				v2 = "0"
			}

			if node.isNumberField(node.Type) {
				v, v2 = ifels(ifExpr(v), decodeExpr(v), v),
					ifels(ifExpr(v2), decodeExpr(v2), v2)
			} else {
				v, v2 = ifels(ifExpr(v), decodeExpr(v), signString(v)),
					ifels(ifExpr(node.Value2), decodeExpr(v2), signString(v2))
			}
			return prev +
				fmt.Sprintf("%s between %s and %s", node.fieldName(getview), v, v2)
		case pageselect.OperatorNotBetween:
			var v, v2 string
			v, v2 = node.Value, node.Value2
			if len(v) == 0 {
				v = "0"
			}
			if len(v2) == 0 {
				v2 = "0"
			}

			if node.isNumberField(node.Type) {
				v, v2 = ifels(ifExpr(v), decodeExpr(v), v),
					ifels(ifExpr(v2), decodeExpr(v2), v2)
			} else {
				v, v2 = ifels(ifExpr(v), decodeExpr(v), signString(v)),
					ifels(ifExpr(node.Value2), decodeExpr(v2), signString(v2))
			}
			return prev +
				fmt.Sprintf("%s not between %s and %s", node.fieldName(getview), v, v2)
		default:
			panic("not impl " + node.Operate.String())
		}

	default:
		panic("not impl")
	}
}
func (node *NodeCondition) isNumberField(dat schema.DataType) bool {
	return len(node.Func) == 0 && (dat == schema.TypeInt || dat == schema.TypeFloat)
}

// WhereString 返回规范化的where条件,传入视图列表，用于关联表查询的语句
func (node *NodeCondition) WhereString(fields map[string]schema.DataType, outerTableName string,
	getview GetUserConditionViewDefineFunc, buildComment bool) string {
	node.setType(fields)
	return node.string("", outerTableName, getview, buildComment)
}
func (node *NodeCondition) setType(fields map[string]schema.DataType) {
	node.Type = fields[node.Field]
	for _, one := range node.Children {
		one.setType(fields)
	}
}

// ReferToColumns 条件中涉及到的列
func (node *NodeCondition) ReferToColumns() []string {
	rev := []string{}
	mapRev := map[string]struct{}{}
	switch node.NodeType {
	case ConditionNodeAnd, ConditionNodeOr:
		for _, one := range node.Children {
			for _, c := range one.ReferToColumns() {
				if _, ok := mapRev[c]; !ok {
					rev = append(rev, c)
					mapRev[c] = struct{}{}
				}
			}

		}
	case ConditionNodeCondition:
		rev = append(rev, node.Field)
	case ConditionNodePlain:
	}
	return rev
}

// ParserNode 根据一个where条件，返回node
func ParserWhereNode(val string) *NodeCondition {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	val, vars = ProcessComment(val)
	stream := antlr.NewInputStream(`where ` + val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.WhereClause()
	visitor := new(sqlWhereVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*NodeCondition)

}
func WhereParseByContext(ctx parser.IWhereClauseContext, vars map[string]interface{}) *NodeCondition {
	visitor := new(sqlWhereVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).(*NodeCondition)
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
func processPlainText(define string) *NodeCondition {
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
func processCount(comment string) *NodeCondition {
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
			if len(arrs) == 2 {
				//去除.号前面的表名
				link = append(link, NodeLinkColumn{
					InnerColumn: arrs[0],
					OuterColumn: strings.Split(arrs[1], ".")[1],
				})
			}
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
func processExists(comment string) *NodeCondition {
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
			if len(arrs) == 2 {
				//去除.号前面的表名
				link = append(link, NodeLinkColumn{
					InnerColumn: arrs[0],
					OuterColumn: strings.Split(arrs[1], ".")[1],
				})
			}
		}
		node := NewExistsNode(from, link, where, reverse)
		return node
	}
	panic("invalid exists format:" + txt)
}
func processIn(comment string) *NodeCondition {
	var txt string
	reverse := false
	if strings.HasPrefix(comment, commentNotIn) {
		//去除 /*not in( 前缀
		txt = comment[9 : len(comment)-3]
		reverse = true
	} else {
		txt = comment[5 : len(comment)-3]
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

// 处理注释，识别关联查询并生成node列表
func ProcessComment(define string) (rev string, vars map[string]interface{}) {
	wait := define
	vars = map[string]interface{}{}
	iDynamic := 0
	addDynamicNode := func(node *NodeCondition) string {
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
		wait = wait[positions[3]:]
		//保留其它注释
		rev += comment
	}
	rev += wait
	return
}

// ConditionLines 遍历树，返回条件数组
func (node *NodeCondition) ConditionLines(fields map[string]schema.DataType, outerTableName string, getview GetUserConditionViewDefineFunc) []*pageselect.ConditionLine {
	rev := []*pageselect.ConditionLine{}
	switch node.NodeType {
	case ConditionNodeAnd:
		for _, one := range node.Children {
			subConts := one.ConditionLines(fields, outerTableName, getview)
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
	case ConditionNodeOr:
		for _, one := range node.Children {
			subConts := one.ConditionLines(fields, outerTableName, getview)
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

	case ConditionNodeCondition:
		op := node.Operate
		if node.Reverse {
			op = op.Reverse()
		}
		rev = append(rev, &pageselect.ConditionLine{
			ColumnName: node.Field,
			Func:       node.Func,
			Args:       node.Args,
			Operators:  op,
			Value:      node.Value,
			Value2:     node.Value2,
		})
	case ConditionNodePlain:
		txt := node.PlainText
		if node.Reverse {
			txt = "not(" + node.PlainText + ")"
		}
		rev = append(rev, &pageselect.ConditionLine{
			PlainText: txt,
		})
	// 关联
	case ConditionNodeCount, ConditionNodeExists, ConditionNodeInTable:
		rev = append(rev, &pageselect.ConditionLine{
			PlainText: node.WhereString(fields, outerTableName, getview, true),
		})

	}
	return rev
}
