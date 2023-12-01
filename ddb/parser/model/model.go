package model

import (
	"github.com/linlexing/dbx/ddb/parser/condition"
)

type SelectElementsType string
type JoinType string

const (
	//*
	NodeStar SelectElementsType = "Star"
	//元素
	NodeElements SelectElementsType = "Elements"

	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	Join      JoinType = "JOIN"
)

type NodeSelectStatement struct {
	SelectElements *NodeSelectelements
	TableSources   []*NodeTableSource
	JoinClause     []*NodeJoinClause
	WhereClause    *condition.Node
	// 	GroupByClause
	// 	HavingClause
	// 	OrderByClause
	// 	LimitClause
	// 	UnionSelect
}

type NodeSelectelements struct {
	NodeType SelectElementsType
	Elements []*Element
}
type Element struct {
	TableAlias string //a.X 单纯字段才识别
	ColumnName string //a.X 单纯字段才识别
	Express    string //表达式
	As         string
	Alias      string
}

type NodeTableSource struct {
	Source *Source
	Alias  string
}
type Source struct {
	TableName       string
	SelectStatement *NodeSelectStatement
}
type NodeJoinClause struct {
	JoinType        JoinType
	TableSources    []*NodeTableSource
	LogicExpression *condition.Node
}
