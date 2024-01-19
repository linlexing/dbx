package selectstatement

type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	Join      JoinType = "JOIN"
)

type NodeSelectStatement struct {
	SelectElements *NodeSelectelements
	TableSources   []*NodeTableSource
	JoinClause     []*NodeJoinClause
	WhereClause    *NodeCondition
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
	OnExpress   *NodeCondition
}
