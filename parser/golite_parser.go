// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
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

type GoliteParser struct {
	*antlr.BaseParser
}

var goliteparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'&&'", "'||'", "'!'", "'<'", "'>'", "'<='", "'>='", "'=='",
		"'!='", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'{'", "'}'", "';'",
		"','", "'var'", "'type'", "'new'", "'delete'", "'struct'", "'.'", "'func'",
		"'return'", "'if'", "'else'", "'for'", "'scan'", "'printf'", "", "",
		"", "'nil'", "'int'", "'bool'",
	}
	staticData.symbolicNames = []string{
		"", "ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE",
		"PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"SEMICOLON", "COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT",
		"FUNC", "RET", "IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "STRING_LIT",
		"BOOL_LIT", "NIL_LIT", "INT", "BOOL", "IDENT", "WS", "COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "types", "typeDeclaration", "fields", "morefields", "decl",
		"type", "declarations", "declaration", "ids", "otherids", "functions",
		"function", "parameters", "parametersPrime", "parametersPPrime", "returnType",
		"statements", "statement", "read", "block", "delete", "assignment",
		"print", "printPrime", "conditional", "conditionalPrime", "loop", "returnRule",
		"invocation", "arguments", "argumentsPrime", "argumentsPrimePrime",
		"lValue", "lValuePrime", "expression", "boolTerm", "equalTerm", "relationTerm",
		"simpleTerm", "term", "unaryTerm", "selectorTerm", "selectorTermPrime",
		"subfactor", "functioncall", "allocation", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 378, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 103, 8, 1, 10, 1, 12,
		1, 106, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 5, 3, 119, 8, 3, 10, 3, 12, 3, 122, 9, 3, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 134, 8, 6, 1, 7, 5, 7, 137,
		8, 7, 10, 7, 12, 7, 140, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		5, 9, 149, 8, 9, 10, 9, 12, 9, 152, 9, 9, 1, 10, 1, 10, 1, 10, 1, 11, 5,
		11, 158, 8, 11, 10, 11, 12, 11, 161, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 167, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3,
		13, 176, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 182, 8, 14, 10, 14,
		12, 14, 185, 9, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 5, 17, 193,
		8, 17, 10, 17, 12, 17, 196, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 3, 18, 207, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 230, 8, 23, 10, 23,
		12, 23, 233, 9, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 247, 8, 25, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 3, 28, 260, 8,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 270,
		8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 276, 8, 31, 10, 31, 12, 31, 279,
		9, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 286, 8, 33, 10, 33, 12,
		33, 289, 9, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 5, 35, 297, 8,
		35, 10, 35, 12, 35, 300, 9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 305, 8, 36,
		10, 36, 12, 36, 308, 9, 36, 1, 37, 1, 37, 1, 37, 5, 37, 313, 8, 37, 10,
		37, 12, 37, 316, 9, 37, 1, 38, 1, 38, 1, 38, 5, 38, 321, 8, 38, 10, 38,
		12, 38, 324, 9, 38, 1, 39, 1, 39, 1, 39, 5, 39, 329, 8, 39, 10, 39, 12,
		39, 332, 9, 39, 1, 40, 1, 40, 1, 40, 5, 40, 337, 8, 40, 10, 40, 12, 40,
		340, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 347, 8, 41, 1, 42,
		1, 42, 5, 42, 351, 8, 42, 10, 42, 12, 42, 354, 9, 42, 1, 43, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 3, 45, 365, 8, 45, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 376, 8,
		47, 1, 47, 0, 0, 48, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 0, 4, 1,
		0, 9, 10, 1, 0, 5, 8, 1, 0, 11, 12, 1, 0, 13, 14, 369, 0, 96, 1, 0, 0,
		0, 2, 104, 1, 0, 0, 0, 4, 107, 1, 0, 0, 0, 6, 115, 1, 0, 0, 0, 8, 123,
		1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 133, 1, 0, 0, 0, 14, 138, 1, 0, 0,
		0, 16, 141, 1, 0, 0, 0, 18, 146, 1, 0, 0, 0, 20, 153, 1, 0, 0, 0, 22, 159,
		1, 0, 0, 0, 24, 162, 1, 0, 0, 0, 26, 173, 1, 0, 0, 0, 28, 179, 1, 0, 0,
		0, 30, 186, 1, 0, 0, 0, 32, 189, 1, 0, 0, 0, 34, 194, 1, 0, 0, 0, 36, 206,
		1, 0, 0, 0, 38, 208, 1, 0, 0, 0, 40, 212, 1, 0, 0, 0, 42, 216, 1, 0, 0,
		0, 44, 220, 1, 0, 0, 0, 46, 225, 1, 0, 0, 0, 48, 237, 1, 0, 0, 0, 50, 240,
		1, 0, 0, 0, 52, 248, 1, 0, 0, 0, 54, 251, 1, 0, 0, 0, 56, 257, 1, 0, 0,
		0, 58, 263, 1, 0, 0, 0, 60, 267, 1, 0, 0, 0, 62, 273, 1, 0, 0, 0, 64, 280,
		1, 0, 0, 0, 66, 283, 1, 0, 0, 0, 68, 290, 1, 0, 0, 0, 70, 293, 1, 0, 0,
		0, 72, 301, 1, 0, 0, 0, 74, 309, 1, 0, 0, 0, 76, 317, 1, 0, 0, 0, 78, 325,
		1, 0, 0, 0, 80, 333, 1, 0, 0, 0, 82, 346, 1, 0, 0, 0, 84, 348, 1, 0, 0,
		0, 86, 355, 1, 0, 0, 0, 88, 358, 1, 0, 0, 0, 90, 362, 1, 0, 0, 0, 92, 366,
		1, 0, 0, 0, 94, 375, 1, 0, 0, 0, 96, 97, 3, 2, 1, 0, 97, 98, 3, 14, 7,
		0, 98, 99, 3, 22, 11, 0, 99, 100, 5, 0, 0, 1, 100, 1, 1, 0, 0, 0, 101,
		103, 3, 4, 2, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 3, 1, 0, 0, 0, 106, 104, 1, 0, 0,
		0, 107, 108, 5, 22, 0, 0, 108, 109, 5, 40, 0, 0, 109, 110, 5, 25, 0, 0,
		110, 111, 5, 17, 0, 0, 111, 112, 3, 6, 3, 0, 112, 113, 5, 18, 0, 0, 113,
		114, 5, 19, 0, 0, 114, 5, 1, 0, 0, 0, 115, 116, 3, 10, 5, 0, 116, 120,
		5, 19, 0, 0, 117, 119, 3, 8, 4, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0,
		0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 7, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 123, 124, 3, 10, 5, 0, 124, 125, 5, 19, 0, 0, 125, 9,
		1, 0, 0, 0, 126, 127, 5, 40, 0, 0, 127, 128, 3, 12, 6, 0, 128, 11, 1, 0,
		0, 0, 129, 134, 5, 38, 0, 0, 130, 134, 5, 39, 0, 0, 131, 132, 5, 13, 0,
		0, 132, 134, 5, 40, 0, 0, 133, 129, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 134, 13, 1, 0, 0, 0, 135, 137, 3, 16, 8, 0, 136, 135,
		1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0,
		0, 0, 139, 15, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 21, 0, 0,
		142, 143, 3, 18, 9, 0, 143, 144, 3, 12, 6, 0, 144, 145, 5, 19, 0, 0, 145,
		17, 1, 0, 0, 0, 146, 150, 5, 40, 0, 0, 147, 149, 3, 20, 10, 0, 148, 147,
		1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 19, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 20, 0, 0,
		154, 155, 5, 40, 0, 0, 155, 21, 1, 0, 0, 0, 156, 158, 3, 24, 12, 0, 157,
		156, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 23, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163, 5, 27,
		0, 0, 163, 164, 5, 40, 0, 0, 164, 166, 3, 26, 13, 0, 165, 167, 3, 32, 16,
		0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		169, 5, 17, 0, 0, 169, 170, 3, 14, 7, 0, 170, 171, 3, 34, 17, 0, 171, 172,
		5, 18, 0, 0, 172, 25, 1, 0, 0, 0, 173, 175, 5, 15, 0, 0, 174, 176, 3, 28,
		14, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0,
		177, 178, 5, 16, 0, 0, 178, 27, 1, 0, 0, 0, 179, 183, 3, 10, 5, 0, 180,
		182, 3, 30, 15, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 29, 1, 0, 0, 0, 185, 183, 1, 0,
		0, 0, 186, 187, 5, 20, 0, 0, 187, 188, 3, 10, 5, 0, 188, 31, 1, 0, 0, 0,
		189, 190, 3, 12, 6, 0, 190, 33, 1, 0, 0, 0, 191, 193, 3, 36, 18, 0, 192,
		191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 35, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 207, 3, 40,
		20, 0, 198, 207, 3, 44, 22, 0, 199, 207, 3, 46, 23, 0, 200, 207, 3, 42,
		21, 0, 201, 207, 3, 38, 19, 0, 202, 207, 3, 50, 25, 0, 203, 207, 3, 54,
		27, 0, 204, 207, 3, 56, 28, 0, 205, 207, 3, 58, 29, 0, 206, 197, 1, 0,
		0, 0, 206, 198, 1, 0, 0, 0, 206, 199, 1, 0, 0, 0, 206, 200, 1, 0, 0, 0,
		206, 201, 1, 0, 0, 0, 206, 202, 1, 0, 0, 0, 206, 203, 1, 0, 0, 0, 206,
		204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 37, 1, 0, 0, 0, 208, 209, 5,
		32, 0, 0, 209, 210, 3, 66, 33, 0, 210, 211, 5, 19, 0, 0, 211, 39, 1, 0,
		0, 0, 212, 213, 5, 17, 0, 0, 213, 214, 3, 34, 17, 0, 214, 215, 5, 18, 0,
		0, 215, 41, 1, 0, 0, 0, 216, 217, 5, 24, 0, 0, 217, 218, 3, 70, 35, 0,
		218, 219, 5, 19, 0, 0, 219, 43, 1, 0, 0, 0, 220, 221, 3, 66, 33, 0, 221,
		222, 5, 1, 0, 0, 222, 223, 3, 70, 35, 0, 223, 224, 5, 19, 0, 0, 224, 45,
		1, 0, 0, 0, 225, 226, 5, 33, 0, 0, 226, 227, 5, 15, 0, 0, 227, 231, 5,
		35, 0, 0, 228, 230, 3, 48, 24, 0, 229, 228, 1, 0, 0, 0, 230, 233, 1, 0,
		0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0,
		233, 231, 1, 0, 0, 0, 234, 235, 5, 16, 0, 0, 235, 236, 5, 19, 0, 0, 236,
		47, 1, 0, 0, 0, 237, 238, 5, 20, 0, 0, 238, 239, 3, 70, 35, 0, 239, 49,
		1, 0, 0, 0, 240, 241, 5, 29, 0, 0, 241, 242, 5, 15, 0, 0, 242, 243, 3,
		70, 35, 0, 243, 244, 5, 16, 0, 0, 244, 246, 3, 40, 20, 0, 245, 247, 3,
		52, 26, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 51, 1, 0, 0,
		0, 248, 249, 5, 30, 0, 0, 249, 250, 3, 40, 20, 0, 250, 53, 1, 0, 0, 0,
		251, 252, 5, 31, 0, 0, 252, 253, 5, 15, 0, 0, 253, 254, 3, 70, 35, 0, 254,
		255, 5, 16, 0, 0, 255, 256, 3, 40, 20, 0, 256, 55, 1, 0, 0, 0, 257, 259,
		5, 28, 0, 0, 258, 260, 3, 70, 35, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1,
		0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 5, 19, 0, 0, 262, 57, 1, 0, 0,
		0, 263, 264, 5, 40, 0, 0, 264, 265, 3, 60, 30, 0, 265, 266, 5, 19, 0, 0,
		266, 59, 1, 0, 0, 0, 267, 269, 5, 15, 0, 0, 268, 270, 3, 62, 31, 0, 269,
		268, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272,
		5, 16, 0, 0, 272, 61, 1, 0, 0, 0, 273, 277, 3, 70, 35, 0, 274, 276, 3,
		64, 32, 0, 275, 274, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 63, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0,
		280, 281, 5, 20, 0, 0, 281, 282, 3, 70, 35, 0, 282, 65, 1, 0, 0, 0, 283,
		287, 5, 40, 0, 0, 284, 286, 3, 68, 34, 0, 285, 284, 1, 0, 0, 0, 286, 289,
		1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 67, 1, 0,
		0, 0, 289, 287, 1, 0, 0, 0, 290, 291, 5, 26, 0, 0, 291, 292, 5, 40, 0,
		0, 292, 69, 1, 0, 0, 0, 293, 298, 3, 72, 36, 0, 294, 295, 5, 3, 0, 0, 295,
		297, 3, 72, 36, 0, 296, 294, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296,
		1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 71, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 301, 306, 3, 74, 37, 0, 302, 303, 5, 2, 0, 0, 303, 305, 3, 74, 37,
		0, 304, 302, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306,
		307, 1, 0, 0, 0, 307, 73, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 314, 3,
		76, 38, 0, 310, 311, 7, 0, 0, 0, 311, 313, 3, 76, 38, 0, 312, 310, 1, 0,
		0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0,
		315, 75, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 322, 3, 78, 39, 0, 318,
		319, 7, 1, 0, 0, 319, 321, 3, 78, 39, 0, 320, 318, 1, 0, 0, 0, 321, 324,
		1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 77, 1, 0,
		0, 0, 324, 322, 1, 0, 0, 0, 325, 330, 3, 80, 40, 0, 326, 327, 7, 2, 0,
		0, 327, 329, 3, 80, 40, 0, 328, 326, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0,
		330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 79, 1, 0, 0, 0, 332, 330,
		1, 0, 0, 0, 333, 338, 3, 82, 41, 0, 334, 335, 7, 3, 0, 0, 335, 337, 3,
		82, 41, 0, 336, 334, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1, 0,
		0, 0, 338, 339, 1, 0, 0, 0, 339, 81, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0,
		341, 342, 5, 4, 0, 0, 342, 347, 3, 84, 42, 0, 343, 344, 5, 12, 0, 0, 344,
		347, 3, 84, 42, 0, 345, 347, 3, 84, 42, 0, 346, 341, 1, 0, 0, 0, 346, 343,
		1, 0, 0, 0, 346, 345, 1, 0, 0, 0, 347, 83, 1, 0, 0, 0, 348, 352, 3, 94,
		47, 0, 349, 351, 3, 86, 43, 0, 350, 349, 1, 0, 0, 0, 351, 354, 1, 0, 0,
		0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 85, 1, 0, 0, 0, 354,
		352, 1, 0, 0, 0, 355, 356, 5, 26, 0, 0, 356, 357, 5, 40, 0, 0, 357, 87,
		1, 0, 0, 0, 358, 359, 5, 15, 0, 0, 359, 360, 3, 70, 35, 0, 360, 361, 5,
		16, 0, 0, 361, 89, 1, 0, 0, 0, 362, 364, 5, 40, 0, 0, 363, 365, 3, 60,
		30, 0, 364, 363, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 91, 1, 0, 0, 0,
		366, 367, 5, 23, 0, 0, 367, 368, 5, 40, 0, 0, 368, 93, 1, 0, 0, 0, 369,
		376, 3, 88, 44, 0, 370, 376, 3, 90, 45, 0, 371, 376, 5, 34, 0, 0, 372,
		376, 3, 92, 46, 0, 373, 376, 5, 36, 0, 0, 374, 376, 5, 37, 0, 0, 375, 369,
		1, 0, 0, 0, 375, 370, 1, 0, 0, 0, 375, 371, 1, 0, 0, 0, 375, 372, 1, 0,
		0, 0, 375, 373, 1, 0, 0, 0, 375, 374, 1, 0, 0, 0, 376, 95, 1, 0, 0, 0,
		27, 104, 120, 133, 138, 150, 159, 166, 175, 183, 194, 206, 231, 246, 259,
		269, 277, 287, 298, 306, 314, 322, 330, 338, 346, 352, 364, 375,
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

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &goliteparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF        = antlr.TokenEOF
	GoliteParserASSIGN     = 1
	GoliteParserAND        = 2
	GoliteParserOR         = 3
	GoliteParserNOT        = 4
	GoliteParserLT         = 5
	GoliteParserGT         = 6
	GoliteParserLE         = 7
	GoliteParserGE         = 8
	GoliteParserEQ         = 9
	GoliteParserNE         = 10
	GoliteParserPLUS       = 11
	GoliteParserMINUS      = 12
	GoliteParserMULT       = 13
	GoliteParserDIV        = 14
	GoliteParserLPAREN     = 15
	GoliteParserRPAREN     = 16
	GoliteParserLBRACE     = 17
	GoliteParserRBRACE     = 18
	GoliteParserSEMICOLON  = 19
	GoliteParserCOMMA      = 20
	GoliteParserVAR        = 21
	GoliteParserTYPE       = 22
	GoliteParserNEW        = 23
	GoliteParserDELETE     = 24
	GoliteParserSTRUCT     = 25
	GoliteParserDOT        = 26
	GoliteParserFUNC       = 27
	GoliteParserRET        = 28
	GoliteParserIF         = 29
	GoliteParserELSE       = 30
	GoliteParserFOR        = 31
	GoliteParserSCAN       = 32
	GoliteParserPRINTF     = 33
	GoliteParserINT_LIT    = 34
	GoliteParserSTRING_LIT = 35
	GoliteParserBOOL_LIT   = 36
	GoliteParserNIL_LIT    = 37
	GoliteParserINT        = 38
	GoliteParserBOOL       = 39
	GoliteParserIDENT      = 40
	GoliteParserWS         = 41
	GoliteParserCOMMENT    = 42
)

