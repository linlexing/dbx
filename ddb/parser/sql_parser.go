// Code generated from Sql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Sql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SqlParser struct {
	*antlr.BaseParser
}

var sqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlParserInit() {
	staticData := &sqlParserStaticData
	staticData.literalNames = []string{
		"", "'*'", "','", "'/'", "'+'", "'-'", "'||'", "'('", "')'", "'='",
		"'>'", "'<'", "'>='", "'<='", "'<>'", "'~'", "'!~'", "'?[]'", "'!?[]'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "AS", "SELECT", "FROM", "MAX", "SUM", "AVG", "MIN", "COUNT",
		"DISTINCT", "WHERE", "GROUP", "BY", "ORDER", "NULLS", "FIRST", "LAST",
		"HAVING", "NOT", "IS", "BETWEEN", "AND", "IN", "NULL", "OR", "ASC",
		"DESC", "LIMIT", "OFFSET", "LIKE", "EXISTS", "CAST", "INNER", "LEFT",
		"RIGHT", "JOIN", "ON", "UNION", "ALL", "DECIMAL_LITERAL", "ID", "TEXT_STRING",
		"TEXT_ALIAS", "BIND_VARIABLE", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"columnName", "tableName", "typeName", "functionName", "alias", "join",
		"union", "decimalLiteral", "textLiteral", "bind_variables", "selectStatement",
		"selectElements", "selectElement", "expr", "value", "functionCall",
		"aggregateFunction", "commonFunction", "functionArg", "tableSources",
		"tableSource", "joinClause", "whereClause", "logicExpression", "comparisonOperator",
		"groupByClause", "groupByItem", "havingClause", "orderByClause", "orderByExpression",
		"limitClause",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 369, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 3, 6, 78, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 93, 8, 10, 1, 10, 3, 10,
		96, 8, 10, 1, 10, 3, 10, 99, 8, 10, 1, 10, 3, 10, 102, 8, 10, 1, 10, 3,
		10, 105, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 111, 8, 10, 10, 10,
		12, 10, 114, 9, 10, 1, 11, 1, 11, 3, 11, 118, 8, 11, 1, 11, 1, 11, 5, 11,
		122, 8, 11, 10, 11, 12, 11, 125, 9, 11, 1, 12, 1, 12, 3, 12, 129, 8, 12,
		1, 12, 3, 12, 132, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 142, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 5, 13, 153, 8, 13, 10, 13, 12, 13, 156, 9, 13, 1,
		14, 1, 14, 1, 14, 3, 14, 161, 8, 14, 1, 15, 1, 15, 3, 15, 165, 8, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 176,
		8, 16, 1, 16, 3, 16, 179, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 195,
		8, 17, 1, 18, 1, 18, 1, 18, 5, 18, 200, 8, 18, 10, 18, 12, 18, 203, 9,
		18, 1, 19, 1, 19, 3, 19, 207, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 212, 8,
		19, 5, 19, 214, 8, 19, 10, 19, 12, 19, 217, 9, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 3, 20, 224, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5,
		21, 231, 8, 21, 10, 21, 12, 21, 234, 9, 21, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 246, 8, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 255, 8, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 5, 23, 262, 8, 23, 10, 23, 12, 23, 265, 9, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 3, 23, 271, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 280, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 288, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 298, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 3, 23, 306, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23,
		314, 8, 23, 10, 23, 12, 23, 317, 9, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 5, 25, 326, 8, 25, 10, 25, 12, 25, 329, 9, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 341,
		8, 28, 10, 28, 12, 28, 344, 9, 28, 1, 29, 1, 29, 3, 29, 348, 8, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 3, 29, 354, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		3, 30, 360, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 367, 8, 30,
		1, 30, 0, 3, 20, 26, 46, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 0, 7, 2, 0, 58, 58, 60, 60, 1, 0, 50, 52, 2, 0, 1, 1, 3, 3, 1, 0, 4,
		5, 1, 0, 22, 25, 1, 0, 9, 18, 1, 0, 43, 44, 391, 0, 62, 1, 0, 0, 0, 2,
		64, 1, 0, 0, 0, 4, 66, 1, 0, 0, 0, 6, 68, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0,
		10, 72, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14, 79, 1, 0, 0, 0, 16, 81, 1,
		0, 0, 0, 18, 83, 1, 0, 0, 0, 20, 85, 1, 0, 0, 0, 22, 117, 1, 0, 0, 0, 24,
		126, 1, 0, 0, 0, 26, 141, 1, 0, 0, 0, 28, 160, 1, 0, 0, 0, 30, 164, 1,
		0, 0, 0, 32, 178, 1, 0, 0, 0, 34, 194, 1, 0, 0, 0, 36, 196, 1, 0, 0, 0,
		38, 204, 1, 0, 0, 0, 40, 223, 1, 0, 0, 0, 42, 232, 1, 0, 0, 0, 44, 235,
		1, 0, 0, 0, 46, 305, 1, 0, 0, 0, 48, 318, 1, 0, 0, 0, 50, 320, 1, 0, 0,
		0, 52, 330, 1, 0, 0, 0, 54, 332, 1, 0, 0, 0, 56, 335, 1, 0, 0, 0, 58, 345,
		1, 0, 0, 0, 60, 355, 1, 0, 0, 0, 62, 63, 5, 58, 0, 0, 63, 1, 1, 0, 0, 0,
		64, 65, 5, 58, 0, 0, 65, 3, 1, 0, 0, 0, 66, 67, 5, 58, 0, 0, 67, 5, 1,
		0, 0, 0, 68, 69, 5, 58, 0, 0, 69, 7, 1, 0, 0, 0, 70, 71, 7, 0, 0, 0, 71,
		9, 1, 0, 0, 0, 72, 73, 7, 1, 0, 0, 73, 74, 5, 53, 0, 0, 74, 11, 1, 0, 0,
		0, 75, 77, 5, 55, 0, 0, 76, 78, 5, 56, 0, 0, 77, 76, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 13, 1, 0, 0, 0, 79, 80, 5, 57, 0, 0, 80, 15, 1, 0, 0, 0,
		81, 82, 5, 59, 0, 0, 82, 17, 1, 0, 0, 0, 83, 84, 5, 61, 0, 0, 84, 19, 1,
		0, 0, 0, 85, 86, 6, 10, -1, 0, 86, 87, 5, 20, 0, 0, 87, 88, 3, 22, 11,
		0, 88, 89, 5, 21, 0, 0, 89, 90, 3, 38, 19, 0, 90, 92, 3, 42, 21, 0, 91,
		93, 3, 44, 22, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0,
		0, 0, 94, 96, 3, 50, 25, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		98, 1, 0, 0, 0, 97, 99, 3, 54, 27, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 99, 101, 1, 0, 0, 0, 100, 102, 3, 56, 28, 0, 101, 100, 1, 0, 0, 0,
		101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 105, 3, 60, 30, 0, 104,
		103, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 112, 1, 0, 0, 0, 106, 107,
		10, 1, 0, 0, 107, 108, 3, 12, 6, 0, 108, 109, 3, 20, 10, 2, 109, 111, 1,
		0, 0, 0, 110, 106, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0,
		0, 112, 113, 1, 0, 0, 0, 113, 21, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115,
		118, 5, 1, 0, 0, 116, 118, 3, 24, 12, 0, 117, 115, 1, 0, 0, 0, 117, 116,
		1, 0, 0, 0, 118, 123, 1, 0, 0, 0, 119, 120, 5, 2, 0, 0, 120, 122, 3, 24,
		12, 0, 121, 119, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0,
		123, 124, 1, 0, 0, 0, 124, 23, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 131,
		3, 26, 13, 0, 127, 129, 5, 19, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1,
		0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 3, 8, 4, 0, 131, 128, 1, 0, 0,
		0, 131, 132, 1, 0, 0, 0, 132, 25, 1, 0, 0, 0, 133, 134, 6, 13, -1, 0, 134,
		142, 3, 0, 0, 0, 135, 142, 3, 30, 15, 0, 136, 142, 3, 28, 14, 0, 137, 138,
		5, 7, 0, 0, 138, 139, 3, 26, 13, 0, 139, 140, 5, 8, 0, 0, 140, 142, 1,
		0, 0, 0, 141, 133, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 136, 1, 0, 0,
		0, 141, 137, 1, 0, 0, 0, 142, 154, 1, 0, 0, 0, 143, 144, 10, 4, 0, 0, 144,
		145, 7, 2, 0, 0, 145, 153, 3, 26, 13, 5, 146, 147, 10, 3, 0, 0, 147, 148,
		7, 3, 0, 0, 148, 153, 3, 26, 13, 4, 149, 150, 10, 2, 0, 0, 150, 151, 5,
		6, 0, 0, 151, 153, 3, 26, 13, 3, 152, 143, 1, 0, 0, 0, 152, 146, 1, 0,
		0, 0, 152, 149, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 27, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 161,
		3, 14, 7, 0, 158, 161, 3, 16, 8, 0, 159, 161, 3, 18, 9, 0, 160, 157, 1,
		0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 29, 1, 0, 0,
		0, 162, 165, 3, 32, 16, 0, 163, 165, 3, 34, 17, 0, 164, 162, 1, 0, 0, 0,
		164, 163, 1, 0, 0, 0, 165, 31, 1, 0, 0, 0, 166, 167, 7, 4, 0, 0, 167, 168,
		5, 7, 0, 0, 168, 169, 3, 36, 18, 0, 169, 170, 5, 8, 0, 0, 170, 179, 1,
		0, 0, 0, 171, 172, 5, 26, 0, 0, 172, 175, 5, 7, 0, 0, 173, 176, 5, 1, 0,
		0, 174, 176, 3, 36, 18, 0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 179, 5, 8, 0, 0, 178, 166, 1, 0, 0, 0, 178,
		171, 1, 0, 0, 0, 179, 33, 1, 0, 0, 0, 180, 181, 3, 6, 3, 0, 181, 182, 5,
		7, 0, 0, 182, 183, 3, 36, 18, 0, 183, 184, 5, 8, 0, 0, 184, 195, 1, 0,
		0, 0, 185, 186, 5, 27, 0, 0, 186, 195, 3, 36, 18, 0, 187, 188, 5, 49, 0,
		0, 188, 189, 5, 7, 0, 0, 189, 190, 3, 36, 18, 0, 190, 191, 5, 19, 0, 0,
		191, 192, 3, 4, 2, 0, 192, 193, 5, 8, 0, 0, 193, 195, 1, 0, 0, 0, 194,
		180, 1, 0, 0, 0, 194, 185, 1, 0, 0, 0, 194, 187, 1, 0, 0, 0, 195, 35, 1,
		0, 0, 0, 196, 201, 3, 26, 13, 0, 197, 198, 5, 2, 0, 0, 198, 200, 3, 26,
		13, 0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0,
		201, 202, 1, 0, 0, 0, 202, 37, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 206,
		3, 40, 20, 0, 205, 207, 3, 8, 4, 0, 206, 205, 1, 0, 0, 0, 206, 207, 1,
		0, 0, 0, 207, 215, 1, 0, 0, 0, 208, 209, 5, 2, 0, 0, 209, 211, 3, 40, 20,
		0, 210, 212, 3, 8, 4, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212,
		214, 1, 0, 0, 0, 213, 208, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213,
		1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 39, 1, 0, 0, 0, 217, 215, 1, 0,
		0, 0, 218, 224, 3, 2, 1, 0, 219, 220, 5, 7, 0, 0, 220, 221, 3, 20, 10,
		0, 221, 222, 5, 8, 0, 0, 222, 224, 1, 0, 0, 0, 223, 218, 1, 0, 0, 0, 223,
		219, 1, 0, 0, 0, 224, 41, 1, 0, 0, 0, 225, 226, 3, 10, 5, 0, 226, 227,
		3, 38, 19, 0, 227, 228, 5, 54, 0, 0, 228, 229, 3, 46, 23, 0, 229, 231,
		1, 0, 0, 0, 230, 225, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0,
		0, 0, 232, 233, 1, 0, 0, 0, 233, 43, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0,
		235, 236, 5, 28, 0, 0, 236, 237, 3, 46, 23, 0, 237, 45, 1, 0, 0, 0, 238,
		239, 6, 23, -1, 0, 239, 240, 3, 26, 13, 0, 240, 241, 3, 48, 24, 0, 241,
		242, 3, 26, 13, 0, 242, 306, 1, 0, 0, 0, 243, 245, 3, 26, 13, 0, 244, 246,
		5, 36, 0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0,
		0, 0, 247, 248, 5, 38, 0, 0, 248, 249, 3, 26, 13, 0, 249, 250, 5, 39, 0,
		0, 250, 251, 3, 26, 13, 0, 251, 306, 1, 0, 0, 0, 252, 254, 3, 26, 13, 0,
		253, 255, 5, 36, 0, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		256, 1, 0, 0, 0, 256, 257, 5, 40, 0, 0, 257, 258, 5, 7, 0, 0, 258, 263,
		3, 26, 13, 0, 259, 260, 5, 2, 0, 0, 260, 262, 3, 26, 13, 0, 261, 259, 1,
		0, 0, 0, 262, 265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0,
		0, 264, 266, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 266, 267, 5, 8, 0, 0, 267,
		306, 1, 0, 0, 0, 268, 270, 3, 26, 13, 0, 269, 271, 5, 36, 0, 0, 270, 269,
		1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 5, 40,
		0, 0, 273, 274, 5, 7, 0, 0, 274, 275, 3, 20, 10, 0, 275, 276, 5, 8, 0,
		0, 276, 306, 1, 0, 0, 0, 277, 279, 3, 26, 13, 0, 278, 280, 5, 36, 0, 0,
		279, 278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		282, 5, 47, 0, 0, 282, 283, 3, 26, 13, 0, 283, 306, 1, 0, 0, 0, 284, 285,
		3, 26, 13, 0, 285, 287, 5, 37, 0, 0, 286, 288, 5, 36, 0, 0, 287, 286, 1,
		0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 5, 41, 0,
		0, 290, 306, 1, 0, 0, 0, 291, 292, 5, 48, 0, 0, 292, 293, 5, 7, 0, 0, 293,
		294, 3, 20, 10, 0, 294, 295, 5, 8, 0, 0, 295, 306, 1, 0, 0, 0, 296, 298,
		5, 62, 0, 0, 297, 296, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 1, 0,
		0, 0, 299, 300, 5, 7, 0, 0, 300, 301, 3, 46, 23, 0, 301, 302, 5, 8, 0,
		0, 302, 306, 1, 0, 0, 0, 303, 304, 5, 36, 0, 0, 304, 306, 3, 46, 23, 3,
		305, 238, 1, 0, 0, 0, 305, 243, 1, 0, 0, 0, 305, 252, 1, 0, 0, 0, 305,
		268, 1, 0, 0, 0, 305, 277, 1, 0, 0, 0, 305, 284, 1, 0, 0, 0, 305, 291,
		1, 0, 0, 0, 305, 297, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 315, 1, 0,
		0, 0, 307, 308, 10, 2, 0, 0, 308, 309, 5, 39, 0, 0, 309, 314, 3, 46, 23,
		3, 310, 311, 10, 1, 0, 0, 311, 312, 5, 42, 0, 0, 312, 314, 3, 46, 23, 2,
		313, 307, 1, 0, 0, 0, 313, 310, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315,
		313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 47, 1, 0, 0, 0, 317, 315, 1,
		0, 0, 0, 318, 319, 7, 5, 0, 0, 319, 49, 1, 0, 0, 0, 320, 321, 5, 29, 0,
		0, 321, 322, 5, 30, 0, 0, 322, 327, 3, 52, 26, 0, 323, 324, 5, 2, 0, 0,
		324, 326, 3, 52, 26, 0, 325, 323, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 51, 1, 0, 0, 0, 329, 327, 1,
		0, 0, 0, 330, 331, 3, 26, 13, 0, 331, 53, 1, 0, 0, 0, 332, 333, 5, 35,
		0, 0, 333, 334, 3, 46, 23, 0, 334, 55, 1, 0, 0, 0, 335, 336, 5, 31, 0,
		0, 336, 337, 5, 30, 0, 0, 337, 342, 3, 58, 29, 0, 338, 339, 5, 2, 0, 0,
		339, 341, 3, 58, 29, 0, 340, 338, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342,
		340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 57, 1, 0, 0, 0, 344, 342, 1,
		0, 0, 0, 345, 347, 3, 26, 13, 0, 346, 348, 7, 6, 0, 0, 347, 346, 1, 0,
		0, 0, 347, 348, 1, 0, 0, 0, 348, 353, 1, 0, 0, 0, 349, 350, 5, 32, 0, 0,
		350, 354, 5, 33, 0, 0, 351, 352, 5, 32, 0, 0, 352, 354, 5, 34, 0, 0, 353,
		349, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 59, 1,
		0, 0, 0, 355, 366, 5, 45, 0, 0, 356, 357, 3, 14, 7, 0, 357, 358, 5, 2,
		0, 0, 358, 360, 1, 0, 0, 0, 359, 356, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0,
		360, 361, 1, 0, 0, 0, 361, 367, 3, 14, 7, 0, 362, 363, 3, 14, 7, 0, 363,
		364, 5, 46, 0, 0, 364, 365, 3, 14, 7, 0, 365, 367, 1, 0, 0, 0, 366, 359,
		1, 0, 0, 0, 366, 362, 1, 0, 0, 0, 367, 61, 1, 0, 0, 0, 41, 77, 92, 95,
		98, 101, 104, 112, 117, 123, 128, 131, 141, 152, 154, 160, 164, 175, 178,
		194, 201, 206, 211, 215, 223, 232, 245, 254, 263, 270, 279, 287, 297, 305,
		313, 315, 327, 342, 347, 353, 359, 366,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SqlParserInit initializes any static state used to implement SqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SqlParserInit() {
	staticData := &sqlParserStaticData
	staticData.once.Do(sqlParserInit)
}

// NewSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSqlParser(input antlr.TokenStream) *SqlParser {
	SqlParserInit()
	this := new(SqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sql.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF             = antlr.TokenEOF
	SqlParserT__0            = 1
	SqlParserT__1            = 2
	SqlParserT__2            = 3
	SqlParserT__3            = 4
	SqlParserT__4            = 5
	SqlParserT__5            = 6
	SqlParserT__6            = 7
	SqlParserT__7            = 8
	SqlParserT__8            = 9
	SqlParserT__9            = 10
	SqlParserT__10           = 11
	SqlParserT__11           = 12
	SqlParserT__12           = 13
	SqlParserT__13           = 14
	SqlParserT__14           = 15
	SqlParserT__15           = 16
	SqlParserT__16           = 17
	SqlParserT__17           = 18
	SqlParserAS              = 19
	SqlParserSELECT          = 20
	SqlParserFROM            = 21
	SqlParserMAX             = 22
	SqlParserSUM             = 23
	SqlParserAVG             = 24
	SqlParserMIN             = 25
	SqlParserCOUNT           = 26
	SqlParserDISTINCT        = 27
	SqlParserWHERE           = 28
	SqlParserGROUP           = 29
	SqlParserBY              = 30
	SqlParserORDER           = 31
	SqlParserNULLS           = 32
	SqlParserFIRST           = 33
	SqlParserLAST            = 34
	SqlParserHAVING          = 35
	SqlParserNOT             = 36
	SqlParserIS              = 37
	SqlParserBETWEEN         = 38
	SqlParserAND             = 39
	SqlParserIN              = 40
	SqlParserNULL            = 41
	SqlParserOR              = 42
	SqlParserASC             = 43
	SqlParserDESC            = 44
	SqlParserLIMIT           = 45
	SqlParserOFFSET          = 46
	SqlParserLIKE            = 47
	SqlParserEXISTS          = 48
	SqlParserCAST            = 49
	SqlParserINNER           = 50
	SqlParserLEFT            = 51
	SqlParserRIGHT           = 52
	SqlParserJOIN            = 53
	SqlParserON              = 54
	SqlParserUNION           = 55
	SqlParserALL             = 56
	SqlParserDECIMAL_LITERAL = 57
	SqlParserID              = 58
	SqlParserTEXT_STRING     = 59
	SqlParserTEXT_ALIAS      = 60
	SqlParserBIND_VARIABLE   = 61
	SqlParserCOMMENT         = 62
	SqlParserWS              = 63
)

// SqlParser rules.
const (
	SqlParserRULE_columnName         = 0
	SqlParserRULE_tableName          = 1
	SqlParserRULE_typeName           = 2
	SqlParserRULE_functionName       = 3
	SqlParserRULE_alias              = 4
	SqlParserRULE_join               = 5
	SqlParserRULE_union              = 6
	SqlParserRULE_decimalLiteral     = 7
	SqlParserRULE_textLiteral        = 8
	SqlParserRULE_bind_variables     = 9
	SqlParserRULE_selectStatement    = 10
	SqlParserRULE_selectElements     = 11
	SqlParserRULE_selectElement      = 12
	SqlParserRULE_expr               = 13
	SqlParserRULE_value              = 14
	SqlParserRULE_functionCall       = 15
	SqlParserRULE_aggregateFunction  = 16
	SqlParserRULE_commonFunction     = 17
	SqlParserRULE_functionArg        = 18
	SqlParserRULE_tableSources       = 19
	SqlParserRULE_tableSource        = 20
	SqlParserRULE_joinClause         = 21
	SqlParserRULE_whereClause        = 22
	SqlParserRULE_logicExpression    = 23
	SqlParserRULE_comparisonOperator = 24
	SqlParserRULE_groupByClause      = 25
	SqlParserRULE_groupByItem        = 26
	SqlParserRULE_havingClause       = 27
	SqlParserRULE_orderByClause      = 28
	SqlParserRULE_orderByExpression  = 29
	SqlParserRULE_limitClause        = 30
)

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnName() (localctx IColumnNameContext) {
	this := p
	_ = this

	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_columnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(SqlParserID)
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableName() (localctx ITableNameContext) {
	this := p
	_ = this

	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(SqlParserID)
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TypeName() (localctx ITypeNameContext) {
	this := p
	_ = this

	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(SqlParserID)
	}

	return localctx
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (s *FunctionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FunctionName() (localctx IFunctionNameContext) {
	this := p
	_ = this

	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_functionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(SqlParserID)
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TEXT_ALIAS() antlr.TerminalNode

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *AliasContext) TEXT_ALIAS() antlr.TerminalNode {
	return s.GetToken(SqlParserTEXT_ALIAS, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Alias() (localctx IAliasContext) {
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_alias)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserID || _la == SqlParserTEXT_ALIAS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJoinContext is an interface to support dynamic dispatch.
type IJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN() antlr.TerminalNode
	INNER() antlr.TerminalNode
	LEFT() antlr.TerminalNode
	RIGHT() antlr.TerminalNode

	// IsJoinContext differentiates from other interfaces.
	IsJoinContext()
}

type JoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinContext() *JoinContext {
	var p = new(JoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_join
	return p
}

func (*JoinContext) IsJoinContext() {}

func NewJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinContext {
	var p = new(JoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_join

	return p
}

func (s *JoinContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(SqlParserJOIN, 0)
}

func (s *JoinContext) INNER() antlr.TerminalNode {
	return s.GetToken(SqlParserINNER, 0)
}

func (s *JoinContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SqlParserLEFT, 0)
}

func (s *JoinContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(SqlParserRIGHT, 0)
}

func (s *JoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterJoin(s)
	}
}

func (s *JoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitJoin(s)
	}
}

func (s *JoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Join() (localctx IJoinContext) {
	this := p
	_ = this

	localctx = NewJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_join)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(73)
		p.Match(SqlParserJOIN)
	}

	return localctx
}

// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNION() antlr.TerminalNode
	ALL() antlr.TerminalNode

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	var p = new(UnionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_union
	return p
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	var p = new(UnionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) UNION() antlr.TerminalNode {
	return s.GetToken(SqlParserUNION, 0)
}

func (s *UnionContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitUnion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Union() (localctx IUnionContext) {
	this := p
	_ = this

	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_union)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(SqlParserUNION)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserALL {
		{
			p.SetState(76)
			p.Match(SqlParserALL)
		}

	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_LITERAL() antlr.TerminalNode

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_decimalLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(SqlParserDECIMAL_LITERAL)
	}

	return localctx
}

// ITextLiteralContext is an interface to support dynamic dispatch.
type ITextLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_STRING() antlr.TerminalNode

	// IsTextLiteralContext differentiates from other interfaces.
	IsTextLiteralContext()
}

type TextLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextLiteralContext() *TextLiteralContext {
	var p = new(TextLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_textLiteral
	return p
}

func (*TextLiteralContext) IsTextLiteralContext() {}

func NewTextLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextLiteralContext {
	var p = new(TextLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_textLiteral

	return p
}

func (s *TextLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TextLiteralContext) TEXT_STRING() antlr.TerminalNode {
	return s.GetToken(SqlParserTEXT_STRING, 0)
}

func (s *TextLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTextLiteral(s)
	}
}

func (s *TextLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTextLiteral(s)
	}
}

func (s *TextLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTextLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TextLiteral() (localctx ITextLiteralContext) {
	this := p
	_ = this

	localctx = NewTextLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_textLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(SqlParserTEXT_STRING)
	}

	return localctx
}

// IBind_variablesContext is an interface to support dynamic dispatch.
type IBind_variablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BIND_VARIABLE() antlr.TerminalNode

	// IsBind_variablesContext differentiates from other interfaces.
	IsBind_variablesContext()
}

type Bind_variablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBind_variablesContext() *Bind_variablesContext {
	var p = new(Bind_variablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_bind_variables
	return p
}

func (*Bind_variablesContext) IsBind_variablesContext() {}

func NewBind_variablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bind_variablesContext {
	var p = new(Bind_variablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_bind_variables

	return p
}

func (s *Bind_variablesContext) GetParser() antlr.Parser { return s.parser }

func (s *Bind_variablesContext) BIND_VARIABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserBIND_VARIABLE, 0)
}

func (s *Bind_variablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bind_variablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bind_variablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterBind_variables(s)
	}
}

func (s *Bind_variablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitBind_variables(s)
	}
}

func (s *Bind_variablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitBind_variables(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Bind_variables() (localctx IBind_variablesContext) {
	this := p
	_ = this

	localctx = NewBind_variablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SqlParserRULE_bind_variables)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(SqlParserBIND_VARIABLE)
	}

	return localctx
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	SelectElements() ISelectElementsContext
	FROM() antlr.TerminalNode
	TableSources() ITableSourcesContext
	JoinClause() IJoinClauseContext
	WhereClause() IWhereClauseContext
	GroupByClause() IGroupByClauseContext
	HavingClause() IHavingClauseContext
	OrderByClause() IOrderByClauseContext
	LimitClause() ILimitClauseContext
	AllSelectStatement() []ISelectStatementContext
	SelectStatement(i int) ISelectStatementContext
	Union() IUnionContext

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *SelectStatementContext) SelectElements() ISelectElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectElementsContext)
}

