package condition

import (
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
