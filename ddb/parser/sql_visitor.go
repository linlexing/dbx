// Code generated from Sql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by SqlParser.
type SqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SqlParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by SqlParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by SqlParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by SqlParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by SqlParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by SqlParser#join.
	VisitJoin(ctx *JoinContext) interface{}

	// Visit a parse tree produced by SqlParser#union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by SqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#textLiteral.
	VisitTextLiteral(ctx *TextLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#bind_variables.
	VisitBind_variables(ctx *Bind_variablesContext) interface{}

	// Visit a parse tree produced by SqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) interface{}

	// Visit a parse tree produced by SqlParser#selectElement.
	VisitSelectElement(ctx *SelectElementContext) interface{}

	// Visit a parse tree produced by SqlParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SqlParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by SqlParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SqlParser#aggregateFunction.
	VisitAggregateFunction(ctx *AggregateFunctionContext) interface{}

	// Visit a parse tree produced by SqlParser#commonFunction.
	VisitCommonFunction(ctx *CommonFunctionContext) interface{}

	// Visit a parse tree produced by SqlParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}

	// Visit a parse tree produced by SqlParser#tableSources.
	VisitTableSources(ctx *TableSourcesContext) interface{}

	// Visit a parse tree produced by SqlParser#tableSource.
	VisitTableSource(ctx *TableSourceContext) interface{}

	// Visit a parse tree produced by SqlParser#joinClause.
	VisitJoinClause(ctx *JoinClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#logicExpression.
	VisitLogicExpression(ctx *LogicExpressionContext) interface{}

	// Visit a parse tree produced by SqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SqlParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) interface{}

	// Visit a parse tree produced by SqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) interface{}

	// Visit a parse tree produced by SqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}
}