func (s *SelectStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *SelectStatementContext) TableSources() ITableSourcesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableSourcesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableSourcesContext)
}

func (s *SelectStatementContext) JoinClause() IJoinClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinClauseContext)
}

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupByClause() IGroupByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByClauseContext)
}

func (s *SelectStatementContext) HavingClause() IHavingClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) AllSelectStatement() []ISelectStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectStatementContext); ok {
			len++
		}
	}

	tst := make([]ISelectStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectStatementContext); ok {
			tst[i] = t.(ISelectStatementContext)
			i++
		}
	}

	return tst
}

func (s *SelectStatementContext) SelectStatement(i int) ISelectStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *SelectStatementContext) Union() IUnionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (s *SelectStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectStatement() (localctx ISelectStatementContext) {
	return p.selectStatement(0)
}

func (p *SqlParser) selectStatement(_p int) (localctx ISelectStatementContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISelectStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SqlParserRULE_selectStatement, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(87)
		p.SelectElements()
	}
	{
		p.SetState(88)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(89)
		p.TableSources()
	}
	{
		p.SetState(90)
		p.JoinClause()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(91)
			p.WhereClause()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(94)
			p.GroupByClause()
		}

	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(97)
			p.HavingClause()
		}

	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(100)
			p.OrderByClause()
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(103)
			p.LimitClause()
		}

	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSelectStatementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_selectStatement)
			p.SetState(106)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(107)
				p.Union()
			}
			{
				p.SetState(108)
				p.selectStatement(2)
			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ISelectElementsContext is an interface to support dynamic dispatch.
type ISelectElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// Getter signatures
	AllSelectElement() []ISelectElementContext
	SelectElement(i int) ISelectElementContext

	// IsSelectElementsContext differentiates from other interfaces.
	IsSelectElementsContext()
}

type SelectElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptySelectElementsContext() *SelectElementsContext {
	var p = new(SelectElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElements

	return p
}

func (s *SelectElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementsContext) GetStar() antlr.Token { return s.star }

func (s *SelectElementsContext) SetStar(v antlr.Token) { s.star = v }

func (s *SelectElementsContext) AllSelectElement() []ISelectElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectElementContext); ok {
			len++
		}
	}

	tst := make([]ISelectElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectElementContext); ok {
			tst[i] = t.(ISelectElementContext)
			i++
		}
	}

	return tst
}

func (s *SelectElementsContext) SelectElement(i int) ISelectElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElements() (localctx ISelectElementsContext) {
	this := p
	_ = this

	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_selectElements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserT__0:
		{
			p.SetState(115)

			var _m = p.Match(SqlParserT__0)

			localctx.(*SelectElementsContext).star = _m
		}

	case SqlParserT__6, SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT, SqlParserDISTINCT, SqlParserCAST, SqlParserDECIMAL_LITERAL, SqlParserID, SqlParserTEXT_STRING, SqlParserBIND_VARIABLE:
		{
			p.SetState(116)
			p.SelectElement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserT__1 {
		{
			p.SetState(119)
			p.Match(SqlParserT__1)
		}
		{
			p.SetState(120)
			p.SelectElement()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Alias() IAliasContext
	AS() antlr.TerminalNode

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelectElementContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *SelectElementContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectElement(s)
	}
}

func (s *SelectElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectElement(s)
	}
}

func (s *SelectElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElement() (localctx ISelectElementContext) {
	this := p
	_ = this

	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_selectElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.expr(0)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1441151880759083008) != 0 {
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserAS {
			{
				p.SetState(127)
				p.Match(SqlParserAS)
			}

		}
		{
			p.SetState(130)
			p.Alias()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnName() IColumnNameContext
	FunctionCall() IFunctionCallContext
	Value() IValueContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExprContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SqlParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, SqlParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(134)
			p.ColumnName()
		}

	case 2:
		{
			p.SetState(135)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(136)
			p.Value()
		}

	case 4:
		{
			p.SetState(137)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(SqlParserT__7)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserT__0 || _la == SqlParserT__2) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(145)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(147)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserT__3 || _la == SqlParserT__4) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(148)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}

				{
					p.SetState(150)
					p.Match(SqlParserT__5)
				}

				{
					p.SetState(151)
					p.expr(3)
				}

			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DecimalLiteral() IDecimalLiteralContext
	TextLiteral() ITextLiteralContext
	Bind_variables() IBind_variablesContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DecimalLiteral() IDecimalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *ValueContext) TextLiteral() ITextLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLiteralContext)
}

