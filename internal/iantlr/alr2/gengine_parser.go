// Code generated from gengine.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gengine

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type gengineParser struct {
	*antlr.BaseParser
}

var GengineParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gengineParserInit() {
	staticData := &GengineParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'@name'", "'@id'", "'@code'", "'@desc'", "'@sal'", "", "",
		"'&&'", "'||'", "'and'", "'or'", "'&'", "'|'", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'",
		"'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "'eq'", "'gt'",
		"'lt'", "'ge'", "'le'", "'ne'", "':='", "'='", "'+='", "'-='", "'*='",
		"'/='", "'['", "']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "AND_STR", "OR_STR",
		"AND_SINGLE", "OR_SINGLE", "CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK",
		"FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE",
		"BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL",
		"EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "EQUALS_STR",
		"GT_STR", "LT_STR", "GTE_STR", "LTE_STR", "NOTEQUALS_STR", "ASSIGN",
		"SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE",
		"RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET",
		"DOT", "DQUOTA_STRING", "DOTTEDNAME", "DOUBLEDOTTEDNAME", "REAL_LITERAL",
		"SL_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"primary", "ruleEntity", "ruleName", "ruleDescription", "salience",
		"ruleContent", "statements", "statement", "concStatement", "expression",
		"mathExpression", "expressionAtom", "assignment", "returnStmt", "ifStmt",
		"elseIfStmt", "elseStmt", "forStmt", "breakStmt", "forRangeStmt", "continueStmt",
		"constant", "functionArgs", "integer", "realLiteral", "stringLiteral",
		"booleanLiteral", "functionCall", "methodCall", "threeLevelCall", "variable",
		"mathPmOperator", "mathMdOperator", "comparisonOperator", "logicalOperator",
		"assignOperator", "rangeOperator", "notOperator", "mapVar", "atName",
		"atId", "atCode", "atDesc", "atSal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 371, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 4, 0, 90, 8, 0, 11, 0, 12, 0, 91, 1, 1,
		1, 1, 1, 1, 3, 1, 97, 8, 1, 1, 1, 3, 1, 100, 8, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 5, 6, 116,
		8, 6, 10, 6, 12, 6, 119, 9, 6, 1, 6, 3, 6, 122, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 134, 8, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 142, 8, 8, 10, 8, 12, 8, 145, 9, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 152, 8, 9, 1, 9, 1, 9, 3, 9, 156, 8, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 3, 9, 162, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 172, 8, 9, 10, 9, 12, 9, 175, 9, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 183, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 193, 8, 10, 10, 10, 12, 10, 196,
		9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 204, 8, 11, 1,
		12, 1, 12, 3, 12, 208, 8, 12, 1, 12, 1, 12, 1, 12, 3, 12, 213, 8, 12, 1,
		13, 1, 13, 3, 13, 217, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		5, 14, 225, 8, 14, 10, 14, 12, 14, 228, 9, 14, 1, 14, 3, 14, 231, 8, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 275, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3,
		22, 284, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 294, 8, 22, 5, 22, 296, 8, 22, 10, 22, 12, 22, 299, 9, 22, 1, 23,
		3, 23, 302, 8, 23, 1, 23, 1, 23, 1, 24, 3, 24, 307, 8, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 318, 8, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 325, 8, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 3, 29, 332, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 357, 8, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 0, 2, 18, 20, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 0, 7, 1, 0,
		23, 24, 2, 0, 29, 29, 63, 64, 1, 0, 31, 32, 1, 0, 33, 34, 2, 0, 35, 40,
		42, 47, 1, 0, 9, 14, 1, 0, 48, 53, 390, 0, 89, 1, 0, 0, 0, 2, 93, 1, 0,
		0, 0, 4, 105, 1, 0, 0, 0, 6, 107, 1, 0, 0, 0, 8, 109, 1, 0, 0, 0, 10, 112,
		1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0, 16, 135, 1, 0, 0,
		0, 18, 161, 1, 0, 0, 0, 20, 182, 1, 0, 0, 0, 22, 203, 1, 0, 0, 0, 24, 207,
		1, 0, 0, 0, 26, 214, 1, 0, 0, 0, 28, 218, 1, 0, 0, 0, 30, 232, 1, 0, 0,
		0, 32, 239, 1, 0, 0, 0, 34, 244, 1, 0, 0, 0, 36, 254, 1, 0, 0, 0, 38, 256,
		1, 0, 0, 0, 40, 264, 1, 0, 0, 0, 42, 274, 1, 0, 0, 0, 44, 283, 1, 0, 0,
		0, 46, 301, 1, 0, 0, 0, 48, 306, 1, 0, 0, 0, 50, 310, 1, 0, 0, 0, 52, 312,
		1, 0, 0, 0, 54, 314, 1, 0, 0, 0, 56, 321, 1, 0, 0, 0, 58, 328, 1, 0, 0,
		0, 60, 335, 1, 0, 0, 0, 62, 337, 1, 0, 0, 0, 64, 339, 1, 0, 0, 0, 66, 341,
		1, 0, 0, 0, 68, 343, 1, 0, 0, 0, 70, 345, 1, 0, 0, 0, 72, 347, 1, 0, 0,
		0, 74, 349, 1, 0, 0, 0, 76, 351, 1, 0, 0, 0, 78, 360, 1, 0, 0, 0, 80, 362,
		1, 0, 0, 0, 82, 364, 1, 0, 0, 0, 84, 366, 1, 0, 0, 0, 86, 368, 1, 0, 0,
		0, 88, 90, 3, 2, 1, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 1, 1, 0, 0, 0, 93, 94, 5, 8, 0, 0,
		94, 96, 3, 4, 2, 0, 95, 97, 3, 6, 3, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 100, 3, 8, 4, 0, 99, 98, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 5, 27, 0, 0, 102, 103,
		3, 10, 5, 0, 103, 104, 5, 28, 0, 0, 104, 3, 1, 0, 0, 0, 105, 106, 3, 50,
		25, 0, 106, 5, 1, 0, 0, 0, 107, 108, 3, 50, 25, 0, 108, 7, 1, 0, 0, 0,
		109, 110, 5, 26, 0, 0, 110, 111, 3, 46, 23, 0, 111, 9, 1, 0, 0, 0, 112,
		113, 3, 12, 6, 0, 113, 11, 1, 0, 0, 0, 114, 116, 3, 14, 7, 0, 115, 114,
		1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0,
		0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122, 3, 26, 13,
		0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 13, 1, 0, 0, 0, 123,
		134, 3, 28, 14, 0, 124, 134, 3, 54, 27, 0, 125, 134, 3, 56, 28, 0, 126,
		134, 3, 58, 29, 0, 127, 134, 3, 24, 12, 0, 128, 134, 3, 16, 8, 0, 129,
		134, 3, 34, 17, 0, 130, 134, 3, 36, 18, 0, 131, 134, 3, 38, 19, 0, 132,
		134, 3, 40, 20, 0, 133, 123, 1, 0, 0, 0, 133, 124, 1, 0, 0, 0, 133, 125,
		1, 0, 0, 0, 133, 126, 1, 0, 0, 0, 133, 127, 1, 0, 0, 0, 133, 128, 1, 0,
		0, 0, 133, 129, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0,
		133, 132, 1, 0, 0, 0, 134, 15, 1, 0, 0, 0, 135, 136, 5, 15, 0, 0, 136,
		143, 5, 57, 0, 0, 137, 142, 3, 54, 27, 0, 138, 142, 3, 56, 28, 0, 139,
		142, 3, 58, 29, 0, 140, 142, 3, 24, 12, 0, 141, 137, 1, 0, 0, 0, 141, 138,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 146, 147, 5, 58, 0, 0, 147, 17, 1, 0, 0, 0, 148,
		149, 6, 9, -1, 0, 149, 162, 3, 20, 10, 0, 150, 152, 3, 74, 37, 0, 151,
		150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 162,
		3, 22, 11, 0, 154, 156, 3, 74, 37, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 59, 0, 0, 158, 159, 3, 18,
		9, 0, 159, 160, 5, 60, 0, 0, 160, 162, 1, 0, 0, 0, 161, 148, 1, 0, 0, 0,
		161, 151, 1, 0, 0, 0, 161, 155, 1, 0, 0, 0, 162, 173, 1, 0, 0, 0, 163,
		164, 10, 4, 0, 0, 164, 165, 3, 66, 33, 0, 165, 166, 3, 18, 9, 5, 166, 172,
		1, 0, 0, 0, 167, 168, 10, 3, 0, 0, 168, 169, 3, 68, 34, 0, 169, 170, 3,
		18, 9, 4, 170, 172, 1, 0, 0, 0, 171, 163, 1, 0, 0, 0, 171, 167, 1, 0, 0,
		0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		19, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 6, 10, -1, 0, 177, 183,
		3, 22, 11, 0, 178, 179, 5, 59, 0, 0, 179, 180, 3, 20, 10, 0, 180, 181,
		5, 60, 0, 0, 181, 183, 1, 0, 0, 0, 182, 176, 1, 0, 0, 0, 182, 178, 1, 0,
		0, 0, 183, 194, 1, 0, 0, 0, 184, 185, 10, 4, 0, 0, 185, 186, 3, 64, 32,
		0, 186, 187, 3, 20, 10, 5, 187, 193, 1, 0, 0, 0, 188, 189, 10, 3, 0, 0,
		189, 190, 3, 62, 31, 0, 190, 191, 3, 20, 10, 4, 191, 193, 1, 0, 0, 0, 192,
		184, 1, 0, 0, 0, 192, 188, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192,
		1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 21, 1, 0, 0, 0, 196, 194, 1, 0,
		0, 0, 197, 204, 3, 54, 27, 0, 198, 204, 3, 56, 28, 0, 199, 204, 3, 58,
		29, 0, 200, 204, 3, 42, 21, 0, 201, 204, 3, 76, 38, 0, 202, 204, 3, 60,
		30, 0, 203, 197, 1, 0, 0, 0, 203, 198, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0,
		203, 200, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204,
		23, 1, 0, 0, 0, 205, 208, 3, 76, 38, 0, 206, 208, 3, 60, 30, 0, 207, 205,
		1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 212, 3, 70,
		35, 0, 210, 213, 3, 20, 10, 0, 211, 213, 3, 18, 9, 0, 212, 210, 1, 0, 0,
		0, 212, 211, 1, 0, 0, 0, 213, 25, 1, 0, 0, 0, 214, 216, 5, 18, 0, 0, 215,
		217, 3, 18, 9, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 27,
		1, 0, 0, 0, 218, 219, 5, 16, 0, 0, 219, 220, 3, 18, 9, 0, 220, 221, 5,
		57, 0, 0, 221, 222, 3, 12, 6, 0, 222, 226, 5, 58, 0, 0, 223, 225, 3, 30,
		15, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0,
		226, 227, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229,
		231, 3, 32, 16, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 29,
		1, 0, 0, 0, 232, 233, 5, 17, 0, 0, 233, 234, 5, 16, 0, 0, 234, 235, 3,
		18, 9, 0, 235, 236, 5, 57, 0, 0, 236, 237, 3, 12, 6, 0, 237, 238, 5, 58,
		0, 0, 238, 31, 1, 0, 0, 0, 239, 240, 5, 17, 0, 0, 240, 241, 5, 57, 0, 0,
		241, 242, 3, 12, 6, 0, 242, 243, 5, 58, 0, 0, 243, 33, 1, 0, 0, 0, 244,
		245, 5, 19, 0, 0, 245, 246, 3, 24, 12, 0, 246, 247, 5, 56, 0, 0, 247, 248,
		3, 18, 9, 0, 248, 249, 5, 56, 0, 0, 249, 250, 3, 24, 12, 0, 250, 251, 5,
		57, 0, 0, 251, 252, 3, 12, 6, 0, 252, 253, 5, 58, 0, 0, 253, 35, 1, 0,
		0, 0, 254, 255, 5, 20, 0, 0, 255, 37, 1, 0, 0, 0, 256, 257, 5, 21, 0, 0,
		257, 258, 3, 60, 30, 0, 258, 259, 3, 72, 36, 0, 259, 260, 3, 60, 30, 0,
		260, 261, 5, 57, 0, 0, 261, 262, 3, 12, 6, 0, 262, 263, 5, 58, 0, 0, 263,
		39, 1, 0, 0, 0, 264, 265, 5, 22, 0, 0, 265, 41, 1, 0, 0, 0, 266, 275, 3,
		52, 26, 0, 267, 275, 3, 48, 24, 0, 268, 275, 3, 50, 25, 0, 269, 275, 3,
		78, 39, 0, 270, 275, 3, 80, 40, 0, 271, 275, 3, 82, 41, 0, 272, 275, 3,
		84, 42, 0, 273, 275, 3, 86, 43, 0, 274, 266, 1, 0, 0, 0, 274, 267, 1, 0,
		0, 0, 274, 268, 1, 0, 0, 0, 274, 269, 1, 0, 0, 0, 274, 270, 1, 0, 0, 0,
		274, 271, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 273, 1, 0, 0, 0, 275,
		43, 1, 0, 0, 0, 276, 284, 3, 42, 21, 0, 277, 284, 3, 60, 30, 0, 278, 284,
		3, 54, 27, 0, 279, 284, 3, 56, 28, 0, 280, 284, 3, 58, 29, 0, 281, 284,
		3, 76, 38, 0, 282, 284, 3, 18, 9, 0, 283, 276, 1, 0, 0, 0, 283, 277, 1,
		0, 0, 0, 283, 278, 1, 0, 0, 0, 283, 279, 1, 0, 0, 0, 283, 280, 1, 0, 0,
		0, 283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0, 284, 297, 1, 0, 0, 0, 285,
		293, 5, 1, 0, 0, 286, 294, 3, 42, 21, 0, 287, 294, 3, 60, 30, 0, 288, 294,
		3, 54, 27, 0, 289, 294, 3, 56, 28, 0, 290, 294, 3, 58, 29, 0, 291, 294,
		3, 76, 38, 0, 292, 294, 3, 18, 9, 0, 293, 286, 1, 0, 0, 0, 293, 287, 1,
		0, 0, 0, 293, 288, 1, 0, 0, 0, 293, 289, 1, 0, 0, 0, 293, 290, 1, 0, 0,
		0, 293, 291, 1, 0, 0, 0, 293, 292, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295,
		285, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298,
		1, 0, 0, 0, 298, 45, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 302, 5, 32,
		0, 0, 301, 300, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0,
		303, 304, 5, 30, 0, 0, 304, 47, 1, 0, 0, 0, 305, 307, 5, 32, 0, 0, 306,
		305, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309,
		5, 65, 0, 0, 309, 49, 1, 0, 0, 0, 310, 311, 5, 62, 0, 0, 311, 51, 1, 0,
		0, 0, 312, 313, 7, 0, 0, 0, 313, 53, 1, 0, 0, 0, 314, 315, 5, 29, 0, 0,
		315, 317, 5, 59, 0, 0, 316, 318, 3, 44, 22, 0, 317, 316, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 5, 60, 0, 0, 320, 55,
		1, 0, 0, 0, 321, 322, 5, 63, 0, 0, 322, 324, 5, 59, 0, 0, 323, 325, 3,
		44, 22, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0,
		0, 0, 326, 327, 5, 60, 0, 0, 327, 57, 1, 0, 0, 0, 328, 329, 5, 64, 0, 0,
		329, 331, 5, 59, 0, 0, 330, 332, 3, 44, 22, 0, 331, 330, 1, 0, 0, 0, 331,
		332, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 5, 60, 0, 0, 334, 59,
		1, 0, 0, 0, 335, 336, 7, 1, 0, 0, 336, 61, 1, 0, 0, 0, 337, 338, 7, 2,
		0, 0, 338, 63, 1, 0, 0, 0, 339, 340, 7, 3, 0, 0, 340, 65, 1, 0, 0, 0, 341,
		342, 7, 4, 0, 0, 342, 67, 1, 0, 0, 0, 343, 344, 7, 5, 0, 0, 344, 69, 1,
		0, 0, 0, 345, 346, 7, 6, 0, 0, 346, 71, 1, 0, 0, 0, 347, 348, 5, 48, 0,
		0, 348, 73, 1, 0, 0, 0, 349, 350, 5, 41, 0, 0, 350, 75, 1, 0, 0, 0, 351,
		352, 3, 60, 30, 0, 352, 356, 5, 54, 0, 0, 353, 357, 3, 46, 23, 0, 354,
		357, 3, 50, 25, 0, 355, 357, 3, 60, 30, 0, 356, 353, 1, 0, 0, 0, 356, 354,
		1, 0, 0, 0, 356, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 5, 55,
		0, 0, 359, 77, 1, 0, 0, 0, 360, 361, 5, 2, 0, 0, 361, 79, 1, 0, 0, 0, 362,
		363, 5, 3, 0, 0, 363, 81, 1, 0, 0, 0, 364, 365, 5, 4, 0, 0, 365, 83, 1,
		0, 0, 0, 366, 367, 5, 5, 0, 0, 367, 85, 1, 0, 0, 0, 368, 369, 5, 6, 0,
		0, 369, 87, 1, 0, 0, 0, 32, 91, 96, 99, 117, 121, 133, 141, 143, 151, 155,
		161, 171, 173, 182, 192, 194, 203, 207, 212, 216, 226, 230, 274, 283, 293,
		297, 301, 306, 317, 324, 331, 356,
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

// gengineParserInit initializes any static state used to implement gengineParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgengineParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GengineParserInit() {
	staticData := &GengineParserStaticData
	staticData.once.Do(gengineParserInit)
}

// NewgengineParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgengineParser(input antlr.TokenStream) *gengineParser {
	GengineParserInit()
	this := new(gengineParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GengineParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gengine.g4"

	return this
}

// gengineParser tokens.
const (
	gengineParserEOF              = antlr.TokenEOF
	gengineParserT__0             = 1
	gengineParserT__1             = 2
	gengineParserT__2             = 3
	gengineParserT__3             = 4
	gengineParserT__4             = 5
	gengineParserT__5             = 6
	gengineParserNIL              = 7
	gengineParserRULE             = 8
	gengineParserAND              = 9
	gengineParserOR               = 10
	gengineParserAND_STR          = 11
	gengineParserOR_STR           = 12
	gengineParserAND_SINGLE       = 13
	gengineParserOR_SINGLE        = 14
	gengineParserCONC             = 15
	gengineParserIF               = 16
	gengineParserELSE             = 17
	gengineParserRETURN           = 18
	gengineParserFOR              = 19
	gengineParserBREAK            = 20
	gengineParserFORRANGE         = 21
	gengineParserCONTINUE         = 22
	gengineParserTRUE             = 23
	gengineParserFALSE            = 24
	gengineParserNULL_LITERAL     = 25
	gengineParserSALIENCE         = 26
	gengineParserBEGIN            = 27
	gengineParserEND              = 28
	gengineParserSIMPLENAME       = 29
	gengineParserINT              = 30
	gengineParserPLUS             = 31
	gengineParserMINUS            = 32
	gengineParserDIV              = 33
	gengineParserMUL              = 34
	gengineParserEQUALS           = 35
	gengineParserGT               = 36
	gengineParserLT               = 37
	gengineParserGTE              = 38
	gengineParserLTE              = 39
	gengineParserNOTEQUALS        = 40
	gengineParserNOT              = 41
	gengineParserEQUALS_STR       = 42
	gengineParserGT_STR           = 43
	gengineParserLT_STR           = 44
	gengineParserGTE_STR          = 45
	gengineParserLTE_STR          = 46
	gengineParserNOTEQUALS_STR    = 47
	gengineParserASSIGN           = 48
	gengineParserSET              = 49
	gengineParserPLUSEQUAL        = 50
	gengineParserMINUSEQUAL       = 51
	gengineParserMULTIEQUAL       = 52
	gengineParserDIVEQUAL         = 53
	gengineParserLSQARE           = 54
	gengineParserRSQARE           = 55
	gengineParserSEMICOLON        = 56
	gengineParserLR_BRACE         = 57
	gengineParserRR_BRACE         = 58
	gengineParserLR_BRACKET       = 59
	gengineParserRR_BRACKET       = 60
	gengineParserDOT              = 61
	gengineParserDQUOTA_STRING    = 62
	gengineParserDOTTEDNAME       = 63
	gengineParserDOUBLEDOTTEDNAME = 64
	gengineParserREAL_LITERAL     = 65
	gengineParserSL_COMMENT       = 66
	gengineParserWS               = 67
)

