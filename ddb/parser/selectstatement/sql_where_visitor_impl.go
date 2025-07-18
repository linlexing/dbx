package selectstatement

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
)

var (
	// 找出所有不在单引号字符串内的注释内容
	// (?:[^']|'[^']')*?(/\*[^*]\*(?:[^/*][^*]\*)*/)
	//
	// Non-capturing group (?:[^']|'[^']')*?
	// *? matches the previous token between zero and unlimited times, as few times as possible, expanding as needed (lazy)
	// 1st Alternative [^']
	// Match a single character not present in the list below [^']
	// ' matches the character ' literally (case sensitive)
	// 2nd Alternative '[^']'
	// ' matches the character ' literally (case sensitive)
	// Match a single character not present in the list below [^']
	// * matches the previous token between zero and unlimited times, as many times as possible, giving back as needed (greedy)
	// ' matches the character ' literally (case sensitive)
	// ' matches the character ' literally (case sensitive)
	// 1st Capturing Group (/\*[^*]\*(?:[^/*][^*]\*)*/)
	// / matches the character / literally (case sensitive)
	// \* matches the character * literally (case sensitive)
	// Match a single character not present in the list below [^*]
	// * matches the previous token between zero and unlimited times, as many times as possible, giving back as needed (greedy)
	// * matches the character * literally (case sensitive)
	// \* matches the character * literally (case sensitive)
	// + matches the previous token between one and unlimited times, as many times as possible, giving back as needed (greedy)
	// Non-capturing group (?:[^/*][^*]\*)*
	// * matches the previous token between zero and unlimited times, as many times as possible, giving back as needed (greedy)
	// Match a single character not present in the list below [^/*]
	// Match a single character not present in the list below [^*]
	// \* matches the character * literally (case sensitive)
	// / matches the character / literally (case sensitive)
	regComment = regexp.MustCompile(`(?:[^']|'[^']*')*?(/\*[^*]*\*+(?:[^/*][^*]*\*+)*/)`)
	regExists  = regexp.MustCompile(`(?ms)^from\s+(.*)\s+on\s+(.*?)\s+where\s+(.*)$`)
	regInTable = regexp.MustCompile(`(?ms)^(.*)\s+in\s+(.*)\((.*)\)\s+where\s+(.*)$`)
	regCount   = regexp.MustCompile(`(?ms)^COUNT\(from\s+(.*)\s+on\s+(.*?)\s+where\s+(.*)\)\s+(.*)\s+(.*)$`)
	regQuote   = regexp.MustCompile(`^'([^']*)'$`)
	regInMini  = regexp.MustCompile(`(?ms)^(.*)\s+in\s+\((.*)\)$`)
)

const (
	commentPlainText = "/*PLAINTEXT*/"
	// commentNotPlainText = "/*NOT PLAINTEXT*/"
	//示例: from ABC on I.abc=O.abc and I.where=O.aaa where a=1 and b=c exist(select 1 from aab where 1=2)
	commentExists    = "/*EXISTS("
	commentNotExists = "/*NOT EXISTS("
	//示例: field in ABC(acd) where a=1 and b=c exist(select 1 from aab where 1=2)
	commentIn    = "/*IN("
	commentNotIn = "/*NOT IN("
	//示例:/*COUNT(from ABC on abc=O.abc and where=O.aaa where a=1 and b=c exist(select 1 from aab where 1=2)) > 0*/
	commentCount = "/*COUNT("
	//特殊的，临时被替换的动态node
	commentDynamicNode = "/*DYNAMIC-NODE*/"
	commentInMini      = "/*InMini("
	commentNotInMini   = "/*NotInMini("
)

// sqlWhereVisitorImpl 完成条件串的转换，只支持简单的 字段名 运算符 值 条件
type sqlWhereVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

// 去除单引号，如果有的话,没有单引号，将自动换成“,用于代表表达式
func decodeExprOrConst(expr parser.IExprContext) string {
	if expr1, ok := expr.(*parser.ExprContext); ok {
		val := expr1.Value()
		if val == nil {
			return "`" + expr.GetText() + "`"
		}
		if val1, ok := val.(*parser.ValueContext); ok {

			if txt := val1.TextLiteral(); txt != nil {
				str := txt.GetText()
				return strings.ReplaceAll(str[1:len(str)-1], "''", "'")
			}

		}

	}
	return expr.GetText()

}
func NewSqlWhereVisitorImpl() *sqlWhereVisitorImpl {
	return &sqlWhereVisitorImpl{
		SqlVisitor: &parser.BaseSqlVisitor{},
		vars:       make(map[string]interface{}),
	}
}

