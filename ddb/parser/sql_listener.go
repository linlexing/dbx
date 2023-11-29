// Code generated from Sql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SqlListener is a complete listener for a parse tree produced by SqlParser.
type SqlListener interface {
	antlr.ParseTreeListener

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterJoin is called when entering the join production.
	EnterJoin(c *JoinContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterTextLiteral is called when entering the textLiteral production.
	EnterTextLiteral(c *TextLiteralContext)

	// EnterBind_variables is called when entering the bind_variables production.
	EnterBind_variables(c *Bind_variablesContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectElements is called when entering the selectElements production.
	EnterSelectElements(c *SelectElementsContext)

	// EnterSelectElement is called when entering the selectElement production.
	EnterSelectElement(c *SelectElementContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAggregateFunction is called when entering the aggregateFunction production.
	EnterAggregateFunction(c *AggregateFunctionContext)

	// EnterCommonFunction is called when entering the commonFunction production.
	EnterCommonFunction(c *CommonFunctionContext)

	// EnterFunctionArg is called when entering the functionArg production.
	EnterFunctionArg(c *FunctionArgContext)

	// EnterTableSources is called when entering the tableSources production.
	EnterTableSources(c *TableSourcesContext)

	// EnterTableSource is called when entering the tableSource production.
	EnterTableSource(c *TableSourceContext)

	// EnterJoinClause is called when entering the joinClause production.
	EnterJoinClause(c *JoinClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterLogicExpression is called when entering the logicExpression production.
	EnterLogicExpression(c *LogicExpressionContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupByItem is called when entering the groupByItem production.
	EnterGroupByItem(c *GroupByItemContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByExpression is called when entering the orderByExpression production.
	EnterOrderByExpression(c *OrderByExpressionContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitJoin is called when exiting the join production.
	ExitJoin(c *JoinContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitTextLiteral is called when exiting the textLiteral production.
	ExitTextLiteral(c *TextLiteralContext)

	// ExitBind_variables is called when exiting the bind_variables production.
	ExitBind_variables(c *Bind_variablesContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectElements is called when exiting the selectElements production.
	ExitSelectElements(c *SelectElementsContext)

	// ExitSelectElement is called when exiting the selectElement production.
	ExitSelectElement(c *SelectElementContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAggregateFunction is called when exiting the aggregateFunction production.
	ExitAggregateFunction(c *AggregateFunctionContext)

	// ExitCommonFunction is called when exiting the commonFunction production.
	ExitCommonFunction(c *CommonFunctionContext)

	// ExitFunctionArg is called when exiting the functionArg production.
	ExitFunctionArg(c *FunctionArgContext)

	// ExitTableSources is called when exiting the tableSources production.
	ExitTableSources(c *TableSourcesContext)

	// ExitTableSource is called when exiting the tableSource production.
	ExitTableSource(c *TableSourceContext)

	// ExitJoinClause is called when exiting the joinClause production.
	ExitJoinClause(c *JoinClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitLogicExpression is called when exiting the logicExpression production.
	ExitLogicExpression(c *LogicExpressionContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupByItem is called when exiting the groupByItem production.
	ExitGroupByItem(c *GroupByItemContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByExpression is called when exiting the orderByExpression production.
	ExitOrderByExpression(c *OrderByExpressionContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)
}