func (s *ValueContext) Bind_variables() IBind_variablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBind_variablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBind_variablesContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.DecimalLiteral()
		}

	case SqlParserTEXT_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.TextLiteral()
		}

	case SqlParserBIND_VARIABLE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Bind_variables()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AggregateFunction() IAggregateFunctionContext
	CommonFunction() ICommonFunctionContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *FunctionCallContext) CommonFunction() ICommonFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommonFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommonFunctionContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_functionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.AggregateFunction()
		}

	case SqlParserDISTINCT, SqlParserCAST, SqlParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.CommonFunction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAggregateFunctionContext is an interface to support dynamic dispatch.
type IAggregateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStarArg returns the starArg token.
	GetStarArg() antlr.Token

	// SetStarArg sets the starArg token.
	SetStarArg(antlr.Token)

	// Getter signatures
	FunctionArg() IFunctionArgContext
	AVG() antlr.TerminalNode
	MAX() antlr.TerminalNode
	MIN() antlr.TerminalNode
	SUM() antlr.TerminalNode
	COUNT() antlr.TerminalNode

	// IsAggregateFunctionContext differentiates from other interfaces.
	IsAggregateFunctionContext()
}

type AggregateFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	starArg antlr.Token
}

func NewEmptyAggregateFunctionContext() *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_aggregateFunction
	return p
}

