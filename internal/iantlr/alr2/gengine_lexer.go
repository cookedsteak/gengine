// Code generated from gengine.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GengineLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func genginelexerLexerInit() {
	staticData := &GengineLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'@name'", "'@id'", "'@code'", "'@desc'", "'@sal'", "", "",
		"'&&'", "'||'", "'and'", "'or'", "'&'", "'|'", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'",
		"'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='",
		"'+='", "'-='", "'*='", "'/='", "'['", "']'", "';'", "'{'", "'}'", "'('",
		"')'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "AND_STR", "OR_STR",
		"AND_SINGLE", "OR_SINGLE", "CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK",
		"FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE",
		"BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL",
		"EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET",
		"PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE",
		"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT",
		"DQUOTA_STRING", "DOTTEDNAME", "DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "DEC_DIGIT", "A", "B",
		"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART",
		"NIL", "RULE", "AND", "OR", "AND_STR", "OR_STR", "AND_SINGLE", "OR_SINGLE",
		"CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK", "FORRANGE", "CONTINUE",
		"TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", "SIMPLENAME",
		"INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE",
		"NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL",
		"DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE",
		"LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "DOUBLEDOTTEDNAME",
		"REAL_LITERAL", "SL_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 61, 543, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 265, 8, 33, 1,
		33, 4, 33, 268, 8, 33, 11, 33, 12, 33, 269, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 56, 4, 56, 382, 8, 56, 11, 56, 12, 56, 383, 1, 56, 5, 56, 387, 8, 56,
		10, 56, 12, 56, 390, 9, 56, 1, 57, 4, 57, 393, 8, 57, 11, 57, 12, 57, 394,
		1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1,
		62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66,
		1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1,
		71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 74, 1, 74,
		1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 78, 1, 78, 1, 79, 1,
		79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 1, 83,
		1, 83, 1, 83, 5, 83, 462, 8, 83, 10, 83, 12, 83, 465, 9, 83, 1, 83, 1,
		83, 1, 84, 1, 84, 1, 84, 1, 84, 1, 85, 1, 85, 1, 85, 1, 85, 1, 85, 1, 85,
		1, 86, 4, 86, 480, 8, 86, 11, 86, 12, 86, 481, 3, 86, 484, 8, 86, 1, 86,
		1, 86, 4, 86, 488, 8, 86, 11, 86, 12, 86, 489, 1, 86, 4, 86, 493, 8, 86,
		11, 86, 12, 86, 494, 1, 86, 1, 86, 1, 86, 1, 86, 4, 86, 501, 8, 86, 11,
		86, 12, 86, 502, 3, 86, 505, 8, 86, 1, 86, 1, 86, 4, 86, 509, 8, 86, 11,
		86, 12, 86, 510, 1, 86, 1, 86, 1, 86, 4, 86, 516, 8, 86, 11, 86, 12, 86,
		517, 1, 86, 1, 86, 3, 86, 522, 8, 86, 1, 87, 1, 87, 1, 87, 1, 87, 5, 87,
		528, 8, 87, 10, 87, 12, 87, 531, 9, 87, 1, 87, 1, 87, 1, 87, 1, 87, 1,
		88, 4, 88, 538, 8, 88, 11, 88, 12, 88, 539, 1, 88, 1, 88, 1, 529, 0, 89,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0,
		23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0, 43,
		0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 0, 55, 0, 57, 0, 59, 0, 61, 0, 63, 0,
		65, 0, 67, 0, 69, 7, 71, 8, 73, 9, 75, 10, 77, 11, 79, 12, 81, 13, 83,
		14, 85, 15, 87, 16, 89, 17, 91, 18, 93, 19, 95, 20, 97, 21, 99, 22, 101,
		23, 103, 24, 105, 25, 107, 26, 109, 27, 111, 28, 113, 29, 115, 30, 117,
		31, 119, 32, 121, 33, 123, 34, 125, 35, 127, 36, 129, 37, 131, 38, 133,
		39, 135, 40, 137, 41, 139, 42, 141, 43, 143, 44, 145, 45, 147, 46, 149,
		47, 151, 48, 153, 49, 155, 50, 157, 51, 159, 52, 161, 53, 163, 54, 165,
		55, 167, 56, 169, 57, 171, 58, 173, 59, 175, 60, 177, 61, 1, 0, 31, 1,
		0, 48, 57, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99,
		99, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102,
		102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105,
		105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108,
		108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111,
		111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114,
		114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117,
		117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120,
		120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92,
		3, 0, 9, 10, 13, 13, 32, 32, 535, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0,
		0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121,
		1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0,
		0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1,
		0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0,
		143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0,
		0, 0, 0, 151, 1, 0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157,
		1, 0, 0, 0, 0, 159, 1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0,
		0, 165, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0, 0, 169, 1, 0, 0, 0, 0, 171, 1,
		0, 0, 0, 0, 173, 1, 0, 0, 0, 0, 175, 1, 0, 0, 0, 0, 177, 1, 0, 0, 0, 1,
		179, 1, 0, 0, 0, 3, 181, 1, 0, 0, 0, 5, 187, 1, 0, 0, 0, 7, 191, 1, 0,
		0, 0, 9, 197, 1, 0, 0, 0, 11, 203, 1, 0, 0, 0, 13, 208, 1, 0, 0, 0, 15,
		210, 1, 0, 0, 0, 17, 212, 1, 0, 0, 0, 19, 214, 1, 0, 0, 0, 21, 216, 1,
		0, 0, 0, 23, 218, 1, 0, 0, 0, 25, 220, 1, 0, 0, 0, 27, 222, 1, 0, 0, 0,
		29, 224, 1, 0, 0, 0, 31, 226, 1, 0, 0, 0, 33, 228, 1, 0, 0, 0, 35, 230,
		1, 0, 0, 0, 37, 232, 1, 0, 0, 0, 39, 234, 1, 0, 0, 0, 41, 236, 1, 0, 0,
		0, 43, 238, 1, 0, 0, 0, 45, 240, 1, 0, 0, 0, 47, 242, 1, 0, 0, 0, 49, 244,
		1, 0, 0, 0, 51, 246, 1, 0, 0, 0, 53, 248, 1, 0, 0, 0, 55, 250, 1, 0, 0,
		0, 57, 252, 1, 0, 0, 0, 59, 254, 1, 0, 0, 0, 61, 256, 1, 0, 0, 0, 63, 258,
		1, 0, 0, 0, 65, 260, 1, 0, 0, 0, 67, 262, 1, 0, 0, 0, 69, 271, 1, 0, 0,
		0, 71, 275, 1, 0, 0, 0, 73, 280, 1, 0, 0, 0, 75, 283, 1, 0, 0, 0, 77, 286,
		1, 0, 0, 0, 79, 290, 1, 0, 0, 0, 81, 293, 1, 0, 0, 0, 83, 295, 1, 0, 0,
		0, 85, 297, 1, 0, 0, 0, 87, 302, 1, 0, 0, 0, 89, 305, 1, 0, 0, 0, 91, 310,
		1, 0, 0, 0, 93, 317, 1, 0, 0, 0, 95, 321, 1, 0, 0, 0, 97, 327, 1, 0, 0,
		0, 99, 336, 1, 0, 0, 0, 101, 345, 1, 0, 0, 0, 103, 350, 1, 0, 0, 0, 105,
		356, 1, 0, 0, 0, 107, 361, 1, 0, 0, 0, 109, 370, 1, 0, 0, 0, 111, 376,
		1, 0, 0, 0, 113, 381, 1, 0, 0, 0, 115, 392, 1, 0, 0, 0, 117, 396, 1, 0,
		0, 0, 119, 398, 1, 0, 0, 0, 121, 400, 1, 0, 0, 0, 123, 402, 1, 0, 0, 0,
		125, 404, 1, 0, 0, 0, 127, 407, 1, 0, 0, 0, 129, 409, 1, 0, 0, 0, 131,
		411, 1, 0, 0, 0, 133, 414, 1, 0, 0, 0, 135, 417, 1, 0, 0, 0, 137, 420,
		1, 0, 0, 0, 139, 422, 1, 0, 0, 0, 141, 425, 1, 0, 0, 0, 143, 427, 1, 0,
		0, 0, 145, 430, 1, 0, 0, 0, 147, 433, 1, 0, 0, 0, 149, 436, 1, 0, 0, 0,
		151, 439, 1, 0, 0, 0, 153, 441, 1, 0, 0, 0, 155, 443, 1, 0, 0, 0, 157,
		445, 1, 0, 0, 0, 159, 447, 1, 0, 0, 0, 161, 449, 1, 0, 0, 0, 163, 451,
		1, 0, 0, 0, 165, 453, 1, 0, 0, 0, 167, 455, 1, 0, 0, 0, 169, 468, 1, 0,
		0, 0, 171, 472, 1, 0, 0, 0, 173, 521, 1, 0, 0, 0, 175, 523, 1, 0, 0, 0,
		177, 537, 1, 0, 0, 0, 179, 180, 5, 44, 0, 0, 180, 2, 1, 0, 0, 0, 181, 182,
		5, 64, 0, 0, 182, 183, 5, 110, 0, 0, 183, 184, 5, 97, 0, 0, 184, 185, 5,
		109, 0, 0, 185, 186, 5, 101, 0, 0, 186, 4, 1, 0, 0, 0, 187, 188, 5, 64,
		0, 0, 188, 189, 5, 105, 0, 0, 189, 190, 5, 100, 0, 0, 190, 6, 1, 0, 0,
		0, 191, 192, 5, 64, 0, 0, 192, 193, 5, 99, 0, 0, 193, 194, 5, 111, 0, 0,
		194, 195, 5, 100, 0, 0, 195, 196, 5, 101, 0, 0, 196, 8, 1, 0, 0, 0, 197,
		198, 5, 64, 0, 0, 198, 199, 5, 100, 0, 0, 199, 200, 5, 101, 0, 0, 200,
		201, 5, 115, 0, 0, 201, 202, 5, 99, 0, 0, 202, 10, 1, 0, 0, 0, 203, 204,
		5, 64, 0, 0, 204, 205, 5, 115, 0, 0, 205, 206, 5, 97, 0, 0, 206, 207, 5,
		108, 0, 0, 207, 12, 1, 0, 0, 0, 208, 209, 7, 0, 0, 0, 209, 14, 1, 0, 0,
		0, 210, 211, 7, 1, 0, 0, 211, 16, 1, 0, 0, 0, 212, 213, 7, 2, 0, 0, 213,
		18, 1, 0, 0, 0, 214, 215, 7, 3, 0, 0, 215, 20, 1, 0, 0, 0, 216, 217, 7,
		4, 0, 0, 217, 22, 1, 0, 0, 0, 218, 219, 7, 5, 0, 0, 219, 24, 1, 0, 0, 0,
		220, 221, 7, 6, 0, 0, 221, 26, 1, 0, 0, 0, 222, 223, 7, 7, 0, 0, 223, 28,
		1, 0, 0, 0, 224, 225, 7, 8, 0, 0, 225, 30, 1, 0, 0, 0, 226, 227, 7, 9,
		0, 0, 227, 32, 1, 0, 0, 0, 228, 229, 7, 10, 0, 0, 229, 34, 1, 0, 0, 0,
		230, 231, 7, 11, 0, 0, 231, 36, 1, 0, 0, 0, 232, 233, 7, 12, 0, 0, 233,
		38, 1, 0, 0, 0, 234, 235, 7, 13, 0, 0, 235, 40, 1, 0, 0, 0, 236, 237, 7,
		14, 0, 0, 237, 42, 1, 0, 0, 0, 238, 239, 7, 15, 0, 0, 239, 44, 1, 0, 0,
		0, 240, 241, 7, 16, 0, 0, 241, 46, 1, 0, 0, 0, 242, 243, 7, 17, 0, 0, 243,
		48, 1, 0, 0, 0, 244, 245, 7, 18, 0, 0, 245, 50, 1, 0, 0, 0, 246, 247, 7,
		19, 0, 0, 247, 52, 1, 0, 0, 0, 248, 249, 7, 20, 0, 0, 249, 54, 1, 0, 0,
		0, 250, 251, 7, 21, 0, 0, 251, 56, 1, 0, 0, 0, 252, 253, 7, 22, 0, 0, 253,
		58, 1, 0, 0, 0, 254, 255, 7, 23, 0, 0, 255, 60, 1, 0, 0, 0, 256, 257, 7,
		24, 0, 0, 257, 62, 1, 0, 0, 0, 258, 259, 7, 25, 0, 0, 259, 64, 1, 0, 0,
		0, 260, 261, 7, 26, 0, 0, 261, 66, 1, 0, 0, 0, 262, 264, 7, 5, 0, 0, 263,
		265, 5, 45, 0, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267,
		1, 0, 0, 0, 266, 268, 3, 13, 6, 0, 267, 266, 1, 0, 0, 0, 268, 269, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 68, 1, 0, 0, 0,
		271, 272, 3, 41, 20, 0, 272, 273, 3, 31, 15, 0, 273, 274, 3, 37, 18, 0,
		274, 70, 1, 0, 0, 0, 275, 276, 3, 49, 24, 0, 276, 277, 3, 55, 27, 0, 277,
		278, 3, 37, 18, 0, 278, 279, 3, 23, 11, 0, 279, 72, 1, 0, 0, 0, 280, 281,
		5, 38, 0, 0, 281, 282, 5, 38, 0, 0, 282, 74, 1, 0, 0, 0, 283, 284, 5, 124,
		0, 0, 284, 285, 5, 124, 0, 0, 285, 76, 1, 0, 0, 0, 286, 287, 5, 97, 0,
		0, 287, 288, 5, 110, 0, 0, 288, 289, 5, 100, 0, 0, 289, 78, 1, 0, 0, 0,
		290, 291, 5, 111, 0, 0, 291, 292, 5, 114, 0, 0, 292, 80, 1, 0, 0, 0, 293,
		294, 5, 38, 0, 0, 294, 82, 1, 0, 0, 0, 295, 296, 5, 124, 0, 0, 296, 84,
		1, 0, 0, 0, 297, 298, 3, 19, 9, 0, 298, 299, 3, 43, 21, 0, 299, 300, 3,
		41, 20, 0, 300, 301, 3, 19, 9, 0, 301, 86, 1, 0, 0, 0, 302, 303, 3, 31,
		15, 0, 303, 304, 3, 25, 12, 0, 304, 88, 1, 0, 0, 0, 305, 306, 3, 23, 11,
		0, 306, 307, 3, 37, 18, 0, 307, 308, 3, 51, 25, 0, 308, 309, 3, 23, 11,
		0, 309, 90, 1, 0, 0, 0, 310, 311, 3, 49, 24, 0, 311, 312, 3, 23, 11, 0,
		312, 313, 3, 53, 26, 0, 313, 314, 3, 55, 27, 0, 314, 315, 3, 49, 24, 0,
		315, 316, 3, 41, 20, 0, 316, 92, 1, 0, 0, 0, 317, 318, 3, 25, 12, 0, 318,
		319, 3, 43, 21, 0, 319, 320, 3, 49, 24, 0, 320, 94, 1, 0, 0, 0, 321, 322,
		3, 17, 8, 0, 322, 323, 3, 49, 24, 0, 323, 324, 3, 23, 11, 0, 324, 325,
		3, 15, 7, 0, 325, 326, 3, 35, 17, 0, 326, 96, 1, 0, 0, 0, 327, 328, 3,
		25, 12, 0, 328, 329, 3, 43, 21, 0, 329, 330, 3, 49, 24, 0, 330, 331, 3,
		49, 24, 0, 331, 332, 3, 15, 7, 0, 332, 333, 3, 41, 20, 0, 333, 334, 3,
		27, 13, 0, 334, 335, 3, 23, 11, 0, 335, 98, 1, 0, 0, 0, 336, 337, 3, 19,
		9, 0, 337, 338, 3, 43, 21, 0, 338, 339, 3, 41, 20, 0, 339, 340, 3, 53,
		26, 0, 340, 341, 3, 31, 15, 0, 341, 342, 3, 41, 20, 0, 342, 343, 3, 55,
		27, 0, 343, 344, 3, 23, 11, 0, 344, 100, 1, 0, 0, 0, 345, 346, 3, 53, 26,
		0, 346, 347, 3, 49, 24, 0, 347, 348, 3, 55, 27, 0, 348, 349, 3, 23, 11,
		0, 349, 102, 1, 0, 0, 0, 350, 351, 3, 25, 12, 0, 351, 352, 3, 15, 7, 0,
		352, 353, 3, 37, 18, 0, 353, 354, 3, 51, 25, 0, 354, 355, 3, 23, 11, 0,
		355, 104, 1, 0, 0, 0, 356, 357, 3, 41, 20, 0, 357, 358, 3, 55, 27, 0, 358,
		359, 3, 37, 18, 0, 359, 360, 3, 37, 18, 0, 360, 106, 1, 0, 0, 0, 361, 362,
		3, 51, 25, 0, 362, 363, 3, 15, 7, 0, 363, 364, 3, 37, 18, 0, 364, 365,
		3, 31, 15, 0, 365, 366, 3, 23, 11, 0, 366, 367, 3, 41, 20, 0, 367, 368,
		3, 19, 9, 0, 368, 369, 3, 23, 11, 0, 369, 108, 1, 0, 0, 0, 370, 371, 3,
		17, 8, 0, 371, 372, 3, 23, 11, 0, 372, 373, 3, 27, 13, 0, 373, 374, 3,
		31, 15, 0, 374, 375, 3, 41, 20, 0, 375, 110, 1, 0, 0, 0, 376, 377, 3, 23,
		11, 0, 377, 378, 3, 41, 20, 0, 378, 379, 3, 21, 10, 0, 379, 112, 1, 0,
		0, 0, 380, 382, 7, 27, 0, 0, 381, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0,
		383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 388, 1, 0, 0, 0, 385,
		387, 7, 28, 0, 0, 386, 385, 1, 0, 0, 0, 387, 390, 1, 0, 0, 0, 388, 386,
		1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 114, 1, 0, 0, 0, 390, 388, 1, 0,
		0, 0, 391, 393, 2, 48, 57, 0, 392, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0,
		0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 116, 1, 0, 0, 0, 396,
		397, 5, 43, 0, 0, 397, 118, 1, 0, 0, 0, 398, 399, 5, 45, 0, 0, 399, 120,
		1, 0, 0, 0, 400, 401, 5, 47, 0, 0, 401, 122, 1, 0, 0, 0, 402, 403, 5, 42,
		0, 0, 403, 124, 1, 0, 0, 0, 404, 405, 5, 61, 0, 0, 405, 406, 5, 61, 0,
		0, 406, 126, 1, 0, 0, 0, 407, 408, 5, 62, 0, 0, 408, 128, 1, 0, 0, 0, 409,
		410, 5, 60, 0, 0, 410, 130, 1, 0, 0, 0, 411, 412, 5, 62, 0, 0, 412, 413,
		5, 61, 0, 0, 413, 132, 1, 0, 0, 0, 414, 415, 5, 60, 0, 0, 415, 416, 5,
		61, 0, 0, 416, 134, 1, 0, 0, 0, 417, 418, 5, 33, 0, 0, 418, 419, 5, 61,
		0, 0, 419, 136, 1, 0, 0, 0, 420, 421, 5, 33, 0, 0, 421, 138, 1, 0, 0, 0,
		422, 423, 5, 58, 0, 0, 423, 424, 5, 61, 0, 0, 424, 140, 1, 0, 0, 0, 425,
		426, 5, 61, 0, 0, 426, 142, 1, 0, 0, 0, 427, 428, 5, 43, 0, 0, 428, 429,
		5, 61, 0, 0, 429, 144, 1, 0, 0, 0, 430, 431, 5, 45, 0, 0, 431, 432, 5,
		61, 0, 0, 432, 146, 1, 0, 0, 0, 433, 434, 5, 42, 0, 0, 434, 435, 5, 61,
		0, 0, 435, 148, 1, 0, 0, 0, 436, 437, 5, 47, 0, 0, 437, 438, 5, 61, 0,
		0, 438, 150, 1, 0, 0, 0, 439, 440, 5, 91, 0, 0, 440, 152, 1, 0, 0, 0, 441,
		442, 5, 93, 0, 0, 442, 154, 1, 0, 0, 0, 443, 444, 5, 59, 0, 0, 444, 156,
		1, 0, 0, 0, 445, 446, 5, 123, 0, 0, 446, 158, 1, 0, 0, 0, 447, 448, 5,
		125, 0, 0, 448, 160, 1, 0, 0, 0, 449, 450, 5, 40, 0, 0, 450, 162, 1, 0,
		0, 0, 451, 452, 5, 41, 0, 0, 452, 164, 1, 0, 0, 0, 453, 454, 5, 46, 0,
		0, 454, 166, 1, 0, 0, 0, 455, 463, 5, 34, 0, 0, 456, 457, 5, 92, 0, 0,
		457, 462, 9, 0, 0, 0, 458, 459, 5, 34, 0, 0, 459, 462, 5, 34, 0, 0, 460,
		462, 8, 29, 0, 0, 461, 456, 1, 0, 0, 0, 461, 458, 1, 0, 0, 0, 461, 460,
		1, 0, 0, 0, 462, 465, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 463, 464, 1, 0,
		0, 0, 464, 466, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 466, 467, 5, 34, 0, 0,
		467, 168, 1, 0, 0, 0, 468, 469, 3, 113, 56, 0, 469, 470, 3, 165, 82, 0,
		470, 471, 3, 113, 56, 0, 471, 170, 1, 0, 0, 0, 472, 473, 3, 113, 56, 0,
		473, 474, 3, 165, 82, 0, 474, 475, 3, 113, 56, 0, 475, 476, 3, 165, 82,
		0, 476, 477, 3, 113, 56, 0, 477, 172, 1, 0, 0, 0, 478, 480, 3, 13, 6, 0,
		479, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 481,
		482, 1, 0, 0, 0, 482, 484, 1, 0, 0, 0, 483, 479, 1, 0, 0, 0, 483, 484,
		1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 487, 5, 46, 0, 0, 486, 488, 3, 13,
		6, 0, 487, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0,
		489, 490, 1, 0, 0, 0, 490, 522, 1, 0, 0, 0, 491, 493, 3, 13, 6, 0, 492,
		491, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 494, 495,
		1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 5, 46, 0, 0, 497, 498, 3, 67,
		33, 0, 498, 522, 1, 0, 0, 0, 499, 501, 3, 13, 6, 0, 500, 499, 1, 0, 0,
		0, 501, 502, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503,
		505, 1, 0, 0, 0, 504, 500, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 506,
		1, 0, 0, 0, 506, 508, 5, 46, 0, 0, 507, 509, 3, 13, 6, 0, 508, 507, 1,
		0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 510, 511, 1, 0, 0,
		0, 511, 512, 1, 0, 0, 0, 512, 513, 3, 67, 33, 0, 513, 522, 1, 0, 0, 0,
		514, 516, 3, 13, 6, 0, 515, 514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517,
		515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 520,
		3, 67, 33, 0, 520, 522, 1, 0, 0, 0, 521, 483, 1, 0, 0, 0, 521, 492, 1,
		0, 0, 0, 521, 504, 1, 0, 0, 0, 521, 515, 1, 0, 0, 0, 522, 174, 1, 0, 0,
		0, 523, 524, 5, 47, 0, 0, 524, 525, 5, 47, 0, 0, 525, 529, 1, 0, 0, 0,
		526, 528, 9, 0, 0, 0, 527, 526, 1, 0, 0, 0, 528, 531, 1, 0, 0, 0, 529,
		530, 1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 530, 532, 1, 0, 0, 0, 531, 529,
		1, 0, 0, 0, 532, 533, 5, 10, 0, 0, 533, 534, 1, 0, 0, 0, 534, 535, 6, 87,
		0, 0, 535, 176, 1, 0, 0, 0, 536, 538, 7, 30, 0, 0, 537, 536, 1, 0, 0, 0,
		538, 539, 1, 0, 0, 0, 539, 537, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540,
		541, 1, 0, 0, 0, 541, 542, 6, 88, 0, 0, 542, 178, 1, 0, 0, 0, 20, 0, 264,
		269, 383, 386, 388, 394, 461, 463, 481, 483, 489, 494, 502, 504, 510, 517,
		521, 529, 539, 1, 6, 0, 0,
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

