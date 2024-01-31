package selectstatement

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/schema"
)

func ParserSelectNode(val string) *NodeSelectStatement {
	if len(val) == 0 {
		return nil
	}
	//注释
	var vars map[string]interface{}
	val, vars = ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectStatement()
	visitor := new(sqlSelectStatementVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*NodeSelectStatement)
}

func parseBySelectStatementContext(ctx parser.ISelectStatementContext, vars map[string]interface{}) *NodeSelectStatement {
	visitor := new(sqlSelectStatementVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).(*NodeSelectStatement)
}

func (node *NodeSelectStatement) SelectStatementString(transform bool) string {
	var fields map[string]schema.DataType
	if node.WhereClause != nil {
		fields = node.WhereClause.Fields
	}
	var sql string
	if len(node.UnionSelect) > 0 {
		var selects []string
		unionStr := " UNION "
		if node.UnionAll {
			unionStr = " UNION ALL "
		}
		for k := range node.UnionSelect {
			selects = append(selects, node.UnionSelect[k].SelectStatementString(transform))
		}
		sql = strings.Join(selects, unionStr)
		return sql
	}
	var tableSources []string
	var alias string
	for _, v := range node.TableSources {
		if len(alias) == 0 {
			alias = v.Alias
		}
		tableSources = append(tableSources, tableSourceString(v, node.GetView, transform))
	}
	whereStr := node.WhereClause.WhereString(fields, alias, nil, true, transform)
	if len(whereStr) > 0 {
		whereStr = "WHERE " + whereStr
	}
	sql = strings.TrimSpace(fmt.Sprintf(`SELECT %s FROM %s %s %s`,
		selectElementsString(node.SelectElements, transform),
		strings.Join(tableSources, ","),
		joinClauseString(node.JoinClause, node.GetView, transform),
		whereStr,
	))
	return sql
}

func parserNodeJoin(val string) []*NodeJoinClause {
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
	return visitor.Visit(tree).([]*NodeJoinClause)
}
func parseByJoinContext(ctx parser.IJoinClauseContext, vars map[string]interface{}) []*NodeJoinClause {
	visitor := new(sqlJoinClauseVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).([]*NodeJoinClause)
}
func joinClauseString(nodes []*NodeJoinClause, getview GetUserConditionViewDefineFunc, transform bool) string {
	var joins []string
	for _, v := range nodes {
		joins = append(joins, fmt.Sprintf("%s %s ON %s",
			v.JoinType, tableSourceString(v.TableSource, getview, transform),
			v.OnExpress.WhereString(nil, v.TableSource.Alias, nil, true, transform)))
	}
	return strings.Join(joins, " ")
}

func parserNodeTableSources(val string) []*NodeTableSource {
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
	return visitor.Visit(tree).([]*NodeTableSource)
}
func parseByTableSourcesContext(ctx parser.ITableSourcesContext, vars map[string]interface{}) []*NodeTableSource {
	visitor := new(sqlTableSourcesVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).([]*NodeTableSource)
}
func tableSourceString(node *NodeTableSource, getview GetUserConditionViewDefineFunc, transform bool) string {
	if len(node.Source.TableName) > 0 {
		if transform {
			str, err := getview(node.Source.TableName)
			if err != nil {
				log.Panic(err)
			}
			if len(str) > 0 {
				return fmt.Sprintf("(%s) %s", str, node.Alias)
			}
		}
		return fmt.Sprintf("%s %s", node.Source.TableName, node.Alias)
	}
	return fmt.Sprintf("(%s) %s",
		node.Source.SelectStatement.SelectStatementString(transform),
		node.Alias)
}
func parserNodeSelectelements(val string) *NodeSelectelements {
	if len(val) == 0 {
		return nil
	}
	var vars map[string]interface{}
	// val, _ = condition.ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectElements()
	visitor := new(sqlSelectelementsVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*NodeSelectelements)
}

func parseBySelectElementsContext(ctx parser.ISelectElementsContext, vars map[string]interface{}) *NodeSelectelements {
	visitor := new(sqlSelectelementsVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(ctx).(*NodeSelectelements)
}

func selectElementsString(node *NodeSelectelements, transform bool) string {
	var elements []string
	if node != nil {
		for _, v := range node.Elements {
			col := v.Express
			if v.Subquery != nil {
				col = "(" + v.Subquery.SelectStatementString(transform) + ")"
			}
			var as, alias string
			if len(v.ColumnName) > 0 {
				if len(v.TableAlias) > 0 {
					col = v.TableAlias + "." + v.ColumnName
				} else {
					col = v.ColumnName
				}
			}
			if len(v.As) > 0 {
				as = " " + v.As
			}
			if len(v.Alias) > 0 {
				alias = " " + v.Alias
			}
			elements = append(elements, fmt.Sprintf("%s%s%s", col, as, alias))
		}
	}
	return strings.Join(elements, ",")
}