// gengineParser rules.
const (
	gengineParserRULE_primary            = 0
	gengineParserRULE_ruleEntity         = 1
	gengineParserRULE_ruleName           = 2
	gengineParserRULE_ruleDescription    = 3
	gengineParserRULE_salience           = 4
	gengineParserRULE_ruleContent        = 5
	gengineParserRULE_statements         = 6
	gengineParserRULE_statement          = 7
	gengineParserRULE_concStatement      = 8
	gengineParserRULE_expression         = 9
	gengineParserRULE_mathExpression     = 10
	gengineParserRULE_expressionAtom     = 11
	gengineParserRULE_assignment         = 12
	gengineParserRULE_returnStmt         = 13
	gengineParserRULE_ifStmt             = 14
	gengineParserRULE_elseIfStmt         = 15
	gengineParserRULE_elseStmt           = 16
	gengineParserRULE_forStmt            = 17
	gengineParserRULE_breakStmt          = 18
	gengineParserRULE_forRangeStmt       = 19
	gengineParserRULE_continueStmt       = 20
	gengineParserRULE_constant           = 21
	gengineParserRULE_functionArgs       = 22
	gengineParserRULE_integer            = 23
	gengineParserRULE_realLiteral        = 24
	gengineParserRULE_stringLiteral      = 25
	gengineParserRULE_booleanLiteral     = 26
	gengineParserRULE_functionCall       = 27
	gengineParserRULE_methodCall         = 28
	gengineParserRULE_threeLevelCall     = 29
	gengineParserRULE_variable           = 30
	gengineParserRULE_mathPmOperator     = 31
	gengineParserRULE_mathMdOperator     = 32
	gengineParserRULE_comparisonOperator = 33
	gengineParserRULE_logicalOperator    = 34
	gengineParserRULE_assignOperator     = 35
	gengineParserRULE_rangeOperator      = 36
	gengineParserRULE_notOperator        = 37
	gengineParserRULE_mapVar             = 38
	gengineParserRULE_atName             = 39
	gengineParserRULE_atId               = 40
	gengineParserRULE_atCode             = 41
	gengineParserRULE_atDesc             = 42
	gengineParserRULE_atSal              = 43
)

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleEntity() []IRuleEntityContext
	RuleEntity(i int) IRuleEntityContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllRuleEntity() []IRuleEntityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleEntityContext); ok {
			len++
		}
	}

	tst := make([]IRuleEntityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleEntityContext); ok {
			tst[i] = t.(IRuleEntityContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) RuleEntity(i int) IRuleEntityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleEntityContext); ok {
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

	return t.(IRuleEntityContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gengineParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gengineParserRULE {
		{
			p.SetState(88)
			p.RuleEntity()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleEntityContext is an interface to support dynamic dispatch.
type IRuleEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE() antlr.TerminalNode
	RuleName() IRuleNameContext
	BEGIN() antlr.TerminalNode
	RuleContent() IRuleContentContext
	END() antlr.TerminalNode
	RuleDescription() IRuleDescriptionContext
	Salience() ISalienceContext

	// IsRuleEntityContext differentiates from other interfaces.
	IsRuleEntityContext()
}

type RuleEntityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntityContext() *RuleEntityContext {
	var p = new(RuleEntityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleEntity
	return p
}

func InitEmptyRuleEntityContext(p *RuleEntityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleEntity
}

func (*RuleEntityContext) IsRuleEntityContext() {}

func NewRuleEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntityContext {
	var p = new(RuleEntityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleEntity

	return p
}

func (s *RuleEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntityContext) RULE() antlr.TerminalNode {
	return s.GetToken(gengineParserRULE, 0)
}

func (s *RuleEntityContext) RuleName() IRuleNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntityContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(gengineParserBEGIN, 0)
}

func (s *RuleEntityContext) RuleContent() IRuleContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleContentContext)
}

func (s *RuleEntityContext) END() antlr.TerminalNode {
	return s.GetToken(gengineParserEND, 0)
}

func (s *RuleEntityContext) RuleDescription() IRuleDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntityContext) Salience() ISalienceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISalienceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleEntity(s)
	}
}

