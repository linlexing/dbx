package condition

import (
	"dbweb/lib/strfun"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/pageselect"
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
)

//Node 一个条件节点，可以有子节点，也可以是叶子
type Node struct {
	NodeType  NodeType
	Field     string
	Operate   pageselect.Operator
	Value     string
	PlainText string
	Children  []*Node
}

func NewLogicNode(nodeType NodeType, children []*Node) *Node {
	return &Node{
		NodeType: nodeType,
		Children: children,
	}
}
func NewConditionNode(field string, operate pageselect.Operator, value string) *Node {
	return &Node{
		NodeType: NodeCondition,
		Field:    field,
		Operate:  operate,
		Value:    value,
	}
}
func NewPlainNode(text string) *Node {
	return &Node{
		NodeType:  NodePlain,
		PlainText: text,
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
func (node *Node) string(prev string) string {
	switch node.NodeType {
	case NodePlain:
		return prev + node.PlainText
	case NodeAnd:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev))
		}
		return prev + "(\n" + strings.Join(list, " AND\n") + "\n" + prev + ")"
	case NodeOr:
		list := []string{}
		for _, one := range node.Children {
			list = append(list, one.string("\t"+prev))
		}
		// list[0] = strings.Repeat("\t", level+1) + "(" + strings.TrimSpace(list[0])
		// list[len(list)-1] = list[len(list)-1] + ")"
		return prev + "(\n" + strings.Join(list, " OR\n") + "\n" + prev + ")"
	case NodeCondition:
		switch node.Operate {
		case pageselect.OperatorEqu, pageselect.OperatorNotEqu, pageselect.OperatorGreaterThan,
			pageselect.OperatorGreaterThanOrEqu, pageselect.OperatorLessThan,
			pageselect.OperatorLessThanOrEqu, pageselect.OperatorLike,
			pageselect.OperatorNotLike, pageselect.OperatorRegexp, pageselect.OperatorNotRegexp:
			return prev +
				fmt.Sprintf("%s %s %s", node.Field, node.Operate.String(), node.Value)

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
			for _, one := range strfun.DecodeCSV(node.Value) {
				list = append(list, signString(one))
			}
			return prev +
				fmt.Sprintf("%s NOT IN (%s)", node.Field, strfun.EncodeCSV(list))
			//OperatorNotIn 不在列表
		case pageselect.OperatorNotIn:
			list := []string{}
			for _, one := range strfun.DecodeCSV(node.Value) {
				list = append(list, signString(one))
			}
			return prev +
				fmt.Sprintf("%s IN (%s)", node.Field, strfun.EncodeCSV(list))
			//OperatorIsNull 为空
		case pageselect.OperatorIsNull:
			return prev +
				fmt.Sprintf("%s IN NULL", node.Field)
			//OperatorIsNotNull is not null
		case pageselect.OperatorIsNotNull:
			return prev +
				fmt.Sprintf("%s IN NOT NULL", node.Field)

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
		default:
			panic("not impl " + node.Operate.String())
		}

	default:
		panic("not impl")
	}
}
func (node *Node) WhereString() string {
	return node.string("")
}

//ParserNode 根据一个where条件，返回node
func ParserNode(val string) *Node {
	if len(val) == 0 {
		return nil
	}
	stream := antlr.NewInputStream(`where ` + val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.WhereClause()
	visitor := new(SqlWhereVisitorImpl)
	return visitor.Visit(tree).(*Node)

}