// GoliteParser rules.
const (
	GoliteParserRULE_program             = 0
	GoliteParserRULE_types               = 1
	GoliteParserRULE_typeDeclaration     = 2
	GoliteParserRULE_fields              = 3
	GoliteParserRULE_morefields          = 4
	GoliteParserRULE_decl                = 5
	GoliteParserRULE_type                = 6
	GoliteParserRULE_declarations        = 7
	GoliteParserRULE_declaration         = 8
	GoliteParserRULE_ids                 = 9
	GoliteParserRULE_otherids            = 10
	GoliteParserRULE_functions           = 11
	GoliteParserRULE_function            = 12
	GoliteParserRULE_parameters          = 13
	GoliteParserRULE_parametersPrime     = 14
	GoliteParserRULE_parametersPPrime    = 15
	GoliteParserRULE_returnType          = 16
	GoliteParserRULE_statements          = 17
	GoliteParserRULE_statement           = 18
	GoliteParserRULE_read                = 19
	GoliteParserRULE_block               = 20
	GoliteParserRULE_delete              = 21
	GoliteParserRULE_assignment          = 22
	GoliteParserRULE_print               = 23
	GoliteParserRULE_printPrime          = 24
	GoliteParserRULE_conditional         = 25
	GoliteParserRULE_conditionalPrime    = 26
	GoliteParserRULE_loop                = 27
	GoliteParserRULE_returnRule          = 28
	GoliteParserRULE_invocation          = 29
	GoliteParserRULE_arguments           = 30
	GoliteParserRULE_argumentsPrime      = 31
	GoliteParserRULE_argumentsPrimePrime = 32
	GoliteParserRULE_lValue              = 33
	GoliteParserRULE_lValuePrime         = 34
	GoliteParserRULE_expression          = 35
	GoliteParserRULE_boolTerm            = 36
	GoliteParserRULE_equalTerm           = 37
	GoliteParserRULE_relationTerm        = 38
	GoliteParserRULE_simpleTerm          = 39
	GoliteParserRULE_term                = 40
	GoliteParserRULE_unaryTerm           = 41
	GoliteParserRULE_selectorTerm        = 42
	GoliteParserRULE_selectorTermPrime   = 43
	GoliteParserRULE_subfactor           = 44
	GoliteParserRULE_functioncall        = 45
	GoliteParserRULE_allocation          = 46
	GoliteParserRULE_factor              = 47
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() ITypesContext

	// GetDcls returns the dcls rule contexts.
	GetDcls() IDeclarationsContext

	// GetFns returns the fns rule contexts.
	GetFns() IFunctionsContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypesContext)

	// SetDcls sets the dcls rule contexts.
	SetDcls(IDeclarationsContext)

	// SetFns sets the fns rule contexts.
	SetFns(IFunctionsContext)

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     ITypesContext
	dcls   IDeclarationsContext
	fns    IFunctionsContext
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) GetTy() ITypesContext { return s.ty }