func (s *RuleEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleEntity(s)
	}
}

func (s *RuleEntityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleEntity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleEntity() (localctx IRuleEntityContext) {
	localctx = NewRuleEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gengineParserRULE_ruleEntity)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(gengineParserRULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.RuleName()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserDQUOTA_STRING {
		{
			p.SetState(95)
			p.RuleDescription()
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserSALIENCE {
		{
			p.SetState(98)
			p.Salience()
		}

	}
	{
		p.SetState(101)
		p.Match(gengineParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.RuleContent()
	}
	{
		p.SetState(103)
		p.Match(gengineParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() IStringLiteralContext

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleName
	return p
}

func InitEmptyRuleNameContext(p *RuleNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleName
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gengineParserRULE_ruleName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.StringLiteral()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() IStringLiteralContext

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleDescription
	return p
}

func InitEmptyRuleDescriptionContext(p *RuleDescriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleDescription
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gengineParserRULE_ruleDescription)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.StringLiteral()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SALIENCE() antlr.TerminalNode
	Integer() IIntegerContext

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_salience
	return p
}

func InitEmptySalienceContext(p *SalienceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_salience
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(gengineParserSALIENCE, 0)
}

func (s *SalienceContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitSalience(s)
	}
}

func (s *SalienceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitSalience(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gengineParserRULE_salience)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(gengineParserSALIENCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Integer()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleContentContext is an interface to support dynamic dispatch.
type IRuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statements() IStatementsContext

	// IsRuleContentContext differentiates from other interfaces.
	IsRuleContentContext()
}

type RuleContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContentContext() *RuleContentContext {
	var p = new(RuleContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleContent
	return p
}

func InitEmptyRuleContentContext(p *RuleContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ruleContent
}

func (*RuleContentContext) IsRuleContentContext() {}

func NewRuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContentContext {
	var p = new(RuleContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleContent

	return p
}

func (s *RuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContentContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *RuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleContent(s)
	}
}

func (s *RuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleContent(s)
	}
}

func (s *RuleContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleContent() (localctx IRuleContentContext) {
	localctx = NewRuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gengineParserRULE_ruleContent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Statements()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	ReturnStmt() IReturnStmtContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *StatementsContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gengineParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&844424930148595) != 0 {
		{
			p.SetState(114)
			p.Statement()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserRETURN {
		{
			p.SetState(120)
			p.ReturnStmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStmt() IIfStmtContext
	FunctionCall() IFunctionCallContext
	MethodCall() IMethodCallContext
	ThreeLevelCall() IThreeLevelCallContext
	Assignment() IAssignmentContext
	ConcStatement() IConcStatementContext
	ForStmt() IForStmtContext
	BreakStmt() IBreakStmtContext
	ForRangeStmt() IForRangeStmtContext
	ContinueStmt() IContinueStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
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

func (s *StatementContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *StatementContext) ThreeLevelCall() IThreeLevelCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ConcStatement() IConcStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcStatementContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ForRangeStmt() IForRangeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForRangeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForRangeStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gengineParserRULE_statement)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.IfStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.FunctionCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.MethodCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.ThreeLevelCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Assignment()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.ConcStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(130)
			p.BreakStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(131)
			p.ForRangeStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(132)
			p.ContinueStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConcStatementContext is an interface to support dynamic dispatch.
type IConcStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONC() antlr.TerminalNode
	LR_BRACE() antlr.TerminalNode
	RR_BRACE() antlr.TerminalNode
	AllFunctionCall() []IFunctionCallContext
	FunctionCall(i int) IFunctionCallContext
	AllMethodCall() []IMethodCallContext
	MethodCall(i int) IMethodCallContext
	AllThreeLevelCall() []IThreeLevelCallContext
	ThreeLevelCall(i int) IThreeLevelCallContext
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext

	// IsConcStatementContext differentiates from other interfaces.
	IsConcStatementContext()
}

type ConcStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcStatementContext() *ConcStatementContext {
	var p = new(ConcStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_concStatement
	return p
}

func InitEmptyConcStatementContext(p *ConcStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_concStatement
}

func (*ConcStatementContext) IsConcStatementContext() {}

func NewConcStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcStatementContext {
	var p = new(ConcStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_concStatement

	return p
}

func (s *ConcStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcStatementContext) CONC() antlr.TerminalNode {
	return s.GetToken(gengineParserCONC, 0)
}

func (s *ConcStatementContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ConcStatementContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ConcStatementContext) AllFunctionCall() []IFunctionCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionCallContext); ok {
			len++
		}
	}

	tst := make([]IFunctionCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionCallContext); ok {
			tst[i] = t.(IFunctionCallContext)
			i++
		}
	}

	return tst
}

func (s *ConcStatementContext) FunctionCall(i int) IFunctionCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
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

	return t.(IFunctionCallContext)
}

func (s *ConcStatementContext) AllMethodCall() []IMethodCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodCallContext); ok {
			len++
		}
	}

	tst := make([]IMethodCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodCallContext); ok {
			tst[i] = t.(IMethodCallContext)
			i++
		}
	}

	return tst
}

func (s *ConcStatementContext) MethodCall(i int) IMethodCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
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

	return t.(IMethodCallContext)
}