func (*AggregateFunctionContext) IsAggregateFunctionContext() {}

func NewAggregateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_aggregateFunction

	return p
}

func (s *AggregateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunctionContext) GetStarArg() antlr.Token { return s.starArg }

func (s *AggregateFunctionContext) SetStarArg(v antlr.Token) { s.starArg = v }

func (s *AggregateFunctionContext) FunctionArg() IFunctionArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *AggregateFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(SqlParserAVG, 0)
}

func (s *AggregateFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(SqlParserMAX, 0)
}

func (s *AggregateFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(SqlParserMIN, 0)
}

func (s *AggregateFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(SqlParserSUM, 0)
}

func (s *AggregateFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOUNT, 0)
}

func (s *AggregateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitAggregateFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) AggregateFunction() (localctx IAggregateFunctionContext) {
	this := p
	_ = this

	localctx = NewAggregateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_aggregateFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(167)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(168)
			p.FunctionArg()
		}
		{
			p.SetState(169)
			p.Match(SqlParserT__7)
		}

	case SqlParserCOUNT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(SqlParserCOUNT)
		}
		{
			p.SetState(172)
			p.Match(SqlParserT__6)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SqlParserT__0:
			{
				p.SetState(173)

				var _m = p.Match(SqlParserT__0)

				localctx.(*AggregateFunctionContext).starArg = _m
			}

		case SqlParserT__6, SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT, SqlParserDISTINCT, SqlParserCAST, SqlParserDECIMAL_LITERAL, SqlParserID, SqlParserTEXT_STRING, SqlParserBIND_VARIABLE:
			{
				p.SetState(174)
				p.FunctionArg()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(177)
			p.Match(SqlParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommonFunctionContext is an interface to support dynamic dispatch.
type ICommonFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionName() IFunctionNameContext
	FunctionArg() IFunctionArgContext
	DISTINCT() antlr.TerminalNode
	CAST() antlr.TerminalNode
	AS() antlr.TerminalNode
	TypeName() ITypeNameContext

	// IsCommonFunctionContext differentiates from other interfaces.
	IsCommonFunctionContext()
}

type CommonFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommonFunctionContext() *CommonFunctionContext {
	var p = new(CommonFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_commonFunction
	return p
}

func (*CommonFunctionContext) IsCommonFunctionContext() {}

func NewCommonFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommonFunctionContext {
	var p = new(CommonFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_commonFunction

	return p
}

func (s *CommonFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CommonFunctionContext) FunctionName() IFunctionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *CommonFunctionContext) FunctionArg() IFunctionArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *CommonFunctionContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *CommonFunctionContext) CAST() antlr.TerminalNode {
	return s.GetToken(SqlParserCAST, 0)
}

func (s *CommonFunctionContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *CommonFunctionContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CommonFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommonFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommonFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterCommonFunction(s)
	}
}

func (s *CommonFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitCommonFunction(s)
	}
}