func (s *sqlWhereVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.WhereClauseContext:
		node := val.Accept(s).(*NodeCondition)
		node.Reduction()
		return node
	default:
		panic("not impl")
	}
}
func (s *sqlWhereVisitorImpl) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return ctx.LogicExpression().Accept(s)
}

//	func isColumn(expr parser.IExprContext) bool {
//		return expr.(*parser.ExprContext).ColumnName() != nil
//	}
//
// 识别SelectStatement
func processSubSelect(expr parser.IExprContext) *NodeCondition {
	if expr.SelectStatement() != nil {
		visitor := new(sqlSelectStatementVisitorImpl)
		subSelect := visitor.Visit(expr.SelectStatement()).(*NodeSelectStatement)
		return NewSubSelectNode(subSelect)
	}
	return nil
}

// 将运算符左边的表达式转换成node的name和func
func expr2NodeName(expr, expr2 parser.IExprContext) *NodeCondition {
	var fieldType schema.DataType
	if expr2 != nil {
		exprTxt := expr2.GetText()
		if !regQuote.MatchString(exprTxt) {
			if _, err := strconv.ParseFloat(exprTxt, 64); err == nil {
				fieldType = schema.TypeFloat
			}
			if _, err := strconv.Atoi(exprTxt); err == nil {
				fieldType = schema.TypeInt
			}
		}
	}
	subSelectNode := processSubSelect(expr)
	if subSelectNode != nil {
		return subSelectNode
	}
	if funcCall := expr.(*parser.ExprContext).FunctionCall(); funcCall != nil {

		switch tv := funcCall.(*parser.FunctionCallContext).CommonFunction().(type) {
		case *parser.CommonFunctionContext:
			// tv := fc.(*parser.CommonFunctionContext)
			funcName := strings.ToUpper(tv.FunctionName().GetText())
			exprList := tv.FunctionArg().(*parser.FunctionArgContext).AllExpr()
			args := make([]string, len(exprList))
			for i, one := range exprList {
				args[i] = one.GetText()
			}
			return NewFuncNode(args[0], funcName, args[1:])
		default:
			panic("invalid function " + tv.GetText())
		}
	}
	return &NodeCondition{Field: expr.GetText(), Type: fieldType, NodeType: ConditionNodeCondition}
}
func (s *sqlWhereVisitorImpl) VisitLogicExpression(ctx *parser.LogicExpressionContext) interface{} {
	return ParseLogicExpression(s, ctx, s.vars)
}
func ParseLogicExpression(s antlr.ParseTreeVisitor, ctx *parser.LogicExpressionContext, vars map[string]interface{}) interface{} {
	//not logicExpression
	if not, logicExpression1 := ctx.NOT(), ctx.LogicExpression(0); not != nil && logicExpression1 != nil {
		node := ParseLogicExpression(s, logicExpression1.(*parser.LogicExpressionContext), vars).(*NodeCondition)
		node.Reverse = true
		return node
	}
	//逻辑关系隔开的条件
	if logicExpression1, logicalOperator, logicExpression2 :=
		ctx.LogicExpression(0), ctx.GetLogicalOperator(), ctx.LogicExpression(1); logicExpression1 != nil && logicalOperator != nil && logicExpression2 != nil {
		var nodeType NodeType
		if ctx.AND() != nil {
			nodeType = ConditionNodeAnd
		} else {
			nodeType = ConditionNodeOr
		}
		return NewLogicNode(nodeType, []*NodeCondition{logicExpression1.Accept(s).(*NodeCondition),
			logicExpression2.Accept(s).(*NodeCondition)})
	}
	//两边表达式都是子查询，PlainNode
	if expr1, expr2 := ctx.Expr(0), ctx.Expr(1); expr1 != nil && expr2 != nil {
		subSelectNode1 := processSubSelect(expr1)
		subSelectNode2 := processSubSelect(expr2)
		if subSelectNode1 != nil && subSelectNode2 != nil {
			return NewPlainNode(getText(ctx))
		}
	}
	//运算符隔开的单个条件
	if expr1, operate, expr2 := ctx.Expr(0), ctx.ComparisonOperator(), ctx.Expr(1); expr1 != nil && operate != nil && expr2 != nil {
		node := expr2NodeName(expr1, expr2)
		//长度返回整型，特殊处理
		if strings.EqualFold(node.Func, "LENGTH") {
			var ope pageselect.Operator
			switch operate.GetText() {
			case "=":
				ope = pageselect.OperatorLengthEqu
			case ">":
				ope = pageselect.OperatorLengthGreaterThan
			case "<":
				ope = pageselect.OperatorLengthLessThan
			case "<=":
				ope = pageselect.OperatorLengthLessThanOrEqu
			case ">=":
				ope = pageselect.OperatorLengthGreaterThanOrEqu
			case "<>":
				ope = pageselect.OperatorLengthNotEqu
			default:
				panic("invalid length opereate " + operate.GetText())
			}
			return NewConditionNode(node.Field, ope, decodeExprOrConst(expr2), "")
		}
		// if funcCall := expr1.(*parser.ExprContext).FunctionCall(); funcCall != nil {

		// 	switch tv := funcCall.(*parser.FunctionCallContext).CommonFunction().(type) {
		// 	case *parser.CommonFunctionContext:
		// 		// tv := fc.(*parser.CommonFunctionContext)
		// 		funcName := strings.ToUpper(tv.FunctionName().GetText())
		// 		exprList := tv.FunctionArg().(*parser.FunctionArgContext).AllExpr()
		// 		switch funcName {
		// 		case "LENGTH":
		// 			return NewFuncNode(tv.FunctionArg().GetText(), funcName, nil, operate.GetText(), decodeExprOrConst(expr2), "")
		// 		}
		// 		return NewConditionNode(tv.FunctionArg().GetText(), ope, decodeExprOrConst(expr2), "")
		// 		if strings.ToUpper(tv.FunctionName().GetText()) == "LENGTH" {
		// 			if exprList := tv.FunctionArg().(*parser.FunctionArgContext).AllExpr(); len(exprList) == 1 &&
		// 				isColumn(exprList[0]) {
		// 				var ope pageselect.Operator
		// 				switch operate.GetText() {
		// 				case "=":
		// 					ope = pageselect.OperatorLengthEqu
		// 				case ">":
		// 					ope = pageselect.OperatorLengthGreaterThan
		// 				case "<":
		// 					ope = pageselect.OperatorLengthLessThan
		// 				case "<=":
		// 					ope = pageselect.OperatorLengthLessThanOrEqu
		// 				case ">=":
		// 					ope = pageselect.OperatorLengthGreaterThanOrEqu
		// 				case "<>":
		// 					ope = pageselect.OperatorLengthNotEqu
		// 				default:
		// 					panic("invalid length opereate " + operate.GetText())
		// 				}
		// 				return NewConditionNode(tv.FunctionArg().GetText(), ope, decodeExprOrConst(expr2), "")
		// 			}
		// 		}
		// 		panic("invalid function " + tv.FunctionName().GetText())

		// 	}
		// }
		// if !isColumn(expr1) {
		// 	return NewPlainNode(getText(ctx))
		// }
		var ope pageselect.Operator
		switch operate.GetText() {
		case "=":
			ope = pageselect.OperatorEqu
		case ">":
			ope = pageselect.OperatorGreaterThan
		case "<":
			ope = pageselect.OperatorLessThan
		case "<=":
			ope = pageselect.OperatorLessThanOrEqu
		case ">=":
			ope = pageselect.OperatorGreaterThanOrEqu
		case "<>":
			ope = pageselect.OperatorNotEqu
		case "~":
			ope = pageselect.OperatorRegexp
		case "!~":
			ope = pageselect.OperatorNotRegexp
		case "?[]":
			ope = pageselect.OperatorLikeArray
		case "!?[]":
			ope = pageselect.OperatorNotLikeArray
		case "?>[]":
			ope = pageselect.OperatorPrefixArray
		case "!?>[]":
			ope = pageselect.OperatorNotPrefixArray
		case "<?[]":
			ope = pageselect.OperatorSuffixArray
		case "!<?[]":
			ope = pageselect.OperatorNotSuffixArray
		default:
			panic("invalid opereate " + operate.GetText())
		}
		node.Operate = ope
		node.Value = decodeExprOrConst(expr2)

		return node

	}
	//BETWEEN
	if not, expr1, between, expr2, expr3 :=
		ctx.NOT(), ctx.Expr(0), ctx.BETWEEN(), ctx.Expr(1), ctx.Expr(2); expr1 != nil &&
		between != nil && expr2 != nil && expr3 != nil {
		node := expr2NodeName(expr1, expr2)
		if not != nil {
			node.Operate = pageselect.OperatorNotBetween
		} else {
			node.Operate = pageselect.OperatorBetween
		}
		node.Value = decodeExprOrConst(expr2)
		node.Value2 = decodeExprOrConst(expr3)
		return node

	}

	//IN/NOT IN Table
	if not, expr1, in, subSelect := ctx.NOT(), ctx.Expr(0), ctx.IN(),
		ctx.SelectStatement(); in != nil && expr1 != nil && subSelect != nil {
		visitor := new(sqlSelectStatementVisitorImpl)
		nodeSelectStatement := visitor.Visit(subSelect).(*NodeSelectStatement)
		return NewInSubSelectNode(expr1.GetText(), nodeSelectStatement, not != nil)
	}

	//IN/NOT IN
	if not, in, expr := ctx.NOT(), ctx.IN(), ctx.AllExpr(); in != nil && len(expr) > 1 {
		node := expr2NodeName(expr[0], expr[1])
		if not != nil {
			node.Operate = pageselect.OperatorNotIn
		} else {
			node.Operate = pageselect.OperatorIn
		}
		strs := []string{}
		//需要将in后面的表达式进行字面量的转换，去掉引号
		for _, one := range expr[1:] {
			strs = append(strs, decodeExprOrConst(one))
		}

		node.Value = strings.Join(strs, ",")
		return node
	}

	//LIKE/NOT LIKE
	if not, like, field, val :=
		ctx.NOT(), ctx.LIKE(), ctx.Expr(0), ctx.Expr(1); like != nil &&
		field != nil && val != nil {
		node := expr2NodeName(field, nil)

		str := decodeExprOrConst(val)
		var first, last byte
		if len(str) > 0 {
			first = str[0]
		}
		if len(str) > 1 {
			last = str[len(str)-1]
		}

		var ope pageselect.Operator
		var valStr string
		if not != nil {
			if first == '%' && last == '%' {
				ope = pageselect.OperatorNotLike
				valStr = str[1 : len(str)-1]
			} else if first == '%' {
				ope = pageselect.OperatorNotSuffix
				valStr = str[1:]
			} else if last == '%' {
				ope = pageselect.OperatorNotPrefix
				valStr = str[:len(str)-1]
			} else {
				ope = pageselect.OperatorNotEqu
				valStr = str
			}
		} else {
			if first == '%' && last == '%' {
				ope = pageselect.OperatorLike
				valStr = str[1 : len(str)-1]
			} else if first == '%' {
				ope = pageselect.OperatorSuffix
				valStr = str[1:]
			} else if last == '%' {
				ope = pageselect.OperatorPrefix
				valStr = str[:len(str)-1]
			} else {
				ope = pageselect.OperatorEqu
				valStr = str
			}
		}
		node.Operate = ope
		node.Value = valStr
		return node
	}
	//IS NULL/IS NOT NULL
	if is, not, null, field :=
		ctx.IS(), ctx.NOT(), ctx.NULL(), ctx.Expr(0); is != nil && null != nil && field != nil {
		node := expr2NodeName(field, nil)

		if not != nil {
			node.Operate = pageselect.OperatorIsNotNull

		} else {
			node.Operate = pageselect.OperatorIsNull
		}
		return node

	}
	//EXISTS (selectStatement)
	if not, exists, subSelect := ctx.NOT(), ctx.EXISTS(), ctx.SelectStatement(); exists != nil && subSelect != nil {
		visitor := new(sqlSelectStatementVisitorImpl)
		nodeSelectStatement := visitor.Visit(subSelect).(*NodeSelectStatement)
		return NewExistsSubSelectNode(nodeSelectStatement, not != nil)
	}
	//动态node '(' logicExpression ')'
	if comment, left, right, logicExpr := ctx.COMMENT(), ctx.GetLeftBracket(), ctx.GetRightBracket(),
		ctx.LogicExpression(0); left != nil && right != nil && logicExpr != nil {
		if comment != nil && comment.GetText() == commentDynamicNode {
			id := strings.Split(logicExpr.GetText(), "=")[1]
			return vars[id].(*NodeCondition)
		}
		return logicExpr.Accept(s)
	}
	return NewPlainNode(getText(ctx))
	// panic("不支持的表达式" + ctx.GetText())
}
func getText(node antlr.ParseTree) string {
	if node.GetChildCount() == 0 {
		return node.GetText()
	}
	list := []string{}
	for _, one := range node.GetChildren() {
		list = append(list, getText(one.(antlr.ParseTree)))
	}
	return strings.Join(list, " ")
}