func (s *ConcStatementContext) AllThreeLevelCall() []IThreeLevelCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
			len++
		}
	}

	tst := make([]IThreeLevelCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThreeLevelCallContext); ok {
			tst[i] = t.(IThreeLevelCallContext)
			i++
		}
	}

	return tst
}

func (s *ConcStatementContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
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

	return t.(IThreeLevelCallContext)
}

func (s *ConcStatementContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ConcStatementContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *ConcStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConcStatement(s)
	}
}

func (s *ConcStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConcStatement(s)
	}
}

func (s *ConcStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConcStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ConcStatement() (localctx IConcStatementContext) {
	localctx = NewConcStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gengineParserRULE_concStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(gengineParserCONC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-29)) & ^0x3f) == 0 && ((int64(1)<<(_la-29))&51539607553) != 0 {
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(137)
				p.FunctionCall()
			}

		case 2:
			{
				p.SetState(138)
				p.MethodCall()
			}

		case 3:
			{
				p.SetState(139)
				p.ThreeLevelCall()
			}

		case 4:
			{
				p.SetState(140)
				p.Assignment()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MathExpression() IMathExpressionContext
	ExpressionAtom() IExpressionAtomContext
	NotOperator() INotOperatorContext
	LR_BRACKET() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RR_BRACKET() antlr.TerminalNode
	ComparisonOperator() IComparisonOperatorContext
	LogicalOperator() ILogicalOperatorContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) NotOperator() INotOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
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

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *gengineParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, gengineParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(149)
			p.mathExpression(0)
		}

	case 2:
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(150)
				p.NotOperator()
			}

		}
		{
			p.SetState(153)
			p.ExpressionAtom()
		}

	case 3:
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(154)
				p.NotOperator()
			}

		}
		{
			p.SetState(157)
			p.Match(gengineParserLR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.expression(0)
		}
		{
			p.SetState(159)
			p.Match(gengineParserRR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(164)
					p.ComparisonOperator()
				}
				{
					p.SetState(165)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(168)
					p.LogicalOperator()
				}
				{
					p.SetState(169)
					p.expression(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMathExpressionContext is an interface to support dynamic dispatch.
type IMathExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionAtom() IExpressionAtomContext
	LR_BRACKET() antlr.TerminalNode
	AllMathExpression() []IMathExpressionContext
	MathExpression(i int) IMathExpressionContext
	RR_BRACKET() antlr.TerminalNode
	MathMdOperator() IMathMdOperatorContext
	MathPmOperator() IMathPmOperatorContext

	// IsMathExpressionContext differentiates from other interfaces.
	IsMathExpressionContext()
}

type MathExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpressionContext() *MathExpressionContext {
	var p = new(MathExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathExpression
	return p
}

func InitEmptyMathExpressionContext(p *MathExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathExpression
}

func (*MathExpressionContext) IsMathExpressionContext() {}

func NewMathExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpressionContext {
	var p = new(MathExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathExpression

	return p
}

func (s *MathExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *MathExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MathExpressionContext) AllMathExpression() []IMathExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMathExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMathExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMathExpressionContext); ok {
			tst[i] = t.(IMathExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MathExpressionContext) MathExpression(i int) IMathExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
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

	return t.(IMathExpressionContext)
}

func (s *MathExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MathExpressionContext) MathMdOperator() IMathMdOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathMdOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathMdOperatorContext)
}

func (s *MathExpressionContext) MathPmOperator() IMathPmOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathPmOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathPmOperatorContext)
}

func (s *MathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathExpression(s)
	}
}

func (s *MathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathExpression(s)
	}
}

func (s *MathExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathExpression() (localctx IMathExpressionContext) {
	return p.mathExpression(0)
}

func (p *gengineParser) mathExpression(_p int) (localctx IMathExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMathExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, gengineParserRULE_mathExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gengineParserT__1, gengineParserT__2, gengineParserT__3, gengineParserT__4, gengineParserT__5, gengineParserTRUE, gengineParserFALSE, gengineParserSIMPLENAME, gengineParserMINUS, gengineParserDQUOTA_STRING, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME, gengineParserREAL_LITERAL:
		{
			p.SetState(177)
			p.ExpressionAtom()
		}

	case gengineParserLR_BRACKET:
		{
			p.SetState(178)
			p.Match(gengineParserLR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.mathExpression(0)
		}
		{
			p.SetState(180)
			p.Match(gengineParserRR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(185)
					p.MathMdOperator()
				}
				{
					p.SetState(186)
					p.mathExpression(5)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(189)
					p.MathPmOperator()
				}
				{
					p.SetState(190)
					p.mathExpression(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	MethodCall() IMethodCallContext
	ThreeLevelCall() IThreeLevelCallContext
	Constant() IConstantContext
	MapVar() IMapVarContext
	Variable() IVariableContext

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_expressionAtom
	return p
}

func InitEmptyExpressionAtomContext(p *ExpressionAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_expressionAtom
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
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

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) ThreeLevelCall() IThreeLevelCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) MapVar() IMapVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gengineParserRULE_expressionAtom)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)
			p.ThreeLevelCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(200)
			p.Constant()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)
			p.MapVar()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(202)
			p.Variable()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignOperator() IAssignOperatorContext
	MapVar() IMapVarContext
	Variable() IVariableContext
	MathExpression() IMathExpressionContext
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AssignOperator() IAssignOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignOperatorContext)
}

