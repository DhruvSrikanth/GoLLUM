// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var golitelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golitelexerLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "'&&'", "'||'", "'!'", "'<'", "'>'", "'<='", "'>='", "'=='",
		"'!='", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'{'", "'}'", "';'",
		"','", "'var'", "'type'", "'new'", "'delete'", "'struct'", "'.'", "'func'",
		"'return'", "'if'", "'else'", "'for'", "'scan'", "'printf'", "", "",
		"'nil'", "'int'", "'bool'", "'string'",
	}
	staticData.symbolicNames = []string{
		"", "ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE",
		"PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"SEMICOLON", "COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT",
		"FUNC", "RET", "IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "BOOL_LIT",
		"NIL_LIT", "INT", "BOOL", "STRING", "PTR", "IDENT", "WS", "COMMENT",
	}
	staticData.ruleNames = []string{
		"ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE", "PLUS",
		"MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "SEMICOLON",
		"COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT", "FUNC", "RET",
		"IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "BOOL_LIT", "NIL_LIT",
		"INT", "BOOL", "STRING", "PTR", "IDENT", "WS", "COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 275, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 3, 33, 200, 8, 33, 1, 33,
		1, 33, 5, 33, 204, 8, 33, 10, 33, 12, 33, 207, 9, 33, 1, 33, 3, 33, 210,
		8, 33, 1, 33, 3, 33, 213, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 224, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40,
		5, 40, 251, 8, 40, 10, 40, 12, 40, 254, 9, 40, 1, 41, 4, 41, 257, 8, 41,
		11, 41, 12, 41, 258, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 267,
		8, 42, 10, 42, 12, 42, 270, 9, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 268,
		0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 1, 0, 7, 1, 0, 45, 45,
		1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 48, 2, 0, 65, 90, 97, 122, 3, 0,
		48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 282, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		1, 87, 1, 0, 0, 0, 3, 89, 1, 0, 0, 0, 5, 92, 1, 0, 0, 0, 7, 95, 1, 0, 0,
		0, 9, 97, 1, 0, 0, 0, 11, 99, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15, 104,
		1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 110, 1, 0, 0, 0, 21, 113, 1, 0, 0,
		0, 23, 115, 1, 0, 0, 0, 25, 117, 1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 121,
		1, 0, 0, 0, 31, 123, 1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 127, 1, 0, 0,
		0, 37, 129, 1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 133, 1, 0, 0, 0, 43, 137,
		1, 0, 0, 0, 45, 142, 1, 0, 0, 0, 47, 146, 1, 0, 0, 0, 49, 153, 1, 0, 0,
		0, 51, 160, 1, 0, 0, 0, 53, 162, 1, 0, 0, 0, 55, 167, 1, 0, 0, 0, 57, 174,
		1, 0, 0, 0, 59, 177, 1, 0, 0, 0, 61, 182, 1, 0, 0, 0, 63, 186, 1, 0, 0,
		0, 65, 191, 1, 0, 0, 0, 67, 212, 1, 0, 0, 0, 69, 223, 1, 0, 0, 0, 71, 225,
		1, 0, 0, 0, 73, 229, 1, 0, 0, 0, 75, 233, 1, 0, 0, 0, 77, 238, 1, 0, 0,
		0, 79, 245, 1, 0, 0, 0, 81, 248, 1, 0, 0, 0, 83, 256, 1, 0, 0, 0, 85, 262,
		1, 0, 0, 0, 87, 88, 5, 61, 0, 0, 88, 2, 1, 0, 0, 0, 89, 90, 5, 38, 0, 0,
		90, 91, 5, 38, 0, 0, 91, 4, 1, 0, 0, 0, 92, 93, 5, 124, 0, 0, 93, 94, 5,
		124, 0, 0, 94, 6, 1, 0, 0, 0, 95, 96, 5, 33, 0, 0, 96, 8, 1, 0, 0, 0, 97,
		98, 5, 60, 0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 62, 0, 0, 100, 12, 1,
		0, 0, 0, 101, 102, 5, 60, 0, 0, 102, 103, 5, 61, 0, 0, 103, 14, 1, 0, 0,
		0, 104, 105, 5, 62, 0, 0, 105, 106, 5, 61, 0, 0, 106, 16, 1, 0, 0, 0, 107,
		108, 5, 61, 0, 0, 108, 109, 5, 61, 0, 0, 109, 18, 1, 0, 0, 0, 110, 111,
		5, 33, 0, 0, 111, 112, 5, 61, 0, 0, 112, 20, 1, 0, 0, 0, 113, 114, 5, 43,
		0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 45, 0, 0, 116, 24, 1, 0, 0, 0,
		117, 118, 5, 42, 0, 0, 118, 26, 1, 0, 0, 0, 119, 120, 5, 47, 0, 0, 120,
		28, 1, 0, 0, 0, 121, 122, 5, 40, 0, 0, 122, 30, 1, 0, 0, 0, 123, 124, 5,
		41, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 123, 0, 0, 126, 34, 1, 0, 0,
		0, 127, 128, 5, 125, 0, 0, 128, 36, 1, 0, 0, 0, 129, 130, 5, 59, 0, 0,
		130, 38, 1, 0, 0, 0, 131, 132, 5, 44, 0, 0, 132, 40, 1, 0, 0, 0, 133, 134,
		5, 118, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 114, 0, 0, 136, 42, 1,
		0, 0, 0, 137, 138, 5, 116, 0, 0, 138, 139, 5, 121, 0, 0, 139, 140, 5, 112,
		0, 0, 140, 141, 5, 101, 0, 0, 141, 44, 1, 0, 0, 0, 142, 143, 5, 110, 0,
		0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 119, 0, 0, 145, 46, 1, 0, 0, 0,
		146, 147, 5, 100, 0, 0, 147, 148, 5, 101, 0, 0, 148, 149, 5, 108, 0, 0,
		149, 150, 5, 101, 0, 0, 150, 151, 5, 116, 0, 0, 151, 152, 5, 101, 0, 0,
		152, 48, 1, 0, 0, 0, 153, 154, 5, 115, 0, 0, 154, 155, 5, 116, 0, 0, 155,
		156, 5, 114, 0, 0, 156, 157, 5, 117, 0, 0, 157, 158, 5, 99, 0, 0, 158,
		159, 5, 116, 0, 0, 159, 50, 1, 0, 0, 0, 160, 161, 5, 46, 0, 0, 161, 52,
		1, 0, 0, 0, 162, 163, 5, 102, 0, 0, 163, 164, 5, 117, 0, 0, 164, 165, 5,
		110, 0, 0, 165, 166, 5, 99, 0, 0, 166, 54, 1, 0, 0, 0, 167, 168, 5, 114,
		0, 0, 168, 169, 5, 101, 0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 117,
		0, 0, 171, 172, 5, 114, 0, 0, 172, 173, 5, 110, 0, 0, 173, 56, 1, 0, 0,
		0, 174, 175, 5, 105, 0, 0, 175, 176, 5, 102, 0, 0, 176, 58, 1, 0, 0, 0,
		177, 178, 5, 101, 0, 0, 178, 179, 5, 108, 0, 0, 179, 180, 5, 115, 0, 0,
		180, 181, 5, 101, 0, 0, 181, 60, 1, 0, 0, 0, 182, 183, 5, 102, 0, 0, 183,
		184, 5, 111, 0, 0, 184, 185, 5, 114, 0, 0, 185, 62, 1, 0, 0, 0, 186, 187,
		5, 115, 0, 0, 187, 188, 5, 99, 0, 0, 188, 189, 5, 97, 0, 0, 189, 190, 5,
		110, 0, 0, 190, 64, 1, 0, 0, 0, 191, 192, 5, 112, 0, 0, 192, 193, 5, 114,
		0, 0, 193, 194, 5, 105, 0, 0, 194, 195, 5, 110, 0, 0, 195, 196, 5, 116,
		0, 0, 196, 197, 5, 102, 0, 0, 197, 66, 1, 0, 0, 0, 198, 200, 7, 0, 0, 0,
		199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		205, 7, 1, 0, 0, 202, 204, 7, 2, 0, 0, 203, 202, 1, 0, 0, 0, 204, 207,
		1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 213, 1, 0,
		0, 0, 207, 205, 1, 0, 0, 0, 208, 210, 7, 0, 0, 0, 209, 208, 1, 0, 0, 0,
		209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 7, 3, 0, 0, 212,
		199, 1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 213, 68, 1, 0, 0, 0, 214, 215, 5,
		116, 0, 0, 215, 216, 5, 114, 0, 0, 216, 217, 5, 117, 0, 0, 217, 224, 5,
		101, 0, 0, 218, 219, 5, 102, 0, 0, 219, 220, 5, 97, 0, 0, 220, 221, 5,
		108, 0, 0, 221, 222, 5, 115, 0, 0, 222, 224, 5, 101, 0, 0, 223, 214, 1,
		0, 0, 0, 223, 218, 1, 0, 0, 0, 224, 70, 1, 0, 0, 0, 225, 226, 5, 110, 0,
		0, 226, 227, 5, 105, 0, 0, 227, 228, 5, 108, 0, 0, 228, 72, 1, 0, 0, 0,
		229, 230, 5, 105, 0, 0, 230, 231, 5, 110, 0, 0, 231, 232, 5, 116, 0, 0,
		232, 74, 1, 0, 0, 0, 233, 234, 5, 98, 0, 0, 234, 235, 5, 111, 0, 0, 235,
		236, 5, 111, 0, 0, 236, 237, 5, 108, 0, 0, 237, 76, 1, 0, 0, 0, 238, 239,
		5, 115, 0, 0, 239, 240, 5, 116, 0, 0, 240, 241, 5, 114, 0, 0, 241, 242,
		5, 105, 0, 0, 242, 243, 5, 110, 0, 0, 243, 244, 5, 103, 0, 0, 244, 78,
		1, 0, 0, 0, 245, 246, 5, 42, 0, 0, 246, 247, 3, 81, 40, 0, 247, 80, 1,
		0, 0, 0, 248, 252, 7, 4, 0, 0, 249, 251, 7, 5, 0, 0, 250, 249, 1, 0, 0,
		0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		82, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 257, 7, 6, 0, 0, 256, 255, 1,
		0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0,
		0, 259, 260, 1, 0, 0, 0, 260, 261, 6, 41, 0, 0, 261, 84, 1, 0, 0, 0, 262,
		263, 5, 47, 0, 0, 263, 264, 5, 47, 0, 0, 264, 268, 1, 0, 0, 0, 265, 267,
		9, 0, 0, 0, 266, 265, 1, 0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 269, 1, 0,
		0, 0, 268, 266, 1, 0, 0, 0, 269, 271, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0,
		271, 272, 5, 10, 0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 6, 42, 0, 0, 274,
		86, 1, 0, 0, 0, 9, 0, 199, 205, 209, 212, 223, 252, 258, 268, 1, 6, 0,
		0,
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

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &golitelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "GoliteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerASSIGN    = 1
	GoliteLexerAND       = 2
	GoliteLexerOR        = 3
	GoliteLexerNOT       = 4
	GoliteLexerLT        = 5
	GoliteLexerGT        = 6
	GoliteLexerLE        = 7
	GoliteLexerGE        = 8
	GoliteLexerEQ        = 9
	GoliteLexerNE        = 10
	GoliteLexerPLUS      = 11
	GoliteLexerMINUS     = 12
	GoliteLexerMULT      = 13
	GoliteLexerDIV       = 14
	GoliteLexerLPAREN    = 15
	GoliteLexerRPAREN    = 16
	GoliteLexerLBRACE    = 17
	GoliteLexerRBRACE    = 18
	GoliteLexerSEMICOLON = 19
	GoliteLexerCOMMA     = 20
	GoliteLexerVAR       = 21
	GoliteLexerTYPE      = 22
	GoliteLexerNEW       = 23
	GoliteLexerDELETE    = 24
	GoliteLexerSTRUCT    = 25
	GoliteLexerDOT       = 26
	GoliteLexerFUNC      = 27
	GoliteLexerRET       = 28
	GoliteLexerIF        = 29
	GoliteLexerELSE      = 30
	GoliteLexerFOR       = 31
	GoliteLexerSCAN      = 32
	GoliteLexerPRINTF    = 33
	GoliteLexerINT_LIT   = 34
	GoliteLexerBOOL_LIT  = 35
	GoliteLexerNIL_LIT   = 36
	GoliteLexerINT       = 37
	GoliteLexerBOOL      = 38
	GoliteLexerSTRING    = 39
	GoliteLexerPTR       = 40
	GoliteLexerIDENT     = 41
	GoliteLexerWS        = 42
	GoliteLexerCOMMENT   = 43
)
