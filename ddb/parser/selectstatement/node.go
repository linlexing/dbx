package selectstatement

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/condition"
	"github.com/linlexing/dbx/ddb/parser/model"
	"github.com/linlexing/dbx/ddb/parser/selectelements"
	"github.com/linlexing/dbx/schema"
)

func ParserNode(val string) *model.NodeSelectStatement {
	if len(val) == 0 {
		return nil
	}
	//注释
	var vars map[string]interface{}
	val, vars = condition.ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectStatement()
	visitor := new(sqlSelectStatementVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*model.NodeSelectStatement)
}

// func ParseByContext(ctx parser.ISelectStatementContext) *model.NodeSelectStatement {
// 	visitor := new(SqlSelectStatementVisitorImpl)
// 	return visitor.Visit(ctx).(*model.NodeSelectStatement)
// }

// 两边自带括号
func SelectStatementString(node *model.NodeSelectStatement,
	fields map[string]schema.DataType, OuterTableName string,
	getview condition.GetUserConditionViewDefineFunc, buildComment bool) string {
	var tableSources []string
	for _, v := range node.TableSources {
		tableSources = append(tableSources, tableSourceString(v, fields, OuterTableName, getview, buildComment))
	}
	whereStr := node.WhereClause.WhereString(fields, OuterTableName, getview, buildComment)
	if len(whereStr) > 0 {
		whereStr = "WHERE " + whereStr
	}
	sql := strings.TrimSpace(fmt.Sprintf(`SELECT %s FROM %s %s %s`,
		selectelements.SelectElementsString(node.SelectElements),
		strings.Join(tableSources, ","),
		joinClauseString(node.JoinClause, fields, OuterTableName, getview, buildComment),
		whereStr,
	))
	return fmt.Sprintf(`(%s)`, sql)
}

func parserNodeJoin(val string) []*model.NodeJoinClause {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	// var vars map[string]interface{}
	// val, vars = condition.ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.JoinClause()
	visitor := new(sqlJoinClauseVisitorImpl)
	// visitor.vars = vars
	return visitor.Visit(tree).([]*model.NodeJoinClause)
}
func parseByJoinContext(ctx parser.IJoinClauseContext, vars map[string]interface{}) []*model.NodeJoinClause {
	visitor := new(sqlJoinClauseVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).([]*model.NodeJoinClause)
}
func joinClauseString(nodes []*model.NodeJoinClause, fields map[string]schema.DataType,
	outerTableName string, getview condition.GetUserConditionViewDefineFunc, buildComment bool) string {
	var joins []string
	for _, v := range nodes {
		joins = append(joins, fmt.Sprintf("%s %s ON %s",
			v.JoinType, tableSourceString(v.TableSource, fields, outerTableName, getview, buildComment),
			v.OnExpress.WhereString(fields, outerTableName, getview, buildComment)))
	}
	return strings.Join(joins, " ")
}

func parserNodeTableSources(val string) []*model.NodeTableSource {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	// val, vars = processComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.TableSources()
	visitor := new(sqlTableSourcesVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).([]*model.NodeTableSource)
}
func parseByTableSourcesContext(ctx parser.ITableSourcesContext, vars map[string]interface{}) []*model.NodeTableSource {
	visitor := new(sqlTableSourcesVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).([]*model.NodeTableSource)
}
func tableSourceString(node *model.NodeTableSource, fields map[string]schema.DataType,
	OuterTableName string, getview condition.GetUserConditionViewDefineFunc, buildComment bool) string {
	if len(node.Source.TableName) > 0 {
		return fmt.Sprintf("%s %s", node.Source.TableName, node.Alias)
	}
	return fmt.Sprintf("%s %s",
		SelectStatementString(node.Source.SelectStatement, fields, OuterTableName, getview, buildComment),
		node.Alias)
}