// gengineLexerInit initializes any static state used to implement gengineLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgengineLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GengineLexerInit() {
	staticData := &GengineLexerLexerStaticData
	staticData.once.Do(genginelexerLexerInit)
}

// NewgengineLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewgengineLexer(input antlr.CharStream) *gengineLexer {
	GengineLexerInit()
	l := new(gengineLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GengineLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0             = 1
	gengineLexerT__1             = 2
	gengineLexerT__2             = 3
	gengineLexerT__3             = 4
	gengineLexerT__4             = 5
	gengineLexerT__5             = 6
	gengineLexerNIL              = 7
	gengineLexerRULE             = 8
	gengineLexerAND              = 9
	gengineLexerOR               = 10
	gengineLexerAND_STR          = 11
	gengineLexerOR_STR           = 12
	gengineLexerAND_SINGLE       = 13
	gengineLexerOR_SINGLE        = 14
	gengineLexerCONC             = 15
	gengineLexerIF               = 16
	gengineLexerELSE             = 17
	gengineLexerRETURN           = 18
	gengineLexerFOR              = 19
	gengineLexerBREAK            = 20
	gengineLexerFORRANGE         = 21
	gengineLexerCONTINUE         = 22
	gengineLexerTRUE             = 23
	gengineLexerFALSE            = 24
	gengineLexerNULL_LITERAL     = 25
	gengineLexerSALIENCE         = 26
	gengineLexerBEGIN            = 27
	gengineLexerEND              = 28
	gengineLexerSIMPLENAME       = 29
	gengineLexerINT              = 30
	gengineLexerPLUS             = 31
	gengineLexerMINUS            = 32
	gengineLexerDIV              = 33
	gengineLexerMUL              = 34
	gengineLexerEQUALS           = 35
	gengineLexerGT               = 36
	gengineLexerLT               = 37
	gengineLexerGTE              = 38
	gengineLexerLTE              = 39
	gengineLexerNOTEQUALS        = 40
	gengineLexerNOT              = 41
	gengineLexerASSIGN           = 42
	gengineLexerSET              = 43
	gengineLexerPLUSEQUAL        = 44
	gengineLexerMINUSEQUAL       = 45
	gengineLexerMULTIEQUAL       = 46
	gengineLexerDIVEQUAL         = 47
	gengineLexerLSQARE           = 48
	gengineLexerRSQARE           = 49
	gengineLexerSEMICOLON        = 50
	gengineLexerLR_BRACE         = 51
	gengineLexerRR_BRACE         = 52
	gengineLexerLR_BRACKET       = 53
	gengineLexerRR_BRACKET       = 54
	gengineLexerDOT              = 55
	gengineLexerDQUOTA_STRING    = 56
	gengineLexerDOTTEDNAME       = 57
	gengineLexerDOUBLEDOTTEDNAME = 58
	gengineLexerREAL_LITERAL     = 59
	gengineLexerSL_COMMENT       = 60
	gengineLexerWS               = 61
)
