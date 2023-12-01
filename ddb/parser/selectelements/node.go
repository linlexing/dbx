package selectelements

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

func ParserNode(val string) *model.NodeSelectelements {
	if len(val) == 0 {
		return nil
	}
	//先进行注释的识别
	var vars map[string]interface{}
	//注释
	// val, vars = processComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectElements()
	visitor := new(SqlSelectelementsVisitorImpl)
	visitor.vars = vars
	return visitor.Visit(tree).(*model.NodeSelectelements)
}

func ParseByContext(ctx parser.ISelectElementsContext) *model.NodeSelectelements {
	visitor := new(SqlSelectelementsVisitorImpl)
	return visitor.Visit(ctx).(*model.NodeSelectelements)
}

func NewElement(tableAlias, columnName, express, as, alias string) *model.Element {
	return &model.Element{
		TableAlias: tableAlias,
		ColumnName: columnName,
		Express:    express,
		As:         as,
		Alias:      alias,
	}
}
func SelectElementsString(node *model.NodeSelectelements) string {
	if node.NodeType == model.NodeStar {
		return "*"
	}
	var elements []string
	for _, v := range node.Elements {
		col := v.Express
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
	return strings.Join(elements, ",")
}

// 处理注释，识别关联查询并生成node列表
// vars [dyna1] *Node
func processComment(define string) (rev string, vars map[string]interface{}) {
	wait := define
	vars = map[string]interface{}{}
	// iDynamic := 0
	// addDynamicNode := func(node *Node) string {
	// 	iDynamic++
	// 	id := fmt.Sprintf("dyna%d", iDynamic)
	// 	vars[id] = node
	// 	return fmt.Sprintf("%s(id=%s)", commentDynamicNode, id)
	// }
	for positions := regComment.FindStringSubmatchIndex(
		wait); len(positions) == 4; positions = regComment.FindStringSubmatchIndex(wait) {

		// comment := wait[positions[2]:positions[3]]
		rev += wait[:positions[2]] //注释之前的截断到返回值中
		// if comment == commentPlainText {
		// 	afterText := strings.TrimSpace(wait[positions[3]:])
		// left, right := findBracketExpr(afterText)
		// 	node := processPlainText(left)
		// 	if node != nil {
		// 		rev += addDynamicNode(node)
		// 	}
		// 	//光有注释，没有后面内容的，注释被清除
		// 	wait = right
		// 	continue
		// }
		// if strings.HasPrefix(comment, commentCount) {
		// 	_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
		// 	rev += addDynamicNode(processCount(comment))
		// 	continue
		// }
		// if strings.HasPrefix(comment, commentExists) || strings.HasPrefix(comment, commentNotExists) {
		// 	_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
		// 	rev += addDynamicNode(processExists(comment))
		// 	continue
		// }
		// if strings.HasPrefix(comment, commentIn) || strings.HasPrefix(comment, commentNotIn) {
		// 	_, wait = findBracketExpr(strings.TrimSpace(wait[positions[3]:]))
		// 	rev += addDynamicNode(processIn(comment))
		// 	continue
		// }
		//普通的注释，不做处理，纳入结果中
		// rev += comment
		wait = wait[positions[3]:]
	}
	rev += wait
	return
}