func (s *ProgramContext) GetDcls() IDeclarationsContext { return s.dcls }

func (s *ProgramContext) GetFns() IFunctionsContext { return s.fns }

func (s *ProgramContext) SetTy(v ITypesContext) { s.ty = v }

func (s *ProgramContext) SetDcls(v IDeclarationsContext) { s.dcls = v }

func (s *ProgramContext) SetFns(v IFunctionsContext) { s.fns = v }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)

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
		p.SetState(96)

		var _x = p.Types()

		localctx.(*ProgramContext).ty = _x
	}
	{
		p.SetState(97)

		var _x = p.Declarations()

		localctx.(*ProgramContext).dcls = _x
	}
	{
		p.SetState(98)

		var _x = p.Functions()

		localctx.(*ProgramContext).fns = _x
	}
	{
		p.SetState(99)
		p.Match(GoliteParserEOF)
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypesContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
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

	return t.(ITypeDeclarationContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *GoliteParser) Types() (localctx ITypesContext) {
	this := p
	_ = this

	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_types)
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
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(101)
			p.TypeDeclaration()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetFlds returns the flds rule contexts.
	GetFlds() IFieldsContext

	// SetFlds sets the flds rule contexts.
	SetFlds(IFieldsContext)

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	flds   IFieldsContext
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) GetId() antlr.Token { return s.id }

func (s *TypeDeclarationContext) SetId(v antlr.Token) { s.id = v }

func (s *TypeDeclarationContext) GetFlds() IFieldsContext { return s.flds }

func (s *TypeDeclarationContext) SetFlds(v IFieldsContext) { s.flds = v }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypeDeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypeDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *TypeDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *TypeDeclarationContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *GoliteParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	this := p
	_ = this

	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_typeDeclaration)

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
		p.SetState(107)
		p.Match(GoliteParserTYPE)
	}
	{
		p.SetState(108)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*TypeDeclarationContext).id = _m
	}
	{
		p.SetState(109)
		p.Match(GoliteParserSTRUCT)
	}
	{
		p.SetState(110)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(111)

		var _x = p.Fields()

		localctx.(*TypeDeclarationContext).flds = _x
	}
	{
		p.SetState(112)
		p.Match(GoliteParserRBRACE)
	}
	{
		p.SetState(113)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDcl returns the dcl rule contexts.
	GetDcl() IDeclContext

	// SetDcl sets the dcl rule contexts.
	SetDcl(IDeclContext)

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dcl    IDeclContext
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) GetDcl() IDeclContext { return s.dcl }

func (s *FieldsContext) SetDcl(v IDeclContext) { s.dcl = v }

func (s *FieldsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsContext) AllMorefields() []IMorefieldsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMorefieldsContext); ok {
			len++
		}
	}

	tst := make([]IMorefieldsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMorefieldsContext); ok {
			tst[i] = t.(IMorefieldsContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Morefields(i int) IMorefieldsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMorefieldsContext); ok {
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

	return t.(IMorefieldsContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fields)
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
		p.SetState(115)

		var _x = p.Decl()

		localctx.(*FieldsContext).dcl = _x
	}
	{
		p.SetState(116)
		p.Match(GoliteParserSEMICOLON)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserIDENT {
		{
			p.SetState(117)
			p.Morefields()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMorefieldsContext is an interface to support dynamic dispatch.
type IMorefieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDcl returns the dcl rule contexts.
	GetDcl() IDeclContext

	// SetDcl sets the dcl rule contexts.
	SetDcl(IDeclContext)

	// IsMorefieldsContext differentiates from other interfaces.
	IsMorefieldsContext()
}

type MorefieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dcl    IDeclContext
}

func NewEmptyMorefieldsContext() *MorefieldsContext {
	var p = new(MorefieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_morefields
	return p
}

func (*MorefieldsContext) IsMorefieldsContext() {}

func NewMorefieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MorefieldsContext {
	var p = new(MorefieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_morefields

	return p
}

func (s *MorefieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *MorefieldsContext) GetDcl() IDeclContext { return s.dcl }

func (s *MorefieldsContext) SetDcl(v IDeclContext) { s.dcl = v }

func (s *MorefieldsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *MorefieldsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *MorefieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MorefieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MorefieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterMorefields(s)
	}
}

func (s *MorefieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitMorefields(s)
	}
}

func (p *GoliteParser) Morefields() (localctx IMorefieldsContext) {
	this := p
	_ = this

	localctx = NewMorefieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_morefields)

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
		p.SetState(123)

		var _x = p.Decl()

		localctx.(*MorefieldsContext).dcl = _x
	}
	{
		p.SetState(124)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	ty     ITypeContext
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) GetId() antlr.Token { return s.id }

func (s *DeclContext) SetId(v antlr.Token) { s.id = v }

func (s *DeclContext) GetTy() ITypeContext { return s.ty }

func (s *DeclContext) SetTy(v ITypeContext) { s.ty = v }

func (s *DeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_decl)

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

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*DeclContext).id = _m
	}
	{
		p.SetState(127)

		var _x = p.Type_()

		localctx.(*DeclContext).ty = _x
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetId() antlr.Token { return s.id }

func (s *TypeContext) SetId(v antlr.Token) { s.id = v }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL, 0)
}

func (s *TypeContext) MULT() antlr.TerminalNode {
	return s.GetToken(GoliteParserMULT, 0)
}

