package parser

import (
	"fmt"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	"golite/types"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Interface for the parser
type Parser interface {
	Parse() *ast.Program
	GetErrors() []*context.CompilerError
}

// parser struct
type goliteParser struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	*BaseGoliteParserListener
	parser *GoliteParser
	lexer  lexer.Lexer
	errors []*context.CompilerError
	nodes  map[string]interface{}
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
	wrappedParser := &goliteParser{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil, make(map[string]interface{})}
	p := NewGoliteParser(lexer.GetTokenStream())
	p.RemoveErrorListeners()
	p.AddErrorListener(wrappedParser)
	wrappedParser.parser = p
	return wrappedParser
}

// Parse the tokens in the token stream as they are lexed
func (gParser *goliteParser) Parse() *ast.Program {
	ctx := gParser.parser.Program()

	// Check if there are any errors in the parsing and lexing stages
	if context.HasErrors(gParser.lexer.GetErrors()) || context.HasErrors(gParser.GetErrors()) {
		return nil
	}

	antlr.ParseTreeWalkerDefault.Walk(gParser, ctx)
	_, _, key := GetTokenInfo(ctx)
	return gParser.nodes[key].(*ast.Program)

}

/******************* Implementation of the Listeners **************************/

// GetTokenInfo gerates a unique key for each node in the ParserTree
func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	key := fmt.Sprintf("%d,%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), key
}

// ExitFactor is called when exiting the program production.
func (gParser *goliteParser) ExitProgram(c *ProgramContext) {
	line, col, key := GetTokenInfo(c)
	var stmts []ast.Statement
	for _, stmtCtx := range c.AllStatement() {
		_, _, key := GetTokenInfo(stmtCtx)
		astStmt := gParser.nodes[key].(ast.Statement)
		stmts = append(stmts, astStmt)
	}
	gParser.nodes[key] = ast.NewProgram(stmts, token.NewToken(line, col))
}

// ExitFactor is called when exiting the decl production.
func (gParser *goliteParser) ExitDecl(c *DeclContext) {

	line, col, key := GetTokenInfo(c)
	var expr ast.Expression
	if exprCtx := c.GetExpr(); exprCtx != nil {
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr = gParser.nodes[exprKey].(ast.Expression)
	} else {
		expr = nil
	}
	decl := ast.NewDeclStm(c.GetTyDec().GetId().GetText(), types.StringToType(c.GetTyDec().GetTy().GetText()), expr, token.NewToken(line, col))
	gParser.nodes[key] = decl
}

// EnterAssignment is called when exiting the assignment production.
func (gParser *goliteParser) ExitAssignment(c *AssignmentContext) {
	line, col, key := GetTokenInfo(c)
	_, _, eprKey := GetTokenInfo(c.Expression())
	expr := gParser.nodes[eprKey].(ast.Expression)
	gParser.nodes[key] = ast.NewAssignStm(c.IDENTIFIER().GetText(), expr, token.NewToken(line, col))
}

// EnterPrint is called when exiting the print production.
func (gParser *goliteParser) ExitPrint(c *PrintContext) {
	line, col, key := GetTokenInfo(c)
	_, _, eprKey := GetTokenInfo(c.Expression())
	expr := gParser.nodes[eprKey].(ast.Expression)
	gParser.nodes[key] = ast.NewPrintStm(expr, token.NewToken(line, col))
}

// ExitExpression is called when exiting the expression production.
func (gParser *goliteParser) ExitExpression(c *ExpressionContext) {
	_, _, key := GetTokenInfo(c)
	termContexts := c.AllExpressionPrime()
	_, _, eprKey := GetTokenInfo(c.RelationTerm())
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		// Exercise for you: You need to think about how you can build a nested BinOp ast from the termContexts.
	}
}

// ExitRelationTerm is called when exiting the relationTerm production.
func (gParser *goliteParser) ExitRelationTerm(c *RelationTermContext) {

	_, _, key := GetTokenInfo(c)
	termContexts := c.AllRelationTermPrime()
	_, _, eprKey := GetTokenInfo(c.Factor())
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		// Exercise for you: You need to think about how you can build a nested BinOp ast from the termContexts.
	}
}

// ExitFactor is called when exiting the factor production.
func (gParser *goliteParser) ExitFactor(c *FactorContext) {
	line, col, key := GetTokenInfo(c)

	if intFactor := c.NUMBER(); intFactor != nil {
		intValue, _ := strconv.Atoi(intFactor.GetText())
		gParser.nodes[key] = ast.NewIntLit(int64(intValue), token.NewToken(line, col))
	} else if trueFactor := c.TRUE(); trueFactor != nil {
		gParser.nodes[key] = ast.NewBoolLit(true, token.NewToken(line, col))
	} else if falseFactor := c.FALSE(); falseFactor != nil {
		gParser.nodes[key] = ast.NewBoolLit(false, token.NewToken(line, col))
	} else if idFactor := c.IDENTIFIER(); idFactor != nil {
		gParser.nodes[key] = ast.NewVariable(idFactor.GetText(), token.NewToken(line, col))
	}
}