func (s *CommonFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitCommonFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) CommonFunction() (localctx ICommonFunctionContext) {
	this := p
	_ = this

	localctx = NewCommonFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SqlParserRULE_commonFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.FunctionName()
		}
		{
			p.SetState(181)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(182)
			p.FunctionArg()
		}
		{
			p.SetState(183)
			p.Match(SqlParserT__7)
		}

	case SqlParserDISTINCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(SqlParserDISTINCT)
		}
		{
			p.SetState(186)
			p.FunctionArg()
		}

	case SqlParserCAST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(SqlParserCAST)
		}
		{
			p.SetState(188)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(189)
			p.FunctionArg()
		}
		{
			p.SetState(190)
			p.Match(SqlParserAS)
		}
		{
			p.SetState(191)
			p.TypeName()
		}
		{
			p.SetState(192)
			p.Match(SqlParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionArg
	return p
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionArg(s)
	}
}

func (s *FunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionArg(s)
	}
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FunctionArg() (localctx IFunctionArgContext) {
	this := p
	_ = this

	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_functionArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.expr(0)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(197)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(198)
				p.expr(0)
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// ITableSourcesContext is an interface to support dynamic dispatch.
type ITableSourcesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTableSource() []ITableSourceContext
	TableSource(i int) ITableSourceContext
	AllAlias() []IAliasContext
	Alias(i int) IAliasContext

	// IsTableSourcesContext differentiates from other interfaces.
	IsTableSourcesContext()
}

type TableSourcesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourcesContext() *TableSourcesContext {
	var p = new(TableSourcesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableSources
	return p
}

func (*TableSourcesContext) IsTableSourcesContext() {}

func NewTableSourcesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourcesContext {
	var p = new(TableSourcesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableSources

	return p
}

func (s *TableSourcesContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourcesContext) AllTableSource() []ITableSourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableSourceContext); ok {
			len++
		}
	}

	tst := make([]ITableSourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableSourceContext); ok {
			tst[i] = t.(ITableSourceContext)
			i++
		}
	}

	return tst
}

func (s *TableSourcesContext) TableSource(i int) ITableSourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableSourceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableSourceContext)
}

func (s *TableSourcesContext) AllAlias() []IAliasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAliasContext); ok {
			len++
		}
	}

	tst := make([]IAliasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAliasContext); ok {
			tst[i] = t.(IAliasContext)
			i++
		}
	}

	return tst
}

func (s *TableSourcesContext) Alias(i int) IAliasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *TableSourcesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourcesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableSourcesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableSources(s)
	}
}

func (s *TableSourcesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableSources(s)
	}
}

func (s *TableSourcesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableSources(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableSources() (localctx ITableSourcesContext) {
	this := p
	_ = this

	localctx = NewTableSourcesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_tableSources)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.TableSource()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(205)
			p.Alias()
		}

	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(208)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(209)
				p.TableSource()
			}
			p.SetState(211)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(210)
					p.Alias()
				}

			}

		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// ITableSourceContext is an interface to support dynamic dispatch.
type ITableSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TableName() ITableNameContext
	SelectStatement() ISelectStatementContext

	// IsTableSourceContext differentiates from other interfaces.
	IsTableSourceContext()
}

type TableSourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourceContext() *TableSourceContext {
	var p = new(TableSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableSource
	return p
}

func (*TableSourceContext) IsTableSourceContext() {}

func NewTableSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourceContext {
	var p = new(TableSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableSource

	return p
}

func (s *TableSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourceContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *TableSourceContext) SelectStatement() ISelectStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *TableSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableSource(s)
	}
}

func (s *TableSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableSource(s)
	}
}

func (s *TableSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableSource() (localctx ITableSourceContext) {
	this := p
	_ = this

	localctx = NewTableSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_tableSource)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.TableName()
		}

	case SqlParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(220)
			p.selectStatement(0)
		}
		{
			p.SetState(221)
			p.Match(SqlParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJoinClauseContext is an interface to support dynamic dispatch.
type IJoinClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJoin() []IJoinContext
	Join(i int) IJoinContext
	AllTableSources() []ITableSourcesContext
	TableSources(i int) ITableSourcesContext
	AllON() []antlr.TerminalNode
	ON(i int) antlr.TerminalNode
	AllLogicExpression() []ILogicExpressionContext
	LogicExpression(i int) ILogicExpressionContext

	// IsJoinClauseContext differentiates from other interfaces.
	IsJoinClauseContext()
}

type JoinClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinClauseContext() *JoinClauseContext {
	var p = new(JoinClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_joinClause
	return p
}

func (*JoinClauseContext) IsJoinClauseContext() {}

func NewJoinClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinClauseContext {
	var p = new(JoinClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_joinClause

	return p
}

func (s *JoinClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinClauseContext) AllJoin() []IJoinContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinContext); ok {
			len++
		}
	}

	tst := make([]IJoinContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinContext); ok {
			tst[i] = t.(IJoinContext)
			i++
		}
	}

	return tst
}

func (s *JoinClauseContext) Join(i int) IJoinContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinContext)
}

func (s *JoinClauseContext) AllTableSources() []ITableSourcesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableSourcesContext); ok {
			len++
		}
	}

	tst := make([]ITableSourcesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableSourcesContext); ok {
			tst[i] = t.(ITableSourcesContext)
			i++
		}
	}

	return tst
}

func (s *JoinClauseContext) TableSources(i int) ITableSourcesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableSourcesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableSourcesContext)
}

func (s *JoinClauseContext) AllON() []antlr.TerminalNode {
	return s.GetTokens(SqlParserON)
}

func (s *JoinClauseContext) ON(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserON, i)
}

func (s *JoinClauseContext) AllLogicExpression() []ILogicExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicExpressionContext); ok {
			tst[i] = t.(ILogicExpressionContext)
			i++
		}
	}

	return tst
}

func (s *JoinClauseContext) LogicExpression(i int) ILogicExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *JoinClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterJoinClause(s)
	}
}

func (s *JoinClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitJoinClause(s)
	}
}

func (s *JoinClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitJoinClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) JoinClause() (localctx IJoinClauseContext) {
	this := p
	_ = this

	localctx = NewJoinClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_joinClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(225)
				p.Join()
			}
			{
				p.SetState(226)
				p.TableSources()
			}
			{
				p.SetState(227)
				p.Match(SqlParserON)
			}
			{
				p.SetState(228)
				p.logicExpression(0)
			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	LogicExpression() ILogicExpressionContext

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
}

func (s *WhereClauseContext) LogicExpression() ILogicExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) WhereClause() (localctx IWhereClauseContext) {
	this := p
	_ = this

	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SqlParserRULE_whereClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(SqlParserWHERE)
	}
	{
		p.SetState(236)
		p.logicExpression(0)
	}

	return localctx
}