func (s *TypeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_type)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(GoliteParserINT)
		}

	case GoliteParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(GoliteParserBOOL)
		}

	case GoliteParserMULT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(GoliteParserMULT)
		}
		{
			p.SetState(132)

			var _m = p.Match(GoliteParserIDENT)

			localctx.(*TypeContext).id = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
	return p
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (p *GoliteParser) Declarations() (localctx IDeclarationsContext) {
	this := p
	_ = this

	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_declarations)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(135)
			p.Declaration()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdx returns the idx rule contexts.
	GetIdx() IIdsContext

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetIdx sets the idx rule contexts.
	SetIdx(IIdsContext)

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	idx    IIdsContext
	ty     ITypeContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetIdx() IIdsContext { return s.idx }

func (s *DeclarationContext) GetTy() ITypeContext { return s.ty }

func (s *DeclarationContext) SetIdx(v IIdsContext) { s.idx = v }

func (s *DeclarationContext) SetTy(v ITypeContext) { s.ty = v }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) Ids() IIdsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsContext)
}

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_declaration)

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
		p.SetState(141)
		p.Match(GoliteParserVAR)
	}
	{
		p.SetState(142)

		var _x = p.Ids()

		localctx.(*DeclarationContext).idx = _x
	}
	{
		p.SetState(143)

		var _x = p.Type_()

		localctx.(*DeclarationContext).ty = _x
	}
	{
		p.SetState(144)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IIdsContext is an interface to support dynamic dispatch.
type IIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsIdsContext differentiates from other interfaces.
	IsIdsContext()
}

type IdsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyIdsContext() *IdsContext {
	var p = new(IdsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
	return p
}

func (*IdsContext) IsIdsContext() {}

func NewIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsContext {
	var p = new(IdsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ids

	return p
}

func (s *IdsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsContext) GetId() antlr.Token { return s.id }

func (s *IdsContext) SetId(v antlr.Token) { s.id = v }

func (s *IdsContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *IdsContext) AllOtherids() []IOtheridsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOtheridsContext); ok {
			len++
		}
	}

	tst := make([]IOtheridsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOtheridsContext); ok {
			tst[i] = t.(IOtheridsContext)
			i++
		}
	}

	return tst
}

func (s *IdsContext) Otherids(i int) IOtheridsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOtheridsContext); ok {
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

	return t.(IOtheridsContext)
}

func (s *IdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIds(s)
	}
}

func (s *IdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIds(s)
	}
}

func (p *GoliteParser) Ids() (localctx IIdsContext) {
	this := p
	_ = this

	localctx = NewIdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_ids)
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
		p.SetState(146)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*IdsContext).id = _m
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(147)
			p.Otherids()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOtheridsContext is an interface to support dynamic dispatch.
type IOtheridsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsOtheridsContext differentiates from other interfaces.
	IsOtheridsContext()
}

type OtheridsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyOtheridsContext() *OtheridsContext {
	var p = new(OtheridsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_otherids
	return p
}

func (*OtheridsContext) IsOtheridsContext() {}

func NewOtheridsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OtheridsContext {
	var p = new(OtheridsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_otherids

	return p
}

func (s *OtheridsContext) GetParser() antlr.Parser { return s.parser }

func (s *OtheridsContext) GetId() antlr.Token { return s.id }

func (s *OtheridsContext) SetId(v antlr.Token) { s.id = v }

func (s *OtheridsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *OtheridsContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *OtheridsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtheridsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OtheridsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterOtherids(s)
	}
}

func (s *OtheridsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitOtherids(s)
	}
}

func (p *GoliteParser) Otherids() (localctx IOtheridsContext) {
	this := p
	_ = this

	localctx = NewOtheridsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_otherids)

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
		p.SetState(153)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(154)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*OtheridsContext).id = _m
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionsContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
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

	return t.(IFunctionContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *GoliteParser) Functions() (localctx IFunctionsContext) {
	this := p
	_ = this

	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_functions)
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
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNC {
		{
			p.SetState(156)
			p.Function()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IParametersContext

	// GetRty returns the rty rule contexts.
	GetRty() IReturnTypeContext

	// GetDcls returns the dcls rule contexts.
	GetDcls() IDeclarationsContext

	// GetStmts returns the stmts rule contexts.
	GetStmts() IStatementsContext

	// SetParams sets the params rule contexts.
	SetParams(IParametersContext)

	// SetRty sets the rty rule contexts.
	SetRty(IReturnTypeContext)

	// SetDcls sets the dcls rule contexts.
	SetDcls(IDeclarationsContext)

	// SetStmts sets the stmts rule contexts.
	SetStmts(IStatementsContext)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	params IParametersContext
	rty    IReturnTypeContext
	dcls   IDeclarationsContext
	stmts  IStatementsContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetId() antlr.Token { return s.id }

func (s *FunctionContext) SetId(v antlr.Token) { s.id = v }

func (s *FunctionContext) GetParams() IParametersContext { return s.params }

func (s *FunctionContext) GetRty() IReturnTypeContext { return s.rty }

func (s *FunctionContext) GetDcls() IDeclarationsContext { return s.dcls }

func (s *FunctionContext) GetStmts() IStatementsContext { return s.stmts }

func (s *FunctionContext) SetParams(v IParametersContext) { s.params = v }

func (s *FunctionContext) SetRty(v IReturnTypeContext) { s.rty = v }

func (s *FunctionContext) SetDcls(v IDeclarationsContext) { s.dcls = v }

func (s *FunctionContext) SetStmts(v IStatementsContext) { s.stmts = v }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNC, 0)
}

func (s *FunctionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *FunctionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *FunctionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *FunctionContext) Statements() IStatementsContext {
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

func (s *FunctionContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_function)
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
		p.SetState(162)
		p.Match(GoliteParserFUNC)
	}
	{
		p.SetState(163)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*FunctionContext).id = _m
	}
	{
		p.SetState(164)

		var _x = p.Parameters()

		localctx.(*FunctionContext).params = _x
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&824633729024) != 0 {
		{
			p.SetState(165)

			var _x = p.ReturnType()

			localctx.(*FunctionContext).rty = _x
		}

	}
	{
		p.SetState(168)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(169)

		var _x = p.Declarations()

		localctx.(*FunctionContext).dcls = _x
	}
	{
		p.SetState(170)

		var _x = p.Statements()

		localctx.(*FunctionContext).stmts = _x
	}
	{
		p.SetState(171)
		p.Match(GoliteParserRBRACE)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ParametersContext) ParametersPrime() IParametersPrimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersPrimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersPrimeContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GoliteParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_parameters)
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
		p.SetState(173)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserIDENT {
		{
			p.SetState(174)
			p.ParametersPrime()
		}

	}
	{
		p.SetState(177)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IParametersPrimeContext is an interface to support dynamic dispatch.
type IParametersPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDcl returns the dcl rule contexts.
	GetDcl() IDeclContext

	// SetDcl sets the dcl rule contexts.
	SetDcl(IDeclContext)

	// IsParametersPrimeContext differentiates from other interfaces.
	IsParametersPrimeContext()
}

type ParametersPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dcl    IDeclContext
}

func NewEmptyParametersPrimeContext() *ParametersPrimeContext {
	var p = new(ParametersPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parametersPrime
	return p
}

func (*ParametersPrimeContext) IsParametersPrimeContext() {}

func NewParametersPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersPrimeContext {
	var p = new(ParametersPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parametersPrime

	return p
}

func (s *ParametersPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersPrimeContext) GetDcl() IDeclContext { return s.dcl }

func (s *ParametersPrimeContext) SetDcl(v IDeclContext) { s.dcl = v }

func (s *ParametersPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersPrimeContext) AllParametersPPrime() []IParametersPPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametersPPrimeContext); ok {
			len++
		}
	}

	tst := make([]IParametersPPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametersPPrimeContext); ok {
			tst[i] = t.(IParametersPPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ParametersPrimeContext) ParametersPPrime(i int) IParametersPPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersPPrimeContext); ok {
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

	return t.(IParametersPPrimeContext)
}

func (s *ParametersPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParametersPrime(s)
	}
}

func (s *ParametersPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParametersPrime(s)
	}
}

func (p *GoliteParser) ParametersPrime() (localctx IParametersPrimeContext) {
	this := p
	_ = this

	localctx = NewParametersPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_parametersPrime)
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
		p.SetState(179)

		var _x = p.Decl()

		localctx.(*ParametersPrimeContext).dcl = _x
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(180)
			p.ParametersPPrime()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParametersPPrimeContext is an interface to support dynamic dispatch.
type IParametersPPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDcl returns the dcl rule contexts.
	GetDcl() IDeclContext

	// SetDcl sets the dcl rule contexts.
	SetDcl(IDeclContext)

	// IsParametersPPrimeContext differentiates from other interfaces.
	IsParametersPPrimeContext()
}

type ParametersPPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dcl    IDeclContext
}

func NewEmptyParametersPPrimeContext() *ParametersPPrimeContext {
	var p = new(ParametersPPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parametersPPrime
	return p
}

func (*ParametersPPrimeContext) IsParametersPPrimeContext() {}

func NewParametersPPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersPPrimeContext {
	var p = new(ParametersPPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parametersPPrime

	return p
}

func (s *ParametersPPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersPPrimeContext) GetDcl() IDeclContext { return s.dcl }

func (s *ParametersPPrimeContext) SetDcl(v IDeclContext) { s.dcl = v }

func (s *ParametersPPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ParametersPPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersPPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersPPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersPPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParametersPPrime(s)
	}
}

func (s *ParametersPPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParametersPPrime(s)
	}
}

