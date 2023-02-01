package lexer

import (
	"fmt"
	"golite/context"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Lexer interface
type Lexer interface {
	GetTokenStream() *antlr.CommonTokenStream
	GetErrors() []*context.CompilerError
	PrintTokens()
}

// Lexer encapsulates the lexer
type lexer struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	lexer                       *GoliteLexer
	stream                      *antlr.CommonTokenStream
	errors                      []*context.CompilerError
}

// Return a beautified error message
func (lex *lexer) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	lex.errors = append(lex.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.LEXER,
	})
}

// Get the token stream for the lexer
func (lexer *lexer) GetTokenStream() *antlr.CommonTokenStream {
	return lexer.stream
}

// Print all tokens
func (lexer *lexer) PrintTokens() {
	// Get the token stream
	tokenStream := lexer.GetTokenStream()

	// Lex all tokens
	tokenStream.Fill()

	// Print all tokens
	for _, token := range tokenStream.GetAllTokens() {
		fmt.Println(token.GetText())
	}

	// Print the errors
	context.HasErrors(lexer.GetErrors())
}

// Get the errors from the lexer
func (lexer *lexer) GetErrors() []*context.CompilerError {
	return lexer.errors
}

// Get a new lexer
func NewLexer(inputSourcePath string) Lexer {
	input, _ := antlr.NewFileStream(inputSourcePath)
	lex := &lexer{antlr.NewDefaultErrorListener(), nil, nil, nil}
	antrLexer := NewGoliteLexer(input)
	antrLexer.RemoveErrorListeners()
	antrLexer.AddErrorListener(lex)
	stream := antlr.NewCommonTokenStream(antrLexer, 0)
	lex.lexer = antrLexer
	lex.stream = stream
	return lex
}