// ILogicExpressionContext is an interface to support dynamic dispatch.
type ILogicExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftBracket returns the leftBracket token.
	GetLeftBracket() antlr.Token

	// GetRightBracket returns the rightBracket token.
	GetRightBracket() antlr.Token

	// GetLogicalOperator returns the logicalOperator token.
	GetLogicalOperator() antlr.Token

	// SetLeftBracket sets the leftBracket token.
	SetLeftBracket(antlr.Token)

	// SetRightBracket sets the rightBracket token.
	SetRightBracket(antlr.Token)

	// SetLogicalOperator sets the logicalOperator token.
	SetLogicalOperator(antlr.Token)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ComparisonOperator() IComparisonOperatorContext
	BETWEEN() antlr.TerminalNode
	AND() antlr.TerminalNode
	NOT() antlr.TerminalNode
	IN() antlr.TerminalNode
	SelectStatement() ISelectStatementContext
	LIKE() antlr.TerminalNode
	IS() antlr.TerminalNode
	NULL() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	AllLogicExpression() []ILogicExpressionContext
	LogicExpression(i int) ILogicExpressionContext
	COMMENT() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsLogicExpressionContext differentiates from other interfaces.
	IsLogicExpressionContext()
}

type LogicExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	leftBracket     antlr.Token
	rightBracket    antlr.Token
	logicalOperator antlr.Token
}

func NewEmptyLogicExpressionContext() *LogicExpressionContext {
	var p = new(LogicExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_logicExpression
	return p
}

func (*LogicExpressionContext) IsLogicExpressionContext() {}

func NewLogicExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicExpressionContext {
	var p = new(LogicExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_logicExpression

	return p
}

func (s *LogicExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicExpressionContext) GetLeftBracket() antlr.Token { return s.leftBracket }

func (s *LogicExpressionContext) GetRightBracket() antlr.Token { return s.rightBracket }

func (s *LogicExpressionContext) GetLogicalOperator() antlr.Token { return s.logicalOperator }

func (s *LogicExpressionContext) SetLeftBracket(v antlr.Token) { s.leftBracket = v }

func (s *LogicExpressionContext) SetRightBracket(v antlr.Token) { s.rightBracket = v }

func (s *LogicExpressionContext) SetLogicalOperator(v antlr.Token) { s.logicalOperator = v }

func (s *LogicExpressionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicExpressionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *LogicExpressionContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SqlParserBETWEEN, 0)
}

func (s *LogicExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *LogicExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *LogicExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(SqlParserIN, 0)
}

func (s *LogicExpressionContext) SelectStatement() ISelectStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *LogicExpressionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserLIKE, 0)
}

func (s *LogicExpressionContext) IS() antlr.TerminalNode {
	return s.GetToken(SqlParserIS, 0)
}

func (s *LogicExpressionContext) NULL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL, 0)
}

func (s *LogicExpressionContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SqlParserEXISTS, 0)
}

func (s *LogicExpressionContext) AllLogicExpression() []ILogicExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicExpressionContext); ok {
			tst[i] = t.(ILogicExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicExpressionContext) LogicExpression(i int) ILogicExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *LogicExpressionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMENT, 0)
}

func (s *LogicExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *LogicExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLogicExpression(s)
	}
}

func (s *LogicExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLogicExpression(s)
	}
}

func (s *LogicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitLogicExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LogicExpression() (localctx ILogicExpressionContext) {
	return p.logicExpression(0)
}

func (p *SqlParser) logicExpression(_p int) (localctx ILogicExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SqlParserRULE_logicExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(239)
			p.expr(0)
		}
		{
			p.SetState(240)
			p.ComparisonOperator()
		}
		{
			p.SetState(241)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(243)
			p.expr(0)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(244)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(247)
			p.Match(SqlParserBETWEEN)
		}
		{
			p.SetState(248)
			p.expr(0)
		}
		{
			p.SetState(249)
			p.Match(SqlParserAND)
		}
		{
			p.SetState(250)
			p.expr(0)
		}

	case 3:
		{
			p.SetState(252)
			p.expr(0)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(253)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(256)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(257)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(258)
			p.expr(0)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(259)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(260)
				p.expr(0)
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(266)
			p.Match(SqlParserT__7)
		}

	case 4:
		{
			p.SetState(268)
			p.expr(0)
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(269)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(272)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(273)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(274)
			p.selectStatement(0)
		}
		{
			p.SetState(275)
			p.Match(SqlParserT__7)
		}

	case 5:
		{
			p.SetState(277)
			p.expr(0)
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(278)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(281)
			p.Match(SqlParserLIKE)
		}
		{
			p.SetState(282)
			p.expr(0)
		}

	case 6:
		{
			p.SetState(284)
			p.expr(0)
		}
		{
			p.SetState(285)
			p.Match(SqlParserIS)
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(286)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(289)
			p.Match(SqlParserNULL)
		}

	case 7:
		{
			p.SetState(291)
			p.Match(SqlParserEXISTS)
		}
		{
			p.SetState(292)

			var _m = p.Match(SqlParserT__6)

			localctx.(*LogicExpressionContext).leftBracket = _m
		}
		{
			p.SetState(293)
			p.selectStatement(0)
		}
		{
			p.SetState(294)

			var _m = p.Match(SqlParserT__7)

			localctx.(*LogicExpressionContext).rightBracket = _m
		}

	case 8:
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserCOMMENT {
			{
				p.SetState(296)
				p.Match(SqlParserCOMMENT)
			}

		}
		{
			p.SetState(299)

			var _m = p.Match(SqlParserT__6)

			localctx.(*LogicExpressionContext).leftBracket = _m
		}
		{
			p.SetState(300)
			p.logicExpression(0)
		}
		{
			p.SetState(301)

			var _m = p.Match(SqlParserT__7)

			localctx.(*LogicExpressionContext).rightBracket = _m
		}

	case 9:
		{
			p.SetState(303)
			p.Match(SqlParserNOT)
		}
		{
			p.SetState(304)
			p.logicExpression(3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_logicExpression)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(308)

					var _m = p.Match(SqlParserAND)

					localctx.(*LogicExpressionContext).logicalOperator = _m
				}
				{
					p.SetState(309)
					p.logicExpression(3)
				}

			case 2:
				localctx = NewLogicExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_logicExpression)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(311)

					var _m = p.Match(SqlParserOR)

					localctx.(*LogicExpressionContext).logicalOperator = _m
				}
				{
					p.SetState(312)
					p.logicExpression(2)
				}

			}

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SqlParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&523776) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGroupByClauseContext is an interface to support dynamic dispatch.
type IGroupByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllGroupByItem() []IGroupByItemContext
	GroupByItem(i int) IGroupByItemContext

	// IsGroupByClauseContext differentiates from other interfaces.
	IsGroupByClauseContext()
}

type GroupByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByClauseContext() *GroupByClauseContext {
	var p = new(GroupByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupByClause
	return p
}

func (*GroupByClauseContext) IsGroupByClauseContext() {}

func NewGroupByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByClauseContext {
	var p = new(GroupByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupByClause

	return p
}

func (s *GroupByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SqlParserGROUP, 0)
}

func (s *GroupByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *GroupByClauseContext) AllGroupByItem() []IGroupByItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupByItemContext); ok {
			len++
		}
	}

	tst := make([]IGroupByItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupByItemContext); ok {
			tst[i] = t.(IGroupByItemContext)
			i++
		}
	}

	return tst
}

