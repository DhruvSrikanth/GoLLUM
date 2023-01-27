package parser

import (
	"golite/context"
	"golite/lexer"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Interface for the parser
type Parser interface {
	Parse() bool
	GetErrors() []*context.CompilerError
}

// parser struct
type goliteParser struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	// *BaseCalParserListener
	parser *GoliteParser
	lexer  lexer.Lexer
	errors []*context.CompilerError
}

// Get errors from the parser
func (gParser *goliteParser) GetErrors() []*context.CompilerError {
	return gParser.errors
}

// Syntax error handler for the parser
func (gParser *goliteParser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	gParser.errors = append(gParser.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.PARSER,
	})
}

// Return a new parser instance
func NewParser(lexer lexer.Lexer) Parser {
	wrappedParser := &goliteParser{antlr.NewDefaultErrorListener(), nil, lexer, nil}
	p := NewGoliteParser(lexer.GetTokenStream())
	p.RemoveErrorListeners()
	p.AddErrorListener(wrappedParser)
	wrappedParser.parser = p
	return wrappedParser
}

// Parse the tokens in the token stream as they are lexed
func (gParser *goliteParser) Parse() bool {
	gParser.parser.Program()
	return context.HasErrors(gParser.lexer.GetErrors()) ||
		context.HasErrors(gParser.GetErrors())
}