func (s *AssignmentContext) MapVar() IMapVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *AssignmentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gengineParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(205)
			p.MapVar()
		}

	case 2:
		{
			p.SetState(206)
			p.Variable()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(209)
		p.AssignOperator()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(210)
			p.mathExpression(0)
		}

	case 2:
		{
			p.SetState(211)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(gengineParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gengineParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(gengineParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-1008805765560926177) != 0 {
		{
			p.SetState(215)
			p.expression(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	LR_BRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RR_BRACE() antlr.TerminalNode
	AllElseIfStmt() []IElseIfStmtContext
	ElseIfStmt(i int) IElseIfStmtContext
	ElseStmt() IElseStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *IfStmtContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *IfStmtContext) AllElseIfStmt() []IElseIfStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStmtContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStmtContext); ok {
			tst[i] = t.(IElseIfStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStmtContext); ok {
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

	return t.(IElseIfStmtContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gengineParserRULE_ifStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(gengineParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.expression(0)
	}
	{
		p.SetState(220)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Statements()
	}
	{
		p.SetState(222)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(223)
				p.ElseIfStmt()
			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserELSE {
		{
			p.SetState(229)
			p.ElseStmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	LR_BRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RR_BRACE() antlr.TerminalNode

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_elseIfStmt
	return p
}

func InitEmptyElseIfStmtContext(p *ElseIfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_elseIfStmt
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseIfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *ElseIfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseIfStmtContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseIfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gengineParserRULE_elseIfStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(gengineParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(gengineParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.expression(0)
	}
	{
		p.SetState(235)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Statements()
	}
	{
		p.SetState(237)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	LR_BRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RR_BRACE() antlr.TerminalNode

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_elseStmt
	return p
}

func InitEmptyElseStmtContext(p *ElseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_elseStmt
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseStmtContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gengineParserRULE_elseStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(gengineParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Statements()
	}
	{
		p.SetState(242)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	Expression() IExpressionContext
	LR_BRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RR_BRACE() antlr.TerminalNode

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(gengineParserFOR, 0)
}

func (s *ForStmtContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(gengineParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(gengineParserSEMICOLON, i)
}

func (s *ForStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ForStmtContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ForStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gengineParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(gengineParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Assignment()
	}
	{
		p.SetState(246)
		p.Match(gengineParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.expression(0)
	}
	{
		p.SetState(248)
		p.Match(gengineParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Assignment()
	}
	{
		p.SetState(250)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Statements()
	}
	{
		p.SetState(252)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(gengineParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gengineParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(gengineParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForRangeStmtContext is an interface to support dynamic dispatch.
type IForRangeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORRANGE() antlr.TerminalNode
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
	RangeOperator() IRangeOperatorContext
	LR_BRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RR_BRACE() antlr.TerminalNode

	// IsForRangeStmtContext differentiates from other interfaces.
	IsForRangeStmtContext()
}

type ForRangeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForRangeStmtContext() *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_forRangeStmt
	return p
}

func InitEmptyForRangeStmtContext(p *ForRangeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_forRangeStmt
}

func (*ForRangeStmtContext) IsForRangeStmtContext() {}

func NewForRangeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_forRangeStmt

	return p
}

func (s *ForRangeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForRangeStmtContext) FORRANGE() antlr.TerminalNode {
	return s.GetToken(gengineParserFORRANGE, 0)
}

func (s *ForRangeStmtContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeStmtContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *ForRangeStmtContext) RangeOperator() IRangeOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeOperatorContext)
}

func (s *ForRangeStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ForRangeStmtContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ForRangeStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ForRangeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForRangeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitForRangeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ForRangeStmt() (localctx IForRangeStmtContext) {
	localctx = NewForRangeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gengineParserRULE_forRangeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(gengineParserFORRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Variable()
	}
	{
		p.SetState(258)
		p.RangeOperator()
	}
	{
		p.SetState(259)
		p.Variable()
	}
	{
		p.SetState(260)
		p.Match(gengineParserLR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Statements()
	}
	{
		p.SetState(262)
		p.Match(gengineParserRR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(gengineParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gengineParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(gengineParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BooleanLiteral() IBooleanLiteralContext
	RealLiteral() IRealLiteralContext
	StringLiteral() IStringLiteralContext
	AtName() IAtNameContext
	AtId() IAtIdContext
	AtCode() IAtCodeContext
	AtDesc() IAtDescContext
	AtSal() IAtSalContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) AtName() IAtNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtNameContext)
}

func (s *ConstantContext) AtId() IAtIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtIdContext)
}

func (s *ConstantContext) AtCode() IAtCodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtCodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtCodeContext)
}

func (s *ConstantContext) AtDesc() IAtDescContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtDescContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtDescContext)
}

func (s *ConstantContext) AtSal() IAtSalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtSalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtSalContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gengineParserRULE_constant)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gengineParserTRUE, gengineParserFALSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.BooleanLiteral()
		}

	case gengineParserMINUS, gengineParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.RealLiteral()
		}

	case gengineParserDQUOTA_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)
			p.StringLiteral()
		}

	case gengineParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(269)
			p.AtName()
		}

	case gengineParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(270)
			p.AtId()
		}

	case gengineParserT__3:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(271)
			p.AtCode()
		}

	case gengineParserT__4:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(272)
			p.AtDesc()
		}

	case gengineParserT__5:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(273)
			p.AtSal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
	AllFunctionCall() []IFunctionCallContext
	FunctionCall(i int) IFunctionCallContext
	AllMethodCall() []IMethodCallContext
	MethodCall(i int) IMethodCallContext
	AllThreeLevelCall() []IThreeLevelCallContext
	ThreeLevelCall(i int) IThreeLevelCallContext
	AllMapVar() []IMapVarContext
	MapVar(i int) IMapVarContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_functionArgs
	return p
}

func InitEmptyFunctionArgsContext(p *FunctionArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_functionArgs
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
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

	return t.(IConstantContext)
}

func (s *FunctionArgsContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *FunctionArgsContext) AllFunctionCall() []IFunctionCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionCallContext); ok {
			len++
		}
	}

	tst := make([]IFunctionCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionCallContext); ok {
			tst[i] = t.(IFunctionCallContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
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

	return t.(IFunctionCallContext)
}

func (s *FunctionArgsContext) AllMethodCall() []IMethodCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodCallContext); ok {
			len++
		}
	}

	tst := make([]IMethodCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodCallContext); ok {
			tst[i] = t.(IMethodCallContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) MethodCall(i int) IMethodCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
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

	return t.(IMethodCallContext)
}