func (s *GroupByClauseContext) GroupByItem(i int) IGroupByItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByItemContext)
}

func (s *GroupByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupByClause(s)
	}
}

func (s *GroupByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupByClause(s)
	}
}

func (s *GroupByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitGroupByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) GroupByClause() (localctx IGroupByClauseContext) {
	this := p
	_ = this

	localctx = NewGroupByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_groupByClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(SqlParserGROUP)
	}
	{
		p.SetState(321)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(322)
		p.GroupByItem()
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(323)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(324)
				p.GroupByItem()
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// IGroupByItemContext is an interface to support dynamic dispatch.
type IGroupByItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsGroupByItemContext differentiates from other interfaces.
	IsGroupByItemContext()
}

type GroupByItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByItemContext() *GroupByItemContext {
	var p = new(GroupByItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupByItem
	return p
}

func (*GroupByItemContext) IsGroupByItemContext() {}

func NewGroupByItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByItemContext {
	var p = new(GroupByItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupByItem

	return p
}

func (s *GroupByItemContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupByItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupByItem(s)
	}
}

func (s *GroupByItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupByItem(s)
	}
}

func (s *GroupByItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitGroupByItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) GroupByItem() (localctx IGroupByItemContext) {
	this := p
	_ = this

	localctx = NewGroupByItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SqlParserRULE_groupByItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.expr(0)
	}

	return localctx
}

// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HAVING() antlr.TerminalNode
	LogicExpression() ILogicExpressionContext

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_havingClause
	return p
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SqlParserHAVING, 0)
}

func (s *HavingClauseContext) LogicExpression() ILogicExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) HavingClause() (localctx IHavingClauseContext) {
	this := p
	_ = this

	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SqlParserRULE_havingClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(SqlParserHAVING)
	}
	{
		p.SetState(333)
		p.logicExpression(0)
	}

	return localctx
}

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllOrderByExpression() []IOrderByExpressionContext
	OrderByExpression(i int) IOrderByExpressionContext

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderByExpression() []IOrderByExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderByExpressionContext); ok {
			len++
		}
	}

	tst := make([]IOrderByExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderByExpressionContext); ok {
			tst[i] = t.(IOrderByExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByExpression(i int) IOrderByExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByExpressionContext)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (s *OrderByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitOrderByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByClause() (localctx IOrderByClauseContext) {
	this := p
	_ = this

	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SqlParserRULE_orderByClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(SqlParserORDER)
	}
	{
		p.SetState(336)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(337)
		p.OrderByExpression()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(338)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(339)
				p.OrderByExpression()
			}

		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// IOrderByExpressionContext is an interface to support dynamic dispatch.
type IOrderByExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrder returns the order token.
	GetOrder() antlr.Token

	// SetOrder sets the order token.
	SetOrder(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	ASC() antlr.TerminalNode
	DESC() antlr.TerminalNode
	NULLS() antlr.TerminalNode
	FIRST() antlr.TerminalNode
	LAST() antlr.TerminalNode

	// IsOrderByExpressionContext differentiates from other interfaces.
	IsOrderByExpressionContext()
}

type OrderByExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	order  antlr.Token
}

func NewEmptyOrderByExpressionContext() *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByExpression
	return p
}

func (*OrderByExpressionContext) IsOrderByExpressionContext() {}

func NewOrderByExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByExpression

	return p
}

func (s *OrderByExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByExpressionContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByExpressionContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByExpressionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrderByExpressionContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *OrderByExpressionContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *OrderByExpressionContext) NULLS() antlr.TerminalNode {
	return s.GetToken(SqlParserNULLS, 0)
}

func (s *OrderByExpressionContext) FIRST() antlr.TerminalNode {
	return s.GetToken(SqlParserFIRST, 0)
}

func (s *OrderByExpressionContext) LAST() antlr.TerminalNode {
	return s.GetToken(SqlParserLAST, 0)
}

func (s *OrderByExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitOrderByExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByExpression() (localctx IOrderByExpressionContext) {
	this := p
	_ = this

	localctx = NewOrderByExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SqlParserRULE_orderByExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.expr(0)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(346)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByExpressionContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserASC || _la == SqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderByExpressionContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(349)
			p.Match(SqlParserNULLS)
		}
		{
			p.SetState(350)
			p.Match(SqlParserFIRST)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(351)
			p.Match(SqlParserNULLS)
		}
		{
			p.SetState(352)
			p.Match(SqlParserLAST)
		}

	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset rule contexts.
	GetOffset() IDecimalLiteralContext

	// GetLimit returns the limit rule contexts.
	GetLimit() IDecimalLiteralContext

	// SetOffset sets the offset rule contexts.
	SetOffset(IDecimalLiteralContext)

	// SetLimit sets the limit rule contexts.
	SetLimit(IDecimalLiteralContext)

	// Getter signatures
	LIMIT() antlr.TerminalNode
	OFFSET() antlr.TerminalNode
	AllDecimalLiteral() []IDecimalLiteralContext
	DecimalLiteral(i int) IDecimalLiteralContext

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	offset IDecimalLiteralContext
	limit  IDecimalLiteralContext
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() IDecimalLiteralContext { return s.offset }

func (s *LimitClauseContext) GetLimit() IDecimalLiteralContext { return s.limit }

func (s *LimitClauseContext) SetOffset(v IDecimalLiteralContext) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v IDecimalLiteralContext) { s.limit = v }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *LimitClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SqlParserOFFSET, 0)
}

func (s *LimitClauseContext) AllDecimalLiteral() []IDecimalLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			len++
		}
	}

	tst := make([]IDecimalLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDecimalLiteralContext); ok {
			tst[i] = t.(IDecimalLiteralContext)
			i++
		}
	}

	return tst
}

func (s *LimitClauseContext) DecimalLiteral(i int) IDecimalLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClause() (localctx ILimitClauseContext) {
	this := p
	_ = this

	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SqlParserRULE_limitClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(SqlParserLIMIT)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.SetState(359)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(356)

				var _x = p.DecimalLiteral()

				localctx.(*LimitClauseContext).offset = _x
			}
			{
				p.SetState(357)
				p.Match(SqlParserT__1)
			}

		}
		{
			p.SetState(361)

			var _x = p.DecimalLiteral()

			localctx.(*LimitClauseContext).limit = _x
		}

	case 2:
		{
			p.SetState(362)

			var _x = p.DecimalLiteral()

			localctx.(*LimitClauseContext).limit = _x
		}
		{
			p.SetState(363)
			p.Match(SqlParserOFFSET)
		}
		{
			p.SetState(364)

			var _x = p.DecimalLiteral()

			localctx.(*LimitClauseContext).offset = _x
		}

	}

	return localctx
}

func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *SelectStatementContext = nil
		if localctx != nil {
			t = localctx.(*SelectStatementContext)
		}
		return p.SelectStatement_Sempred(t, predIndex)

	case 13:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 23:
		var t *LogicExpressionContext = nil
		if localctx != nil {
			t = localctx.(*LogicExpressionContext)
		}
		return p.LogicExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SqlParser) SelectStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) LogicExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
