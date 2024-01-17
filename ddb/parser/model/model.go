package model

import (
	"regexp"

	"github.com/linlexing/dbx/ddb/parser/condition"
)

type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	Join      JoinType = "JOIN"
)

var regComment = regexp.MustCompile(`(?:[^']|'[^']*')*?(/\*[^*]*\*+(?:[^/*][^*]*\*+)*/)`)

type NodeSelectStatement struct {
	SelectElements *NodeSelectelements
	TableSources   []*NodeTableSource
	JoinClause     []*NodeJoinClause
	WhereClause    *condition.Node
	UnionSelect    []*NodeSelectStatement
	UnionAll       bool
	// 	GroupByClause
	// 	HavingClause
	// 	OrderByClause
	// 	LimitClause
}

type NodeSelectelements struct {
	Elements []*Element
}
type Element struct {
	TableAlias string //a.X 单纯字段才识别
	ColumnName string //a.X 单纯字段才识别
	Express    string //表达式
	Subquery   *NodeSelectStatement
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
	JoinType    JoinType
	TableSource *NodeTableSource
	OnExpress   *condition.Node
}
