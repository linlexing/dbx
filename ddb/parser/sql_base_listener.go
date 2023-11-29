// Code generated from Sql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSqlListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlListener struct{}

var _ SqlListener = &BaseSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseSqlListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseSqlListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlListener) ExitTableName(ctx *TableNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseSqlListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseSqlListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseSqlListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseSqlListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSqlListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSqlListener) ExitAlias(ctx *AliasContext) {}

// EnterJoin is called when production join is entered.
func (s *BaseSqlListener) EnterJoin(ctx *JoinContext) {}

// ExitJoin is called when production join is exited.
func (s *BaseSqlListener) ExitJoin(ctx *JoinContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseSqlListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseSqlListener) ExitUnion(ctx *UnionContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterTextLiteral is called when production textLiteral is entered.
func (s *BaseSqlListener) EnterTextLiteral(ctx *TextLiteralContext) {}

// ExitTextLiteral is called when production textLiteral is exited.
func (s *BaseSqlListener) ExitTextLiteral(ctx *TextLiteralContext) {}

// EnterBind_variables is called when production bind_variables is entered.
func (s *BaseSqlListener) EnterBind_variables(ctx *Bind_variablesContext) {}

// ExitBind_variables is called when production bind_variables is exited.
func (s *BaseSqlListener) ExitBind_variables(ctx *Bind_variablesContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseSqlListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseSqlListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseSqlListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseSqlListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectElement is called when production selectElement is entered.
func (s *BaseSqlListener) EnterSelectElement(ctx *SelectElementContext) {}

// ExitSelectElement is called when production selectElement is exited.
func (s *BaseSqlListener) ExitSelectElement(ctx *SelectElementContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSqlListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSqlListener) ExitExpr(ctx *ExprContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSqlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSqlListener) ExitValue(ctx *ValueContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSqlListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSqlListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregateFunction is called when production aggregateFunction is entered.
func (s *BaseSqlListener) EnterAggregateFunction(ctx *AggregateFunctionContext) {}

// ExitAggregateFunction is called when production aggregateFunction is exited.
func (s *BaseSqlListener) ExitAggregateFunction(ctx *AggregateFunctionContext) {}

// EnterCommonFunction is called when production commonFunction is entered.
func (s *BaseSqlListener) EnterCommonFunction(ctx *CommonFunctionContext) {}

// ExitCommonFunction is called when production commonFunction is exited.
func (s *BaseSqlListener) ExitCommonFunction(ctx *CommonFunctionContext) {}

// EnterFunctionArg is called when production functionArg is entered.
func (s *BaseSqlListener) EnterFunctionArg(ctx *FunctionArgContext) {}

// ExitFunctionArg is called when production functionArg is exited.
func (s *BaseSqlListener) ExitFunctionArg(ctx *FunctionArgContext) {}

// EnterTableSources is called when production tableSources is entered.
func (s *BaseSqlListener) EnterTableSources(ctx *TableSourcesContext) {}

// ExitTableSources is called when production tableSources is exited.
func (s *BaseSqlListener) ExitTableSources(ctx *TableSourcesContext) {}

// EnterTableSource is called when production tableSource is entered.
func (s *BaseSqlListener) EnterTableSource(ctx *TableSourceContext) {}

// ExitTableSource is called when production tableSource is exited.
func (s *BaseSqlListener) ExitTableSource(ctx *TableSourceContext) {}

// EnterJoinClause is called when production joinClause is entered.
func (s *BaseSqlListener) EnterJoinClause(ctx *JoinClauseContext) {}

// ExitJoinClause is called when production joinClause is exited.
func (s *BaseSqlListener) ExitJoinClause(ctx *JoinClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSqlListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSqlListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterLogicExpression is called when production logicExpression is entered.
func (s *BaseSqlListener) EnterLogicExpression(ctx *LogicExpressionContext) {}

// ExitLogicExpression is called when production logicExpression is exited.
func (s *BaseSqlListener) ExitLogicExpression(ctx *LogicExpressionContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSqlListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSqlListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupByItem is called when production groupByItem is entered.
func (s *BaseSqlListener) EnterGroupByItem(ctx *GroupByItemContext) {}

// ExitGroupByItem is called when production groupByItem is exited.
func (s *BaseSqlListener) ExitGroupByItem(ctx *GroupByItemContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSqlListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSqlListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSqlListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSqlListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByExpression is called when production orderByExpression is entered.
func (s *BaseSqlListener) EnterOrderByExpression(ctx *OrderByExpressionContext) {}

// ExitOrderByExpression is called when production orderByExpression is exited.
func (s *BaseSqlListener) ExitOrderByExpression(ctx *OrderByExpressionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSqlListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSqlListener) ExitLimitClause(ctx *LimitClauseContext) {}