func (s *FunctionArgsContext) AllThreeLevelCall() []IThreeLevelCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
			len++
		}
	}

	tst := make([]IThreeLevelCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThreeLevelCallContext); ok {
			tst[i] = t.(IThreeLevelCallContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThreeLevelCallContext); ok {
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

	return t.(IThreeLevelCallContext)
}

func (s *FunctionArgsContext) AllMapVar() []IMapVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapVarContext); ok {
			len++
		}
	}

	tst := make([]IMapVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapVarContext); ok {
			tst[i] = t.(IMapVarContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) MapVar(i int) IMapVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapVarContext); ok {
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

	return t.(IMapVarContext)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gengineParserRULE_functionArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(276)
			p.Constant()
		}

	case 2:
		{
			p.SetState(277)
			p.Variable()
		}

	case 3:
		{
			p.SetState(278)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(279)
			p.MethodCall()
		}

	case 5:
		{
			p.SetState(280)
			p.ThreeLevelCall()
		}

	case 6:
		{
			p.SetState(281)
			p.MapVar()
		}

	case 7:
		{
			p.SetState(282)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gengineParserT__0 {
		{
			p.SetState(285)
			p.Match(gengineParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(286)
				p.Constant()
			}

		case 2:
			{
				p.SetState(287)
				p.Variable()
			}

		case 3:
			{
				p.SetState(288)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(289)
				p.MethodCall()
			}

		case 5:
			{
				p.SetState(290)
				p.ThreeLevelCall()
			}

		case 6:
			{
				p.SetState(291)
				p.MapVar()
			}

		case 7:
			{
				p.SetState(292)
				p.expression(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *IntegerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gengineParserRULE_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(300)
			p.Match(gengineParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(303)
		p.Match(gengineParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REAL_LITERAL() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_realLiteral
	return p
}

func InitEmptyRealLiteralContext(p *RealLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_realLiteral
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(gengineParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gengineParserRULE_realLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(305)
			p.Match(gengineParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(308)
		p.Match(gengineParserREAL_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DQUOTA_STRING() antlr.TerminalNode

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(gengineParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gengineParserRULE_stringLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(gengineParserDQUOTA_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gengineParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gengineParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gengineParserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserTRUE || _la == gengineParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIMPLENAME() antlr.TerminalNode
	LR_BRACKET() antlr.TerminalNode
	RR_BRACKET() antlr.TerminalNode
	FunctionArgs() IFunctionArgsContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gengineParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(gengineParserSIMPLENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(gengineParserLR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-1008805765560926177) != 0 {
		{
			p.SetState(316)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(319)
		p.Match(gengineParserRR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOTTEDNAME() antlr.TerminalNode
	LR_BRACKET() antlr.TerminalNode
	RR_BRACKET() antlr.TerminalNode
	FunctionArgs() IFunctionArgsContext

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_methodCall
	return p
}

func InitEmptyMethodCallContext(p *MethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_methodCall
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *MethodCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MethodCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gengineParserRULE_methodCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(gengineParserDOTTEDNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(gengineParserLR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-1008805765560926177) != 0 {
		{
			p.SetState(323)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(326)
		p.Match(gengineParserRR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IThreeLevelCallContext is an interface to support dynamic dispatch.
type IThreeLevelCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLEDOTTEDNAME() antlr.TerminalNode
	LR_BRACKET() antlr.TerminalNode
	RR_BRACKET() antlr.TerminalNode
	FunctionArgs() IFunctionArgsContext

	// IsThreeLevelCallContext differentiates from other interfaces.
	IsThreeLevelCallContext()
}

type ThreeLevelCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThreeLevelCallContext() *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_threeLevelCall
	return p
}

func InitEmptyThreeLevelCallContext(p *ThreeLevelCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_threeLevelCall
}

func (*ThreeLevelCallContext) IsThreeLevelCallContext() {}

func NewThreeLevelCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_threeLevelCall

	return p
}

func (s *ThreeLevelCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ThreeLevelCallContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *ThreeLevelCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *ThreeLevelCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreeLevelCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThreeLevelCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitThreeLevelCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ThreeLevelCall() (localctx IThreeLevelCallContext) {
	localctx = NewThreeLevelCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gengineParserRULE_threeLevelCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(gengineParserDOUBLEDOTTEDNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(gengineParserLR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-1008805765560926177) != 0 {
		{
			p.SetState(330)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(333)
		p.Match(gengineParserRR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIMPLENAME() antlr.TerminalNode
	DOTTEDNAME() antlr.TerminalNode
	DOUBLEDOTTEDNAME() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *VariableContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *VariableContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gengineParserRULE_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-29)) & ^0x3f) == 0 && ((int64(1)<<(_la-29))&51539607553) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMathPmOperatorContext is an interface to support dynamic dispatch.
type IMathPmOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsMathPmOperatorContext differentiates from other interfaces.
	IsMathPmOperatorContext()
}

type MathPmOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathPmOperatorContext() *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathPmOperator
	return p
}

func InitEmptyMathPmOperatorContext(p *MathPmOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathPmOperator
}

func (*MathPmOperatorContext) IsMathPmOperatorContext() {}

func NewMathPmOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathPmOperator

	return p
}

func (s *MathPmOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathPmOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUS, 0)
}

func (s *MathPmOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *MathPmOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathPmOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathPmOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathPmOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathPmOperator() (localctx IMathPmOperatorContext) {
	localctx = NewMathPmOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gengineParserRULE_mathPmOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserPLUS || _la == gengineParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMathMdOperatorContext is an interface to support dynamic dispatch.
type IMathMdOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsMathMdOperatorContext differentiates from other interfaces.
	IsMathMdOperatorContext()
}

type MathMdOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathMdOperatorContext() *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathMdOperator
	return p
}

func InitEmptyMathMdOperatorContext(p *MathMdOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mathMdOperator
}

func (*MathMdOperatorContext) IsMathMdOperatorContext() {}

func NewMathMdOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathMdOperator

	return p
}

func (s *MathMdOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathMdOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(gengineParserMUL, 0)
}

func (s *MathMdOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(gengineParserDIV, 0)
}

func (s *MathMdOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathMdOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathMdOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathMdOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathMdOperator() (localctx IMathMdOperatorContext) {
	localctx = NewMathMdOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gengineParserRULE_mathMdOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserDIV || _la == gengineParserMUL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LTE() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	NOTEQUALS() antlr.TerminalNode
	GT_STR() antlr.TerminalNode
	LT_STR() antlr.TerminalNode
	GTE_STR() antlr.TerminalNode
	LTE_STR() antlr.TerminalNode
	EQUALS_STR() antlr.TerminalNode
	NOTEQUALS_STR() antlr.TerminalNode

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(gengineParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(gengineParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GT_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserGT_STR, 0)
}

func (s *ComparisonOperatorContext) LT_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserLT_STR, 0)
}

func (s *ComparisonOperatorContext) GTE_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE_STR, 0)
}

func (s *ComparisonOperatorContext) LTE_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE_STR, 0)
}

func (s *ComparisonOperatorContext) EQUALS_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS_STR, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS_STR, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gengineParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279241593716736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	AND_STR() antlr.TerminalNode
	OR_STR() antlr.TerminalNode
	AND_SINGLE() antlr.TerminalNode
	OR_SINGLE() antlr.TerminalNode

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_logicalOperator
	return p
}

func InitEmptyLogicalOperatorContext(p *LogicalOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_logicalOperator
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(gengineParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(gengineParserOR, 0)
}

func (s *LogicalOperatorContext) AND_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserAND_STR, 0)
}

func (s *LogicalOperatorContext) OR_STR() antlr.TerminalNode {
	return s.GetToken(gengineParserOR_STR, 0)
}

func (s *LogicalOperatorContext) AND_SINGLE() antlr.TerminalNode {
	return s.GetToken(gengineParserAND_SINGLE, 0)
}

func (s *LogicalOperatorContext) OR_SINGLE() antlr.TerminalNode {
	return s.GetToken(gengineParserOR_SINGLE, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gengineParserRULE_logicalOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignOperatorContext is an interface to support dynamic dispatch.
type IAssignOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	SET() antlr.TerminalNode
	PLUSEQUAL() antlr.TerminalNode
	MINUSEQUAL() antlr.TerminalNode
	MULTIEQUAL() antlr.TerminalNode
	DIVEQUAL() antlr.TerminalNode

	// IsAssignOperatorContext differentiates from other interfaces.
	IsAssignOperatorContext()
}

type AssignOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOperatorContext() *AssignOperatorContext {
	var p = new(AssignOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_assignOperator
	return p
}

func InitEmptyAssignOperatorContext(p *AssignOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_assignOperator
}

func (*AssignOperatorContext) IsAssignOperatorContext() {}

func NewAssignOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOperatorContext {
	var p = new(AssignOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignOperator

	return p
}

func (s *AssignOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *AssignOperatorContext) SET() antlr.TerminalNode {
	return s.GetToken(gengineParserSET, 0)
}

func (s *AssignOperatorContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUSEQUAL, 0)
}

func (s *AssignOperatorContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUSEQUAL, 0)
}

func (s *AssignOperatorContext) MULTIEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMULTIEQUAL, 0)
}

func (s *AssignOperatorContext) DIVEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserDIVEQUAL, 0)
}

func (s *AssignOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignOperator(s)
	}
}

func (s *AssignOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignOperator(s)
	}
}

func (s *AssignOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AssignOperator() (localctx IAssignOperatorContext) {
	localctx = NewAssignOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gengineParserRULE_assignOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeOperatorContext is an interface to support dynamic dispatch.
type IRangeOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode

	// IsRangeOperatorContext differentiates from other interfaces.
	IsRangeOperatorContext()
}

type RangeOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeOperatorContext() *RangeOperatorContext {
	var p = new(RangeOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_rangeOperator
	return p
}

func InitEmptyRangeOperatorContext(p *RangeOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_rangeOperator
}

func (*RangeOperatorContext) IsRangeOperatorContext() {}

func NewRangeOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeOperatorContext {
	var p = new(RangeOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_rangeOperator

	return p
}

func (s *RangeOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *RangeOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRangeOperator(s)
	}
}

func (s *RangeOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRangeOperator(s)
	}
}

func (s *RangeOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRangeOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RangeOperator() (localctx IRangeOperatorContext) {
	localctx = NewRangeOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gengineParserRULE_rangeOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(gengineParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotOperatorContext is an interface to support dynamic dispatch.
type INotOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode

	// IsNotOperatorContext differentiates from other interfaces.
	IsNotOperatorContext()
}

type NotOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOperatorContext() *NotOperatorContext {
	var p = new(NotOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_notOperator
	return p
}

func InitEmptyNotOperatorContext(p *NotOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_notOperator
}

func (*NotOperatorContext) IsNotOperatorContext() {}

func NewNotOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOperatorContext {
	var p = new(NotOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_notOperator

	return p
}

func (s *NotOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(gengineParserNOT, 0)
}

func (s *NotOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterNotOperator(s)
	}
}

func (s *NotOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitNotOperator(s)
	}
}

func (s *NotOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitNotOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) NotOperator() (localctx INotOperatorContext) {
	localctx = NewNotOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, gengineParserRULE_notOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(gengineParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapVarContext is an interface to support dynamic dispatch.
type IMapVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
	LSQARE() antlr.TerminalNode
	RSQARE() antlr.TerminalNode
	Integer() IIntegerContext
	StringLiteral() IStringLiteralContext

	// IsMapVarContext differentiates from other interfaces.
	IsMapVarContext()
}

type MapVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapVarContext() *MapVarContext {
	var p = new(MapVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mapVar
	return p
}

func InitEmptyMapVarContext(p *MapVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_mapVar
}

func (*MapVarContext) IsMapVarContext() {}

func NewMapVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapVarContext {
	var p = new(MapVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mapVar

	return p
}

func (s *MapVarContext) GetParser() antlr.Parser { return s.parser }

func (s *MapVarContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *MapVarContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *MapVarContext) LSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserLSQARE, 0)
}

func (s *MapVarContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserRSQARE, 0)
}

func (s *MapVarContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *MapVarContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *MapVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMapVar(s)
	}
}

func (s *MapVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMapVar(s)
	}
}

func (s *MapVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMapVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MapVar() (localctx IMapVarContext) {
	localctx = NewMapVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, gengineParserRULE_mapVar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Variable()
	}
	{
		p.SetState(352)
		p.Match(gengineParserLSQARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gengineParserINT, gengineParserMINUS:
		{
			p.SetState(353)
			p.Integer()
		}

	case gengineParserDQUOTA_STRING:
		{
			p.SetState(354)
			p.StringLiteral()
		}

	case gengineParserSIMPLENAME, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME:
		{
			p.SetState(355)
			p.Variable()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(358)
		p.Match(gengineParserRSQARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtNameContext is an interface to support dynamic dispatch.
type IAtNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtNameContext differentiates from other interfaces.
	IsAtNameContext()
}

type AtNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtNameContext() *AtNameContext {
	var p = new(AtNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atName
	return p
}

func InitEmptyAtNameContext(p *AtNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atName
}

func (*AtNameContext) IsAtNameContext() {}

func NewAtNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtNameContext {
	var p = new(AtNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atName

	return p
}

func (s *AtNameContext) GetParser() antlr.Parser { return s.parser }
func (s *AtNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtName(s)
	}
}

func (s *AtNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtName(s)
	}
}

func (s *AtNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtName() (localctx IAtNameContext) {
	localctx = NewAtNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, gengineParserRULE_atName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(gengineParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtIdContext is an interface to support dynamic dispatch.
type IAtIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtIdContext differentiates from other interfaces.
	IsAtIdContext()
}

type AtIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtIdContext() *AtIdContext {
	var p = new(AtIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atId
	return p
}

func InitEmptyAtIdContext(p *AtIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atId
}

func (*AtIdContext) IsAtIdContext() {}

func NewAtIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtIdContext {
	var p = new(AtIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atId

	return p
}

func (s *AtIdContext) GetParser() antlr.Parser { return s.parser }
func (s *AtIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtId(s)
	}
}

func (s *AtIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtId(s)
	}
}

func (s *AtIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtId() (localctx IAtIdContext) {
	localctx = NewAtIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, gengineParserRULE_atId)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(gengineParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtCodeContext is an interface to support dynamic dispatch.
type IAtCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtCodeContext differentiates from other interfaces.
	IsAtCodeContext()
}

type AtCodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtCodeContext() *AtCodeContext {
	var p = new(AtCodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atCode
	return p
}

func InitEmptyAtCodeContext(p *AtCodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atCode
}

func (*AtCodeContext) IsAtCodeContext() {}

func NewAtCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtCodeContext {
	var p = new(AtCodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atCode

	return p
}

func (s *AtCodeContext) GetParser() antlr.Parser { return s.parser }
func (s *AtCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtCode(s)
	}
}

func (s *AtCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtCode(s)
	}
}

func (s *AtCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtCode() (localctx IAtCodeContext) {
	localctx = NewAtCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, gengineParserRULE_atCode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(gengineParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtDescContext is an interface to support dynamic dispatch.
type IAtDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtDescContext differentiates from other interfaces.
	IsAtDescContext()
}

type AtDescContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtDescContext() *AtDescContext {
	var p = new(AtDescContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atDesc
	return p
}

func InitEmptyAtDescContext(p *AtDescContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atDesc
}

func (*AtDescContext) IsAtDescContext() {}

func NewAtDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDescContext {
	var p = new(AtDescContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atDesc

	return p
}

func (s *AtDescContext) GetParser() antlr.Parser { return s.parser }
func (s *AtDescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtDescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtDesc(s)
	}
}

func (s *AtDescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtDesc(s)
	}
}

func (s *AtDescContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtDesc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtDesc() (localctx IAtDescContext) {
	localctx = NewAtDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, gengineParserRULE_atDesc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(gengineParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtSalContext is an interface to support dynamic dispatch.
type IAtSalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtSalContext differentiates from other interfaces.
	IsAtSalContext()
}

type AtSalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtSalContext() *AtSalContext {
	var p = new(AtSalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atSal
	return p
}

func InitEmptyAtSalContext(p *AtSalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gengineParserRULE_atSal
}

func (*AtSalContext) IsAtSalContext() {}

func NewAtSalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtSalContext {
	var p = new(AtSalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atSal

	return p
}

func (s *AtSalContext) GetParser() antlr.Parser { return s.parser }
func (s *AtSalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtSalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtSalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtSal(s)
	}
}

func (s *AtSalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtSal(s)
	}
}

func (s *AtSalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtSal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtSal() (localctx IAtSalContext) {
	localctx = NewAtSalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, gengineParserRULE_atSal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(gengineParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *gengineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 10:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gengineParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gengineParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