func (p *GoliteParser) ParametersPPrime() (localctx IParametersPPrimeContext) {
	this := p
	_ = this

	localctx = NewParametersPPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_parametersPPrime)

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
		p.SetState(186)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(187)

		var _x = p.Decl()

		localctx.(*ParametersPPrimeContext).dcl = _x
	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     ITypeContext
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) GetTy() ITypeContext { return s.ty }

func (s *ReturnTypeContext) SetTy(v ITypeContext) { s.ty = v }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *GoliteParser) ReturnType() (localctx IReturnTypeContext) {
	this := p
	_ = this

	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_returnType)

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
		p.SetState(189)

		var _x = p.Type_()

		localctx.(*ReturnTypeContext).ty = _x
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statements

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

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GoliteParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_statements)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1115366227968) != 0 {
		{
			p.SetState(191)
			p.Statement()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBl returns the bl rule contexts.
	GetBl() IBlockContext

	// GetAsmt returns the asmt rule contexts.
	GetAsmt() IAssignmentContext

	// GetPrnt returns the prnt rule contexts.
	GetPrnt() IPrintContext

	// GetDel returns the del rule contexts.
	GetDel() IDeleteContext

	// GetRd returns the rd rule contexts.
	GetRd() IReadContext

	// GetCond returns the cond rule contexts.
	GetCond() IConditionalContext

	// GetLp returns the lp rule contexts.
	GetLp() ILoopContext

	// GetRet returns the ret rule contexts.
	GetRet() IReturnRuleContext

	// GetInvoke returns the invoke rule contexts.
	GetInvoke() IInvocationContext

	// SetBl sets the bl rule contexts.
	SetBl(IBlockContext)

	// SetAsmt sets the asmt rule contexts.
	SetAsmt(IAssignmentContext)

	// SetPrnt sets the prnt rule contexts.
	SetPrnt(IPrintContext)

	// SetDel sets the del rule contexts.
	SetDel(IDeleteContext)

	// SetRd sets the rd rule contexts.
	SetRd(IReadContext)

	// SetCond sets the cond rule contexts.
	SetCond(IConditionalContext)

	// SetLp sets the lp rule contexts.
	SetLp(ILoopContext)

	// SetRet sets the ret rule contexts.
	SetRet(IReturnRuleContext)

	// SetInvoke sets the invoke rule contexts.
	SetInvoke(IInvocationContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	bl     IBlockContext
	asmt   IAssignmentContext
	prnt   IPrintContext
	del    IDeleteContext
	rd     IReadContext
	cond   IConditionalContext
	lp     ILoopContext
	ret    IReturnRuleContext
	invoke IInvocationContext
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetBl() IBlockContext { return s.bl }

func (s *StatementContext) GetAsmt() IAssignmentContext { return s.asmt }

func (s *StatementContext) GetPrnt() IPrintContext { return s.prnt }

func (s *StatementContext) GetDel() IDeleteContext { return s.del }

func (s *StatementContext) GetRd() IReadContext { return s.rd }

func (s *StatementContext) GetCond() IConditionalContext { return s.cond }

func (s *StatementContext) GetLp() ILoopContext { return s.lp }

func (s *StatementContext) GetRet() IReturnRuleContext { return s.ret }

func (s *StatementContext) GetInvoke() IInvocationContext { return s.invoke }

func (s *StatementContext) SetBl(v IBlockContext) { s.bl = v }

func (s *StatementContext) SetAsmt(v IAssignmentContext) { s.asmt = v }

func (s *StatementContext) SetPrnt(v IPrintContext) { s.prnt = v }

func (s *StatementContext) SetDel(v IDeleteContext) { s.del = v }

func (s *StatementContext) SetRd(v IReadContext) { s.rd = v }

func (s *StatementContext) SetCond(v IConditionalContext) { s.cond = v }

func (s *StatementContext) SetLp(v ILoopContext) { s.lp = v }

func (s *StatementContext) SetRet(v IReturnRuleContext) { s.ret = v }

func (s *StatementContext) SetInvoke(v IInvocationContext) { s.invoke = v }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) ReturnRule() IReturnRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnRuleContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_statement)

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

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)

			var _x = p.Block()

			localctx.(*StatementContext).bl = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)

			var _x = p.Assignment()

			localctx.(*StatementContext).asmt = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)

			var _x = p.Print_()

			localctx.(*StatementContext).prnt = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(200)

			var _x = p.Delete_()

			localctx.(*StatementContext).del = _x
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)

			var _x = p.Read()

			localctx.(*StatementContext).rd = _x
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(202)

			var _x = p.Conditional()

			localctx.(*StatementContext).cond = _x
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(203)

			var _x = p.Loop()

			localctx.(*StatementContext).lp = _x
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(204)

			var _x = p.ReturnRule()

			localctx.(*StatementContext).ret = _x
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(205)

			var _x = p.Invocation()

			localctx.(*StatementContext).invoke = _x
		}

	}

	return localctx
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLval returns the lval rule contexts.
	GetLval() ILValueContext

	// SetLval sets the lval rule contexts.
	SetLval(ILValueContext)

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lval   ILValueContext
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) GetLval() ILValueContext { return s.lval }

func (s *ReadContext) SetLval(v ILValueContext) { s.lval = v }

func (s *ReadContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *ReadContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReadContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *GoliteParser) Read() (localctx IReadContext) {
	this := p
	_ = this

	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_read)

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
		p.SetState(208)
		p.Match(GoliteParserSCAN)
	}
	{
		p.SetState(209)

		var _x = p.LValue()

		localctx.(*ReadContext).lval = _x
	}
	{
		p.SetState(210)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStmts returns the stmts rule contexts.
	GetStmts() IStatementsContext

	// SetStmts sets the stmts rule contexts.
	SetStmts(IStatementsContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stmts  IStatementsContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) GetStmts() IStatementsContext { return s.stmts }

func (s *BlockContext) SetStmts(v IStatementsContext) { s.stmts = v }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_block)

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
		p.SetState(212)
		p.Match(GoliteParserLBRACE)
	}
	{
		p.SetState(213)

		var _x = p.Statements()

		localctx.(*BlockContext).stmts = _x
	}
	{
		p.SetState(214)
		p.Match(GoliteParserRBRACE)
	}

	return localctx
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) GetExpr() IExpressionContext { return s.expr }

func (s *DeleteContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
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

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	this := p
	_ = this

	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_delete)

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
		p.SetState(216)
		p.Match(GoliteParserDELETE)
	}
	{
		p.SetState(217)

		var _x = p.Expression()

		localctx.(*DeleteContext).expr = _x
	}
	{
		p.SetState(218)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLval returns the lval rule contexts.
	GetLval() ILValueContext

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetLval sets the lval rule contexts.
	SetLval(ILValueContext)

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lval   ILValueContext
	expr   IExpressionContext
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) GetLval() ILValueContext { return s.lval }

func (s *AssignmentContext) GetExpr() IExpressionContext { return s.expr }

func (s *AssignmentContext) SetLval(v ILValueContext) { s.lval = v }

func (s *AssignmentContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoliteParserASSIGN, 0)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
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
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_assignment)

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
		p.SetState(220)

		var _x = p.LValue()

		localctx.(*AssignmentContext).lval = _x
	}
	{
		p.SetState(221)
		p.Match(GoliteParserASSIGN)
	}
	{
		p.SetState(222)

		var _x = p.Expression()

		localctx.(*AssignmentContext).expr = _x
	}
	{
		p.SetState(223)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStr returns the str token.
	GetStr() antlr.Token

	// SetStr sets the str token.
	SetStr(antlr.Token)

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	str    antlr.Token
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) GetStr() antlr.Token { return s.str }

func (s *PrintContext) SetStr(v antlr.Token) { s.str = v }

func (s *PrintContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINTF, 0)
}

func (s *PrintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *PrintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING_LIT, 0)
}

func (s *PrintContext) AllPrintPrime() []IPrintPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintPrimeContext); ok {
			len++
		}
	}

	tst := make([]IPrintPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintPrimeContext); ok {
			tst[i] = t.(IPrintPrimeContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) PrintPrime(i int) IPrintPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintPrimeContext); ok {
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

	return t.(IPrintPrimeContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	this := p
	_ = this

	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_print)
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
		p.SetState(225)
		p.Match(GoliteParserPRINTF)
	}
	{
		p.SetState(226)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(227)

		var _m = p.Match(GoliteParserSTRING_LIT)

		localctx.(*PrintContext).str = _m
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(228)
			p.PrintPrime()
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(235)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IPrintPrimeContext is an interface to support dynamic dispatch.
type IPrintPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsPrintPrimeContext differentiates from other interfaces.
	IsPrintPrimeContext()
}

type PrintPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyPrintPrimeContext() *PrintPrimeContext {
	var p = new(PrintPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_printPrime
	return p
}

func (*PrintPrimeContext) IsPrintPrimeContext() {}

func NewPrintPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintPrimeContext {
	var p = new(PrintPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_printPrime

	return p
}

func (s *PrintPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintPrimeContext) GetExpr() IExpressionContext { return s.expr }

func (s *PrintPrimeContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *PrintPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *PrintPrimeContext) Expression() IExpressionContext {
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

func (s *PrintPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrintPrime(s)
	}
}

func (s *PrintPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrintPrime(s)
	}
}

func (p *GoliteParser) PrintPrime() (localctx IPrintPrimeContext) {
	this := p
	_ = this

	localctx = NewPrintPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_printPrime)

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
		p.SetState(237)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(238)

		var _x = p.Expression()

		localctx.(*PrintPrimeContext).expr = _x
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// GetBl returns the bl rule contexts.
	GetBl() IBlockContext

	// GetThen returns the then rule contexts.
	GetThen() IConditionalPrimeContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// SetBl sets the bl rule contexts.
	SetBl(IBlockContext)

	// SetThen sets the then rule contexts.
	SetThen(IConditionalPrimeContext)

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
	bl     IBlockContext
	then   IConditionalPrimeContext
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) GetExpr() IExpressionContext { return s.expr }

func (s *ConditionalContext) GetBl() IBlockContext { return s.bl }

func (s *ConditionalContext) GetThen() IConditionalPrimeContext { return s.then }

func (s *ConditionalContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *ConditionalContext) SetBl(v IBlockContext) { s.bl = v }

func (s *ConditionalContext) SetThen(v IConditionalPrimeContext) { s.then = v }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ConditionalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
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

func (s *ConditionalContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) ConditionalPrime() IConditionalPrimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalPrimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalPrimeContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_conditional)
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
		p.SetState(240)
		p.Match(GoliteParserIF)
	}
	{
		p.SetState(241)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(242)

		var _x = p.Expression()

		localctx.(*ConditionalContext).expr = _x
	}
	{
		p.SetState(243)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(244)

		var _x = p.Block()

		localctx.(*ConditionalContext).bl = _x
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(245)

			var _x = p.ConditionalPrime()

			localctx.(*ConditionalContext).then = _x
		}

	}

	return localctx
}

// IConditionalPrimeContext is an interface to support dynamic dispatch.
type IConditionalPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBl returns the bl rule contexts.
	GetBl() IBlockContext

	// SetBl sets the bl rule contexts.
	SetBl(IBlockContext)

	// IsConditionalPrimeContext differentiates from other interfaces.
	IsConditionalPrimeContext()
}

type ConditionalPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	bl     IBlockContext
}

func NewEmptyConditionalPrimeContext() *ConditionalPrimeContext {
	var p = new(ConditionalPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditionalPrime
	return p
}

func (*ConditionalPrimeContext) IsConditionalPrimeContext() {}

func NewConditionalPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalPrimeContext {
	var p = new(ConditionalPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditionalPrime

	return p
}

func (s *ConditionalPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalPrimeContext) GetBl() IBlockContext { return s.bl }

func (s *ConditionalPrimeContext) SetBl(v IBlockContext) { s.bl = v }

func (s *ConditionalPrimeContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalPrimeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditionalPrime(s)
	}
}

func (s *ConditionalPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditionalPrime(s)
	}
}

func (p *GoliteParser) ConditionalPrime() (localctx IConditionalPrimeContext) {
	this := p
	_ = this

	localctx = NewConditionalPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_conditionalPrime)

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
		p.SetState(248)
		p.Match(GoliteParserELSE)
	}
	{
		p.SetState(249)

		var _x = p.Block()

		localctx.(*ConditionalPrimeContext).bl = _x
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// GetBl returns the bl rule contexts.
	GetBl() IBlockContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// SetBl sets the bl rule contexts.
	SetBl(IBlockContext)

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
	bl     IBlockContext
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) GetExpr() IExpressionContext { return s.expr }

func (s *LoopContext) GetBl() IBlockContext { return s.bl }

func (s *LoopContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *LoopContext) SetBl(v IBlockContext) { s.bl = v }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *LoopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
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

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_loop)

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
		p.SetState(251)
		p.Match(GoliteParserFOR)
	}
	{
		p.SetState(252)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(253)

		var _x = p.Expression()

		localctx.(*LoopContext).expr = _x
	}
	{
		p.SetState(254)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(255)

		var _x = p.Block()

		localctx.(*LoopContext).bl = _x
	}

	return localctx
}

// IReturnRuleContext is an interface to support dynamic dispatch.
type IReturnRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsReturnRuleContext differentiates from other interfaces.
	IsReturnRuleContext()
}

type ReturnRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyReturnRuleContext() *ReturnRuleContext {
	var p = new(ReturnRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_returnRule
	return p
}

func (*ReturnRuleContext) IsReturnRuleContext() {}

func NewReturnRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnRuleContext {
	var p = new(ReturnRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnRule

	return p
}

func (s *ReturnRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnRuleContext) GetExpr() IExpressionContext { return s.expr }

func (s *ReturnRuleContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *ReturnRuleContext) RET() antlr.TerminalNode {
	return s.GetToken(GoliteParserRET, 0)
}

func (s *ReturnRuleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReturnRuleContext) Expression() IExpressionContext {
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

func (s *ReturnRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnRule(s)
	}
}

func (s *ReturnRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnRule(s)
	}
}

func (p *GoliteParser) ReturnRule() (localctx IReturnRuleContext) {
	this := p
	_ = this

	localctx = NewReturnRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_returnRule)
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
		p.SetState(257)
		p.Match(GoliteParserRET)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1322858352656) != 0 {
		{
			p.SetState(258)

			var _x = p.Expression()

			localctx.(*ReturnRuleContext).expr = _x
		}

	}
	{
		p.SetState(261)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetArgs returns the args rule contexts.
	GetArgs() IArgumentsContext

	// SetArgs sets the args rule contexts.
	SetArgs(IArgumentsContext)

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	args   IArgumentsContext
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) GetId() antlr.Token { return s.id }

func (s *InvocationContext) SetId(v antlr.Token) { s.id = v }

func (s *InvocationContext) GetArgs() IArgumentsContext { return s.args }

func (s *InvocationContext) SetArgs(v IArgumentsContext) { s.args = v }

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	this := p
	_ = this

	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_invocation)

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
		p.SetState(263)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*InvocationContext).id = _m
	}
	{
		p.SetState(264)

		var _x = p.Arguments()

		localctx.(*InvocationContext).args = _x
	}
	{
		p.SetState(265)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgs returns the args rule contexts.
	GetArgs() IArgumentsPrimeContext

	// SetArgs sets the args rule contexts.
	SetArgs(IArgumentsPrimeContext)

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	args   IArgumentsPrimeContext
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) GetArgs() IArgumentsPrimeContext { return s.args }

func (s *ArgumentsContext) SetArgs(v IArgumentsPrimeContext) { s.args = v }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ArgumentsContext) ArgumentsPrime() IArgumentsPrimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsPrimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsPrimeContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GoliteParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_arguments)
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
		p.SetState(267)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1322858352656) != 0 {
		{
			p.SetState(268)

			var _x = p.ArgumentsPrime()

			localctx.(*ArgumentsContext).args = _x
		}

	}
	{
		p.SetState(271)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IArgumentsPrimeContext is an interface to support dynamic dispatch.
type IArgumentsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsArgumentsPrimeContext differentiates from other interfaces.
	IsArgumentsPrimeContext()
}

type ArgumentsPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyArgumentsPrimeContext() *ArgumentsPrimeContext {
	var p = new(ArgumentsPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_argumentsPrime
	return p
}

func (*ArgumentsPrimeContext) IsArgumentsPrimeContext() {}

func NewArgumentsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsPrimeContext {
	var p = new(ArgumentsPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_argumentsPrime

	return p
}

func (s *ArgumentsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsPrimeContext) GetExpr() IExpressionContext { return s.expr }

func (s *ArgumentsPrimeContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *ArgumentsPrimeContext) Expression() IExpressionContext {
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

func (s *ArgumentsPrimeContext) AllArgumentsPrimePrime() []IArgumentsPrimePrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentsPrimePrimeContext); ok {
			len++
		}
	}

	tst := make([]IArgumentsPrimePrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentsPrimePrimeContext); ok {
			tst[i] = t.(IArgumentsPrimePrimeContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsPrimeContext) ArgumentsPrimePrime(i int) IArgumentsPrimePrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsPrimePrimeContext); ok {
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

	return t.(IArgumentsPrimePrimeContext)
}

func (s *ArgumentsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgumentsPrime(s)
	}
}

func (s *ArgumentsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgumentsPrime(s)
	}
}

func (p *GoliteParser) ArgumentsPrime() (localctx IArgumentsPrimeContext) {
	this := p
	_ = this

	localctx = NewArgumentsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_argumentsPrime)
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
		p.SetState(273)

		var _x = p.Expression()

		localctx.(*ArgumentsPrimeContext).expr = _x
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(274)
			p.ArgumentsPrimePrime()
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgumentsPrimePrimeContext is an interface to support dynamic dispatch.
type IArgumentsPrimePrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsArgumentsPrimePrimeContext differentiates from other interfaces.
	IsArgumentsPrimePrimeContext()
}

type ArgumentsPrimePrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyArgumentsPrimePrimeContext() *ArgumentsPrimePrimeContext {
	var p = new(ArgumentsPrimePrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_argumentsPrimePrime
	return p
}

func (*ArgumentsPrimePrimeContext) IsArgumentsPrimePrimeContext() {}

func NewArgumentsPrimePrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsPrimePrimeContext {
	var p = new(ArgumentsPrimePrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_argumentsPrimePrime

	return p
}

func (s *ArgumentsPrimePrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsPrimePrimeContext) GetExpr() IExpressionContext { return s.expr }

func (s *ArgumentsPrimePrimeContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *ArgumentsPrimePrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ArgumentsPrimePrimeContext) Expression() IExpressionContext {
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

func (s *ArgumentsPrimePrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsPrimePrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsPrimePrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgumentsPrimePrime(s)
	}
}

func (s *ArgumentsPrimePrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgumentsPrimePrime(s)
	}
}

func (p *GoliteParser) ArgumentsPrimePrime() (localctx IArgumentsPrimePrimeContext) {
	this := p
	_ = this

	localctx = NewArgumentsPrimePrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_argumentsPrimePrime)

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
		p.SetState(280)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(281)

		var _x = p.Expression()

		localctx.(*ArgumentsPrimePrimeContext).expr = _x
	}

	return localctx
}

// ILValueContext is an interface to support dynamic dispatch.
type ILValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsLValueContext differentiates from other interfaces.
	IsLValueContext()
}

type LValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyLValueContext() *LValueContext {
	var p = new(LValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_lValue
	return p
}

func (*LValueContext) IsLValueContext() {}

func NewLValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValueContext {
	var p = new(LValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lValue

	return p
}

func (s *LValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LValueContext) GetId() antlr.Token { return s.id }

func (s *LValueContext) SetId(v antlr.Token) { s.id = v }

func (s *LValueContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *LValueContext) AllLValuePrime() []ILValuePrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILValuePrimeContext); ok {
			len++
		}
	}

	tst := make([]ILValuePrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILValuePrimeContext); ok {
			tst[i] = t.(ILValuePrimeContext)
			i++
		}
	}

	return tst
}

func (s *LValueContext) LValuePrime(i int) ILValuePrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValuePrimeContext); ok {
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

	return t.(ILValuePrimeContext)
}

func (s *LValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLValue(s)
	}
}

func (s *LValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLValue(s)
	}
}

func (p *GoliteParser) LValue() (localctx ILValueContext) {
	this := p
	_ = this

	localctx = NewLValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_lValue)
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
		p.SetState(283)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*LValueContext).id = _m
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(284)
			p.LValuePrime()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILValuePrimeContext is an interface to support dynamic dispatch.
type ILValuePrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsLValuePrimeContext differentiates from other interfaces.
	IsLValuePrimeContext()
}

type LValuePrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyLValuePrimeContext() *LValuePrimeContext {
	var p = new(LValuePrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_lValuePrime
	return p
}

func (*LValuePrimeContext) IsLValuePrimeContext() {}

func NewLValuePrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValuePrimeContext {
	var p = new(LValuePrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lValuePrime

	return p
}

func (s *LValuePrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *LValuePrimeContext) GetId() antlr.Token { return s.id }

func (s *LValuePrimeContext) SetId(v antlr.Token) { s.id = v }

func (s *LValuePrimeContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, 0)
}

func (s *LValuePrimeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *LValuePrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValuePrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValuePrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLValuePrime(s)
	}
}

func (s *LValuePrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLValuePrime(s)
	}
}

func (p *GoliteParser) LValuePrime() (localctx ILValuePrimeContext) {
	this := p
	_ = this

	localctx = NewLValuePrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_lValuePrime)

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
		p.SetState(290)
		p.Match(GoliteParserDOT)
	}
	{
		p.SetState(291)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*LValuePrimeContext).id = _m
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllBoolTerm() []IBoolTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolTermContext); ok {
			len++
		}
	}

	tst := make([]IBoolTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolTermContext); ok {
			tst[i] = t.(IBoolTermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) BoolTerm(i int) IBoolTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
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

	return t.(IBoolTermContext)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoliteParserRULE_expression)
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
		p.SetState(293)
		p.BoolTerm()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(294)
			p.Match(GoliteParserOR)
		}
		{
			p.SetState(295)
			p.BoolTerm()
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolTermContext is an interface to support dynamic dispatch.
type IBoolTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTermContext differentiates from other interfaces.
	IsBoolTermContext()
}

type BoolTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTermContext() *BoolTermContext {
	var p = new(BoolTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTerm
	return p
}

func (*BoolTermContext) IsBoolTermContext() {}

func NewBoolTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTermContext {
	var p = new(BoolTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolTerm

	return p
}

func (s *BoolTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTermContext) AllEqualTerm() []IEqualTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualTermContext); ok {
			len++
		}
	}

	tst := make([]IEqualTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualTermContext); ok {
			tst[i] = t.(IEqualTermContext)
			i++
		}
	}

	return tst
}

func (s *BoolTermContext) EqualTerm(i int) IEqualTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
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

	return t.(IEqualTermContext)
}

func (s *BoolTermContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserAND)
}

func (s *BoolTermContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserAND, i)
}

func (s *BoolTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolTerm(s)
	}
}

func (s *BoolTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolTerm(s)
	}
}

func (p *GoliteParser) BoolTerm() (localctx IBoolTermContext) {
	this := p
	_ = this

	localctx = NewBoolTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoliteParserRULE_boolTerm)
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
		p.SetState(301)
		p.EqualTerm()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserAND {
		{
			p.SetState(302)
			p.Match(GoliteParserAND)
		}
		{
			p.SetState(303)
			p.EqualTerm()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEqualTermContext is an interface to support dynamic dispatch.
type IEqualTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualTermContext differentiates from other interfaces.
	IsEqualTermContext()
}

type EqualTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualTermContext() *EqualTermContext {
	var p = new(EqualTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTerm
	return p
}

func (*EqualTermContext) IsEqualTermContext() {}

func NewEqualTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermContext {
	var p = new(EqualTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTerm

	return p
}

func (s *EqualTermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermContext) AllRelationTerm() []IRelationTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationTermContext); ok {
			len++
		}
	}

	tst := make([]IRelationTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationTermContext); ok {
			tst[i] = t.(IRelationTermContext)
			i++
		}
	}

	return tst
}

func (s *EqualTermContext) RelationTerm(i int) IRelationTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermContext); ok {
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

	return t.(IRelationTermContext)
}

func (s *EqualTermContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserEQ)
}

func (s *EqualTermContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserEQ, i)
}

func (s *EqualTermContext) AllNE() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserNE)
}

func (s *EqualTermContext) NE(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserNE, i)
}

func (s *EqualTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTerm(s)
	}
}

func (s *EqualTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTerm(s)
	}
}

func (p *GoliteParser) EqualTerm() (localctx IEqualTermContext) {
	this := p
	_ = this

	localctx = NewEqualTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoliteParserRULE_equalTerm)
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
		p.SetState(309)
		p.RelationTerm()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserEQ || _la == GoliteParserNE {
		{
			p.SetState(310)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserEQ || _la == GoliteParserNE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(311)
			p.RelationTerm()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationTermContext is an interface to support dynamic dispatch.
type IRelationTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationTermContext differentiates from other interfaces.
	IsRelationTermContext()
}

type RelationTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationTermContext() *RelationTermContext {
	var p = new(RelationTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTerm
	return p
}

func (*RelationTermContext) IsRelationTermContext() {}

func NewRelationTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationTermContext {
	var p = new(RelationTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationTerm

	return p
}

func (s *RelationTermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationTermContext) AllSimpleTerm() []ISimpleTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleTermContext); ok {
			len++
		}
	}

	tst := make([]ISimpleTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleTermContext); ok {
			tst[i] = t.(ISimpleTermContext)
			i++
		}
	}

	return tst
}

func (s *RelationTermContext) SimpleTerm(i int) ISimpleTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
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

	return t.(ISimpleTermContext)
}

func (s *RelationTermContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserGE)
}

func (s *RelationTermContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserGE, i)
}

func (s *RelationTermContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserLT)
}

func (s *RelationTermContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserLT, i)
}

func (s *RelationTermContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserLE)
}

func (s *RelationTermContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserLE, i)
}

func (s *RelationTermContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserGT)
}

func (s *RelationTermContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserGT, i)
}

func (s *RelationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationTerm(s)
	}
}

func (s *RelationTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationTerm(s)
	}
}

func (p *GoliteParser) RelationTerm() (localctx IRelationTermContext) {
	this := p
	_ = this

	localctx = NewRelationTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoliteParserRULE_relationTerm)
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
		p.SetState(317)
		p.SimpleTerm()
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0 {
		{
			p.SetState(318)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(319)
			p.SimpleTerm()
		}

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISimpleTermContext is an interface to support dynamic dispatch.
type ISimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleTermContext differentiates from other interfaces.
	IsSimpleTermContext()
}

type SimpleTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTermContext() *SimpleTermContext {
	var p = new(SimpleTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTerm
	return p
}

func (*SimpleTermContext) IsSimpleTermContext() {}

func NewSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermContext {
	var p = new(SimpleTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTerm

	return p
}

func (s *SimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *SimpleTermContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *SimpleTermContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserPLUS)
}

func (s *SimpleTermContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, i)
}

func (s *SimpleTermContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserMINUS)
}

func (s *SimpleTermContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, i)
}

func (s *SimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTerm(s)
	}
}

func (s *SimpleTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTerm(s)
	}
}

func (p *GoliteParser) SimpleTerm() (localctx ISimpleTermContext) {
	this := p
	_ = this

	localctx = NewSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoliteParserRULE_simpleTerm)
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
		p.SetState(325)
		p.Term()
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPLUS || _la == GoliteParserMINUS {
		{
			p.SetState(326)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserPLUS || _la == GoliteParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(327)
			p.Term()
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllUnaryTerm() []IUnaryTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryTermContext); ok {
			len++
		}
	}

	tst := make([]IUnaryTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryTermContext); ok {
			tst[i] = t.(IUnaryTermContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) UnaryTerm(i int) IUnaryTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
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

	return t.(IUnaryTermContext)
}

func (s *TermContext) AllMULT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserMULT)
}

func (s *TermContext) MULT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserMULT, i)
}

func (s *TermContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDIV)
}

func (s *TermContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDIV, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoliteParserRULE_term)
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
		p.SetState(333)
		p.UnaryTerm()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserMULT || _la == GoliteParserDIV {
		{
			p.SetState(334)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserMULT || _la == GoliteParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(335)
			p.UnaryTerm()
		}

		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnaryTermContext is an interface to support dynamic dispatch.
type IUnaryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryTermContext differentiates from other interfaces.
	IsUnaryTermContext()
}

type UnaryTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryTermContext() *UnaryTermContext {
	var p = new(UnaryTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryTerm
	return p
}

func (*UnaryTermContext) IsUnaryTermContext() {}

func NewUnaryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryTermContext {
	var p = new(UnaryTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryTerm

	return p
}

func (s *UnaryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryTermContext) NOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOT, 0)
}

func (s *UnaryTermContext) SelectorTerm() ISelectorTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTermContext)
}

func (s *UnaryTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *UnaryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryTerm(s)
	}
}

func (s *UnaryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryTerm(s)
	}
}

func (p *GoliteParser) UnaryTerm() (localctx IUnaryTermContext) {
	this := p
	_ = this

	localctx = NewUnaryTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoliteParserRULE_unaryTerm)

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

	p.SetState(346)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.Match(GoliteParserNOT)
		}
		{
			p.SetState(342)
			p.SelectorTerm()
		}

	case GoliteParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(GoliteParserMINUS)
		}
		{
			p.SetState(344)
			p.SelectorTerm()
		}

	case GoliteParserLPAREN, GoliteParserNEW, GoliteParserINT_LIT, GoliteParserBOOL_LIT, GoliteParserNIL_LIT, GoliteParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(345)
			p.SelectorTerm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectorTermContext is an interface to support dynamic dispatch.
type ISelectorTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorTermContext differentiates from other interfaces.
	IsSelectorTermContext()
}

type SelectorTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermContext() *SelectorTermContext {
	var p = new(SelectorTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTerm
	return p
}

func (*SelectorTermContext) IsSelectorTermContext() {}

func NewSelectorTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermContext {
	var p = new(SelectorTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTerm

	return p
}

func (s *SelectorTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectorTermContext) AllSelectorTermPrime() []ISelectorTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ISelectorTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorTermPrimeContext); ok {
			tst[i] = t.(ISelectorTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *SelectorTermContext) SelectorTermPrime(i int) ISelectorTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermPrimeContext); ok {
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

	return t.(ISelectorTermPrimeContext)
}

func (s *SelectorTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTerm(s)
	}
}

func (s *SelectorTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTerm(s)
	}
}

func (p *GoliteParser) SelectorTerm() (localctx ISelectorTermContext) {
	this := p
	_ = this

	localctx = NewSelectorTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoliteParserRULE_selectorTerm)
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
		p.SetState(348)
		p.Factor()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(349)
			p.SelectorTermPrime()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectorTermPrimeContext is an interface to support dynamic dispatch.
type ISelectorTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorTermPrimeContext differentiates from other interfaces.
	IsSelectorTermPrimeContext()
}

type SelectorTermPrimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermPrimeContext() *SelectorTermPrimeContext {
	var p = new(SelectorTermPrimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTermPrime
	return p
}

func (*SelectorTermPrimeContext) IsSelectorTermPrimeContext() {}

func NewSelectorTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermPrimeContext {
	var p = new(SelectorTermPrimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTermPrime

	return p
}

func (s *SelectorTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermPrimeContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, 0)
}

func (s *SelectorTermPrimeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *SelectorTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTermPrime(s)
	}
}

func (s *SelectorTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTermPrime(s)
	}
}

func (p *GoliteParser) SelectorTermPrime() (localctx ISelectorTermPrimeContext) {
	this := p
	_ = this

	localctx = NewSelectorTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GoliteParserRULE_selectorTermPrime)

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
		p.Match(GoliteParserDOT)
	}
	{
		p.SetState(356)
		p.Match(GoliteParserIDENT)
	}

	return localctx
}

// ISubfactorContext is an interface to support dynamic dispatch.
type ISubfactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubfactorContext differentiates from other interfaces.
	IsSubfactorContext()
}

type SubfactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubfactorContext() *SubfactorContext {
	var p = new(SubfactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_subfactor
	return p
}

func (*SubfactorContext) IsSubfactorContext() {}

func NewSubfactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubfactorContext {
	var p = new(SubfactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_subfactor

	return p
}

func (s *SubfactorContext) GetParser() antlr.Parser { return s.parser }

func (s *SubfactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *SubfactorContext) Expression() IExpressionContext {
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

func (s *SubfactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *SubfactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubfactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubfactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSubfactor(s)
	}
}

func (s *SubfactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSubfactor(s)
	}
}

func (p *GoliteParser) Subfactor() (localctx ISubfactorContext) {
	this := p
	_ = this

	localctx = NewSubfactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GoliteParserRULE_subfactor)

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
		p.SetState(358)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(359)
		p.Expression()
	}
	{
		p.SetState(360)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_functioncall
	return p
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *FunctioncallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (p *GoliteParser) Functioncall() (localctx IFunctioncallContext) {
	this := p
	_ = this

	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GoliteParserRULE_functioncall)
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
		p.SetState(362)
		p.Match(GoliteParserIDENT)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserLPAREN {
		{
			p.SetState(363)
			p.Arguments()
		}

	}

	return localctx
}

// IAllocationContext is an interface to support dynamic dispatch.
type IAllocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// IsAllocationContext differentiates from other interfaces.
	IsAllocationContext()
}

type AllocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyAllocationContext() *AllocationContext {
	var p = new(AllocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_allocation
	return p
}

func (*AllocationContext) IsAllocationContext() {}

func NewAllocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocationContext {
	var p = new(AllocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_allocation

	return p
}

func (s *AllocationContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocationContext) GetKey() antlr.Token { return s.key }

func (s *AllocationContext) SetKey(v antlr.Token) { s.key = v }

func (s *AllocationContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *AllocationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENT, 0)
}

func (s *AllocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAllocation(s)
	}
}

func (s *AllocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAllocation(s)
	}
}

func (p *GoliteParser) Allocation() (localctx IAllocationContext) {
	this := p
	_ = this

	localctx = NewAllocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GoliteParserRULE_allocation)

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
		p.SetState(366)
		p.Match(GoliteParserNEW)
	}
	{
		p.SetState(367)

		var _m = p.Match(GoliteParserIDENT)

		localctx.(*AllocationContext).key = _m
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Subfactor() ISubfactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubfactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubfactorContext)
}

func (s *FactorContext) Functioncall() IFunctioncallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctioncallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *FactorContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT_LIT, 0)
}

func (s *FactorContext) Allocation() IAllocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllocationContext)
}

func (s *FactorContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL_LIT, 0)
}

func (s *FactorContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL_LIT, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GoliteParserRULE_factor)

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

	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.Subfactor()
		}

	case GoliteParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Functioncall()
		}

	case GoliteParserINT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.Match(GoliteParserINT_LIT)
		}

	case GoliteParserNEW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(372)
			p.Allocation()
		}

	case GoliteParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(373)
			p.Match(GoliteParserBOOL_LIT)
		}

	case GoliteParserNIL_LIT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(374)
			p.Match(GoliteParserNIL_LIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
