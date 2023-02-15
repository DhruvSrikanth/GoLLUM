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
		"ASSIGN", "AND", "OR", "NOT", "LT", "GT", "LE", "GE", "EQ", "NE", "PLUS",
		"MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "SEMICOLON",
		"COMMA", "VAR", "TYPE", "NEW", "DELETE", "STRUCT", "DOT", "FUNC", "RET",
		"IF", "ELSE", "FOR", "SCAN", "PRINTF", "INT_LIT", "STRING_LIT", "BOOL_LIT",
		"NIL_LIT", "INT", "BOOL", "IDENT", "WS", "COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 266, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 199, 8, 33, 10, 33, 12, 33,
		202, 9, 33, 1, 33, 3, 33, 205, 8, 33, 1, 34, 1, 34, 5, 34, 209, 8, 34,
		10, 34, 12, 34, 212, 9, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 225, 8, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 5, 39, 242, 8, 39, 10, 39, 12, 39, 245, 9, 39, 1, 40, 4, 40,
		248, 8, 40, 11, 40, 12, 40, 249, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 5, 41, 258, 8, 41, 10, 41, 12, 41, 261, 9, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 2, 210, 259, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 1, 0, 6, 1,
		0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 48, 2, 0, 65, 90, 97, 122, 3, 0, 48,
		57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 272, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3, 87,
		1, 0, 0, 0, 5, 90, 1, 0, 0, 0, 7, 93, 1, 0, 0, 0, 9, 95, 1, 0, 0, 0, 11,
		97, 1, 0, 0, 0, 13, 99, 1, 0, 0, 0, 15, 102, 1, 0, 0, 0, 17, 105, 1, 0,
		0, 0, 19, 108, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 113, 1, 0, 0, 0, 25,
		115, 1, 0, 0, 0, 27, 117, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 121, 1,
		0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 127, 1, 0, 0, 0,
		39, 129, 1, 0, 0, 0, 41, 131, 1, 0, 0, 0, 43, 135, 1, 0, 0, 0, 45, 140,
		1, 0, 0, 0, 47, 144, 1, 0, 0, 0, 49, 151, 1, 0, 0, 0, 51, 158, 1, 0, 0,
		0, 53, 160, 1, 0, 0, 0, 55, 165, 1, 0, 0, 0, 57, 172, 1, 0, 0, 0, 59, 175,
		1, 0, 0, 0, 61, 180, 1, 0, 0, 0, 63, 184, 1, 0, 0, 0, 65, 189, 1, 0, 0,
		0, 67, 204, 1, 0, 0, 0, 69, 206, 1, 0, 0, 0, 71, 224, 1, 0, 0, 0, 73, 226,
		1, 0, 0, 0, 75, 230, 1, 0, 0, 0, 77, 234, 1, 0, 0, 0, 79, 239, 1, 0, 0,
		0, 81, 247, 1, 0, 0, 0, 83, 253, 1, 0, 0, 0, 85, 86, 5, 61, 0, 0, 86, 2,
		1, 0, 0, 0, 87, 88, 5, 38, 0, 0, 88, 89, 5, 38, 0, 0, 89, 4, 1, 0, 0, 0,
		90, 91, 5, 124, 0, 0, 91, 92, 5, 124, 0, 0, 92, 6, 1, 0, 0, 0, 93, 94,
		5, 33, 0, 0, 94, 8, 1, 0, 0, 0, 95, 96, 5, 60, 0, 0, 96, 10, 1, 0, 0, 0,
		97, 98, 5, 62, 0, 0, 98, 12, 1, 0, 0, 0, 99, 100, 5, 60, 0, 0, 100, 101,
		5, 61, 0, 0, 101, 14, 1, 0, 0, 0, 102, 103, 5, 62, 0, 0, 103, 104, 5, 61,
		0, 0, 104, 16, 1, 0, 0, 0, 105, 106, 5, 61, 0, 0, 106, 107, 5, 61, 0, 0,
		107, 18, 1, 0, 0, 0, 108, 109, 5, 33, 0, 0, 109, 110, 5, 61, 0, 0, 110,
		20, 1, 0, 0, 0, 111, 112, 5, 43, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114, 5,
		45, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 5, 42, 0, 0, 116, 26, 1, 0, 0,
		0, 117, 118, 5, 47, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 40, 0, 0, 120,
		30, 1, 0, 0, 0, 121, 122, 5, 41, 0, 0, 122, 32, 1, 0, 0, 0, 123, 124, 5,
		123, 0, 0, 124, 34, 1, 0, 0, 0, 125, 126, 5, 125, 0, 0, 126, 36, 1, 0,
		0, 0, 127, 128, 5, 59, 0, 0, 128, 38, 1, 0, 0, 0, 129, 130, 5, 44, 0, 0,
		130, 40, 1, 0, 0, 0, 131, 132, 5, 118, 0, 0, 132, 133, 5, 97, 0, 0, 133,
		134, 5, 114, 0, 0, 134, 42, 1, 0, 0, 0, 135, 136, 5, 116, 0, 0, 136, 137,
		5, 121, 0, 0, 137, 138, 5, 112, 0, 0, 138, 139, 5, 101, 0, 0, 139, 44,
		1, 0, 0, 0, 140, 141, 5, 110, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5,
		119, 0, 0, 143, 46, 1, 0, 0, 0, 144, 145, 5, 100, 0, 0, 145, 146, 5, 101,
		0, 0, 146, 147, 5, 108, 0, 0, 147, 148, 5, 101, 0, 0, 148, 149, 5, 116,
		0, 0, 149, 150, 5, 101, 0, 0, 150, 48, 1, 0, 0, 0, 151, 152, 5, 115, 0,
		0, 152, 153, 5, 116, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 117, 0,
		0, 155, 156, 5, 99, 0, 0, 156, 157, 5, 116, 0, 0, 157, 50, 1, 0, 0, 0,
		158, 159, 5, 46, 0, 0, 159, 52, 1, 0, 0, 0, 160, 161, 5, 102, 0, 0, 161,
		162, 5, 117, 0, 0, 162, 163, 5, 110, 0, 0, 163, 164, 5, 99, 0, 0, 164,
		54, 1, 0, 0, 0, 165, 166, 5, 114, 0, 0, 166, 167, 5, 101, 0, 0, 167, 168,
		5, 116, 0, 0, 168, 169, 5, 117, 0, 0, 169, 170, 5, 114, 0, 0, 170, 171,
		5, 110, 0, 0, 171, 56, 1, 0, 0, 0, 172, 173, 5, 105, 0, 0, 173, 174, 5,
		102, 0, 0, 174, 58, 1, 0, 0, 0, 175, 176, 5, 101, 0, 0, 176, 177, 5, 108,
		0, 0, 177, 178, 5, 115, 0, 0, 178, 179, 5, 101, 0, 0, 179, 60, 1, 0, 0,
		0, 180, 181, 5, 102, 0, 0, 181, 182, 5, 111, 0, 0, 182, 183, 5, 114, 0,
		0, 183, 62, 1, 0, 0, 0, 184, 185, 5, 115, 0, 0, 185, 186, 5, 99, 0, 0,
		186, 187, 5, 97, 0, 0, 187, 188, 5, 110, 0, 0, 188, 64, 1, 0, 0, 0, 189,
		190, 5, 112, 0, 0, 190, 191, 5, 114, 0, 0, 191, 192, 5, 105, 0, 0, 192,
		193, 5, 110, 0, 0, 193, 194, 5, 116, 0, 0, 194, 195, 5, 102, 0, 0, 195,
		66, 1, 0, 0, 0, 196, 200, 7, 0, 0, 0, 197, 199, 7, 1, 0, 0, 198, 197, 1,
		0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0,
		0, 201, 205, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 205, 7, 2, 0, 0, 204,
		196, 1, 0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 68, 1, 0, 0, 0, 206, 210, 5,
		34, 0, 0, 207, 209, 9, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0,
		0, 210, 211, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212,
		210, 1, 0, 0, 0, 213, 214, 5, 34, 0, 0, 214, 70, 1, 0, 0, 0, 215, 216,
		5, 116, 0, 0, 216, 217, 5, 114, 0, 0, 217, 218, 5, 117, 0, 0, 218, 225,
		5, 101, 0, 0, 219, 220, 5, 102, 0, 0, 220, 221, 5, 97, 0, 0, 221, 222,
		5, 108, 0, 0, 222, 223, 5, 115, 0, 0, 223, 225, 5, 101, 0, 0, 224, 215,
		1, 0, 0, 0, 224, 219, 1, 0, 0, 0, 225, 72, 1, 0, 0, 0, 226, 227, 5, 110,
		0, 0, 227, 228, 5, 105, 0, 0, 228, 229, 5, 108, 0, 0, 229, 74, 1, 0, 0,
		0, 230, 231, 5, 105, 0, 0, 231, 232, 5, 110, 0, 0, 232, 233, 5, 116, 0,
		0, 233, 76, 1, 0, 0, 0, 234, 235, 5, 98, 0, 0, 235, 236, 5, 111, 0, 0,
		236, 237, 5, 111, 0, 0, 237, 238, 5, 108, 0, 0, 238, 78, 1, 0, 0, 0, 239,
		243, 7, 3, 0, 0, 240, 242, 7, 4, 0, 0, 241, 240, 1, 0, 0, 0, 242, 245,
		1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 80, 1, 0,
		0, 0, 245, 243, 1, 0, 0, 0, 246, 248, 7, 5, 0, 0, 247, 246, 1, 0, 0, 0,
		248, 249, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 252, 6, 40, 0, 0, 252, 82, 1, 0, 0, 0, 253, 254,
		5, 47, 0, 0, 254, 255, 5, 47, 0, 0, 255, 259, 1, 0, 0, 0, 256, 258, 9,
		0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 260, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262,
		263, 5, 10, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 6, 41, 0, 0, 265, 84,
		1, 0, 0, 0, 8, 0, 200, 204, 210, 224, 243, 249, 259, 1, 6, 0, 0,
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
	GoliteLexerASSIGN     = 1
	GoliteLexerAND        = 2
	GoliteLexerOR         = 3
	GoliteLexerNOT        = 4
	GoliteLexerLT         = 5
	GoliteLexerGT         = 6
	GoliteLexerLE         = 7
	GoliteLexerGE         = 8
	GoliteLexerEQ         = 9
	GoliteLexerNE         = 10
	GoliteLexerPLUS       = 11
	GoliteLexerMINUS      = 12
	GoliteLexerMULT       = 13
	GoliteLexerDIV        = 14
	GoliteLexerLPAREN     = 15
	GoliteLexerRPAREN     = 16
	GoliteLexerLBRACE     = 17
	GoliteLexerRBRACE     = 18
	GoliteLexerSEMICOLON  = 19
	GoliteLexerCOMMA      = 20
	GoliteLexerVAR        = 21
	GoliteLexerTYPE       = 22
	GoliteLexerNEW        = 23
	GoliteLexerDELETE     = 24
	GoliteLexerSTRUCT     = 25
	GoliteLexerDOT        = 26
	GoliteLexerFUNC       = 27
	GoliteLexerRET        = 28
	GoliteLexerIF         = 29
	GoliteLexerELSE       = 30
	GoliteLexerFOR        = 31
	GoliteLexerSCAN       = 32
	GoliteLexerPRINTF     = 33
	GoliteLexerINT_LIT    = 34
	GoliteLexerSTRING_LIT = 35
	GoliteLexerBOOL_LIT   = 36
	GoliteLexerNIL_LIT    = 37
	GoliteLexerINT        = 38
	GoliteLexerBOOL       = 39
	GoliteLexerIDENT      = 40
	GoliteLexerWS         = 41
	GoliteLexerCOMMENT    = 42
)
