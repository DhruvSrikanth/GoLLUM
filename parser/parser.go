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

	// Collect all struct declarations
	var strctTypes []ast.StructTypeStatement
	for _, strctTypeCtx := range c.GetTy().(*TypesContext).AllTypeDeclaration() {
		_, _, key := GetTokenInfo(strctTypeCtx)
		astStrctType := gParser.nodes[key].(ast.StructTypeStatement)
		strctTypes = append(strctTypes, astStrctType)
	}

	// Collect all global variable declarations
	var decls []ast.DeclarationStatement
	for _, declCtx := range c.GetDcls().(*DeclarationsContext).AllDeclaration() {
		_, _, declKey := GetTokenInfo(declCtx)
		decl := gParser.nodes[declKey].(ast.DeclarationStatement)
		decls = append(decls, decl)
		for _, moreDeclCtx := range declCtx.GetIdx().(*IdsContext).AllOtherids() {
			_, _, moreDeclKey := GetTokenInfo(moreDeclCtx)
			decl := gParser.nodes[moreDeclKey].(ast.DeclarationStatement)
			decls = append(decls, decl)
		}
	}

	// Collect all function declarations
	var funcs []ast.FunctionStatement
	for _, funcCtx := range c.GetFns().(*FunctionsContext).AllFunction() {
		_, _, funcKey := GetTokenInfo(funcCtx)
		funcStmt := gParser.nodes[funcKey].(ast.FunctionStatement)
		funcs = append(funcs, funcStmt)
	}

	gParser.nodes[key] = ast.NewProgram(strctTypes, decls, funcs, token.NewToken(line, col))
}

// ExitTypeDeclaration is called when exiting the typeDeclaration production.
func (gParser *goliteParser) ExitTypeDeclaration(c *TypeDeclarationContext) {
	// Get the key for the struct declaration
	line, col, key := GetTokenInfo(c)

	// Get the struct name
	strctName := c.GetId().GetText()

	// Get the struct fields
	var varDecls []ast.Decl
	fieldsCtx := c.GetFlds().(*FieldsContext)
	declCtx := fieldsCtx.GetDcl().(*DeclContext)
	_, _, varKey := GetTokenInfo(declCtx)
	// Add the first field (mandatory)
	varDecl := gParser.nodes[varKey].(*ast.Decl)
	varDecls = append(varDecls, *varDecl)
	for _, fieldCtx := range fieldsCtx.AllMorefields() {
		declCtx := fieldCtx.GetDcl().(*DeclContext)
		_, _, varKey := GetTokenInfo(declCtx)
		// Add the rest of the fields (optional)
		varDecl := gParser.nodes[varKey].(*ast.Decl)
		varDecls = append(varDecls, *varDecl)
	}

	// Create the struct type
	gParser.nodes[key] = ast.NewStructDecl(strctName, varDecls, token.NewToken(line, col))
}

// ExitDecl is called when exiting the decl production.
func (gParser *goliteParser) ExitDecl(c *DeclContext) {
	// Get the key for the declaration
	line, col, key := GetTokenInfo(c)
	// Get the struct field name
	varName := c.GetId().GetText()
	// Get the type of struct field
	varType := c.GetTy().GetText()
	// Create the struct field
	gParser.nodes[key] = ast.NewDeclStm(varName, types.StringToType(varType), token.NewToken(line, col))
}

// ExitDeclaration is called when exiting the declaration production.
func (gParser *goliteParser) ExitDeclaration(c *DeclarationContext) {
	// Get the key for the declaration
	line, col, key := GetTokenInfo(c)
	// Get the variable type
	varType := c.GetTy().GetText()
	// Get the variable name
	varName := c.Ids().GetId().GetText()
	// Create the variable declaration
	gParser.nodes[key] = ast.NewDeclaration(varName, types.StringToType(varType), token.NewToken(line, col))
	for _, moreIdCtx := range c.Ids().(*IdsContext).AllOtherids() {
		line, col, key := GetTokenInfo(moreIdCtx)
		// Get the variable name declared in the same statement
		varName = moreIdCtx.GetId().GetText()
		// Create the variable declaration
		gParser.nodes[key] = ast.NewDeclaration(varName, types.StringToType(varType), token.NewToken(line, col))
	}
}

// ExitFunction is called when exiting the function production.
func (gParser *goliteParser) ExitFunction(c *FunctionContext) {
	// Get the key for the function declaration
	line, col, key := GetTokenInfo(c)
	// Get the function name
	funcName := c.GetId().GetText()
	// Get the return type (if nil, cast to void)
	var returnType string
	if c.GetRty() != nil {
		returnType = c.GetRty().(*ReturnTypeContext).GetTy().GetText()
	} else {
		returnType = "void"
	}
	// Get the function parameters (if nil, cast to empty list)
	paramsCtx := c.GetParams().(*ParametersContext).ParametersPrime()
	var params []ast.Decl
	if paramsCtx != nil {
		// Get the first parameter
		declCtx := paramsCtx.GetDcl().(*DeclContext)
		_, _, paramKey := GetTokenInfo(declCtx)
		// Add the first parameter
		param := gParser.nodes[paramKey].(*ast.Decl)
		params = append(params, *param)

		// Get the rest of the parameters
		for _, paramCtx := range paramsCtx.(*ParametersPrimeContext).AllParametersPPrime() {
			declCtx := paramCtx.GetDcl().(*DeclContext)
			_, _, paramKey := GetTokenInfo(declCtx)
			// Add the rest of the parameters
			param := gParser.nodes[paramKey].(*ast.Decl)
			params = append(params, *param)
		}
	}

	// Get the declaration of the function body
	var bodyDecls []ast.Declaration
	for _, declCtx := range c.GetDcls().(*DeclarationsContext).AllDeclaration() {
		// Get the declaration
		_, _, declKey := GetTokenInfo(declCtx)
		decl := gParser.nodes[declKey].(*ast.Declaration)
		bodyDecls = append(bodyDecls, *decl)
		// Get the rest of the declarations
		for _, moreDeclCtx := range declCtx.GetIdx().(*IdsContext).AllOtherids() {
			_, _, moreDeclKey := GetTokenInfo(moreDeclCtx)
			decl := gParser.nodes[moreDeclKey].(*ast.Declaration)
			bodyDecls = append(bodyDecls, *decl)
		}
	}

	// Get the statements of the function body
	var bodyStmts []ast.Statement
	for _, stmtCtx := range c.GetStmts().(*StatementsContext).AllStatement() {
		// Get the first statement
		_, _, stmtKey := GetTokenInfo(stmtCtx)
		stmt := gParser.nodes[stmtKey].(*ast.Statement)
		bodyStmts = append(bodyStmts, *stmt)
	}

	// Add the function node to the AST
	gParser.nodes[key] = ast.NewFunction(funcName, params, bodyDecls, bodyStmts, types.StringToType(returnType), token.NewToken(line, col))
}

// ExitStatement is called when exiting the statement production.
func (gParser *goliteParser) ExitStatement(c *StatementContext) {
	// Get the key for the statement
	// line, col, key := GetTokenInfo(c)
	// Get the statement type
	if blockStatement := c.GetBl(); blockStatement != nil {
		fmt.Println("block statement")
	} else if assignStatement := c.GetAsmt(); assignStatement != nil {
		fmt.Println("assign statement")
	} else if printStatement := c.GetPrnt(); printStatement != nil {
		fmt.Println("print statement")
	} else if delStatement := c.GetDel(); delStatement != nil {
		fmt.Println("delete statement")
	} else if rdStatement := c.GetRd(); rdStatement != nil {
		fmt.Println("read statement")
	} else if condStatement := c.GetCond(); condStatement != nil {
		fmt.Println("conditional statement")
	} else if lpStatement := c.GetLp(); lpStatement != nil {
		fmt.Println("loop statement")
	} else if returnStatement := c.GetRet(); returnStatement != nil {
		fmt.Println("return statement")
	} else if invokeStatement := c.GetInvoke(); invokeStatement != nil {
		fmt.Println("invoke statement")
	} else {
		panic(fmt.Sprintf("Unknown statement - %s", c.GetText()))
	}

	// // Get the statement type
	// stmtType := c.GetStmtType().GetText()
	// // Get the statement expression
	// exprCtx := c.GetExpr().(*ExpressionContext)
	// _, _, exprKey := GetTokenInfo(exprCtx)
	// expr := gParser.nodes[exprKey].(*ast.Expression)

	// // Create the statement
	// gParser.nodes[key] = ast.NewStatement(stmtType, *expr, token.NewToken(line, col))
}

// ExitBlock is called when exiting the block production.
func (gParser *goliteParser) ExitBlock(c *BlockContext) {
	// Get the key for the block
	line, col, key := GetTokenInfo(c)

	// Get the statements of the block
	var stmts []ast.Statement
	for _, stmtCtx := range c.GetStmts().(*StatementsContext).AllStatement() {
		// Get the statement key
		_, _, stmtKey := GetTokenInfo(stmtCtx)
		// Get the statement
		stmt := gParser.nodes[stmtKey].(*ast.Statement)
		// Add the statement to the list of statements
		stmts = append(stmts, *stmt)
	}

	// Add the block node to the AST
	gParser.nodes[key] = ast.NewBlock(stmts, token.NewToken(line, col))

}

// ExitAssignment is called when exiting the assignment production.
func (gParser *goliteParser) ExitAssignment(c *AssignmentContext) {
	// Get the key for the assignment
	line, col, key := GetTokenInfo(c)

	// Get the left hand side of the assignment
	lhsCtx := c.GetLval().(*LValueContext)
	_, _, lhsKey := GetTokenInfo(lhsCtx)
	lhs := gParser.nodes[lhsKey].(*ast.LValue)

	// Get the right hand side of the assignment
	exprCtx := c.GetExpr()
	_, _, exprKey := GetTokenInfo(exprCtx)
	expr := gParser.nodes[exprKey].(*ast.Expression)

	// Create the assignment
	gParser.nodes[key] = ast.NewAssignment(*lhs, *expr, token.NewToken(line, col))
}

// ExitLValue is called when exiting the lValue production.
func (gParser *goliteParser) ExitLValue(c *LValueContext) {
	// Get the key for the lValue
	line, col, key := GetTokenInfo(c)

	varName := c.GetId().GetText()

	// Get the lValue fields
	var fields []string
	for _, fieldCtx := range c.AllLValuePrime() {
		field := fieldCtx.GetId().GetText()
		fields = append(fields, field)
	}

	// Create the lValue
	// Initialize the type to nil
	// Type will be set in the semantic analysis
	gParser.nodes[key] = ast.NewLValue(varName, fields, types.StringToType("nil"), token.NewToken(line, col))
}

// ExitPrint is called when exiting the print production.
func (gParser *goliteParser) ExitPrint(c *PrintContext) {
	// Get the key for the print statement
	line, col, key := GetTokenInfo(c)

	// Get the format string
	formatString := c.GetStr().GetText()

	// Get the expressions to print
	var exprs []ast.Expression
	for _, printPrimeCtx := range c.AllPrintPrime() {
		exprCtx := printPrimeCtx.(*PrintPrimeContext).GetExpr()
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr := gParser.nodes[exprKey].(*ast.Expression)
		exprs = append(exprs, *expr)
	}

	// Add the print node to the AST
	gParser.nodes[key] = ast.NewPrint(formatString, exprs, token.NewToken(line, col))
}

// ExitDelete is called when exiting the delete production.
func (gParser *goliteParser) ExitDelete(c *DeleteContext) {
	// Get the key for the delete statement
	line, col, key := GetTokenInfo(c)

	// Get the expression
	exprCtx := c.GetExpr().(*ExpressionContext)
	fmt.Println(exprCtx.GetText())
	_, _, exprKey := GetTokenInfo(exprCtx)
	expr := gParser.nodes[exprKey].(*ast.Expression)
	// Create the delete statement
	gParser.nodes[key] = ast.NewDelete(*expr, token.NewToken(line, col))
}

// ExitRead is called when exiting the read production.
func (gParser *goliteParser) ExitRead(c *ReadContext) {
	// Get the key for the read statement
	line, col, key := GetTokenInfo(c)

	// Get the lValue
	lvalCtx := c.GetLval().(*LValueContext)
	_, _, lvalKey := GetTokenInfo(lvalCtx)
	lval := gParser.nodes[lvalKey].(*ast.LValue)

	// Create the read statement
	gParser.nodes[key] = ast.NewRead(*lval, token.NewToken(line, col))
}

// ExitConditional is called when exiting the conditional production.
func (gParser *goliteParser) ExitConditional(c *ConditionalContext) {
	// Get the key for the conditional
	line, col, key := GetTokenInfo(c)

	// Get the condition
	condCtx := c.GetExpr().(*ExpressionContext)
	_, _, condKey := GetTokenInfo(condCtx)
	cond := gParser.nodes[condKey].(*ast.Expression)

	// Get the true block
	trueBlockCtx := c.GetBl().(*BlockContext)
	_, _, trueBlockKey := GetTokenInfo(trueBlockCtx)
	trueBlock := gParser.nodes[trueBlockKey].(*ast.Block)

	var falseBlock *ast.Block

	// Check for the false block
	if elseCtx := c.GetThen(); elseCtx != nil {
		// Get the false block
		falseBlockCtx := elseCtx.(*ConditionalPrimeContext).GetBl().(*BlockContext)
		_, _, falseBlockKey := GetTokenInfo(falseBlockCtx)
		falseBlock = gParser.nodes[falseBlockKey].(*ast.Block)
	}

	// Create the conditional
	gParser.nodes[key] = ast.NewConditional(*cond, trueBlock, falseBlock, token.NewToken(line, col))

}

// ExitLoop is called when exiting the loop production.
func (gParser *goliteParser) ExitLoop(c *LoopContext) {
	// Get the key for the loop
	line, col, key := GetTokenInfo(c)

	// Get the condition
	condCtx := c.GetExpr().(*ExpressionContext)
	_, _, condKey := GetTokenInfo(condCtx)
	cond := gParser.nodes[condKey].(*ast.Expression)

	// Get the block
	blockCtx := c.GetBl().(*BlockContext)
	_, _, blockKey := GetTokenInfo(blockCtx)
	block := gParser.nodes[blockKey].(*ast.Block)

	// Create the loop
	gParser.nodes[key] = ast.NewLoop(*cond, *block, token.NewToken(line, col))

}

// ExitReturnRule is called when exiting the returnRule production.
func (gParser *goliteParser) ExitReturnRule(c *ReturnRuleContext) {
	// Get the key for the return statement
	line, col, key := GetTokenInfo(c)

	// Get the expression
	exprCtx := c.GetExpr().(*ExpressionContext)
	_, _, exprKey := GetTokenInfo(exprCtx)
	expr := gParser.nodes[exprKey].(*ast.Expression)

	// Create the return statement
	gParser.nodes[key] = ast.NewReturn(expr, token.NewToken(line, col))
}

// ExitInvocation is called when exiting the invocation production.
func (gParser *goliteParser) ExitInvocation(c *InvocationContext) {
	// Get the key for the invocation
	line, col, key := GetTokenInfo(c)

	// Get the function name
	funcName := c.GetId().GetText()

	// Get the arguments
	var args []ast.Expression

	if argsCtx := c.GetArgs().(*ArgumentsContext).GetArgs(); argsCtx != nil {
		// First argument
		exprCtx := argsCtx.GetExpr().(*ExpressionContext)
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr := gParser.nodes[exprKey].(*ast.Expression)
		args = append(args, *expr)

		// Get the rest of the arguments (if any)
		for _, exprWrapperCtx := range argsCtx.(*ArgumentsPrimeContext).AllArgumentsPrimePrime() {
			exprCtx := exprWrapperCtx.(*ArgumentsPrimePrimeContext).GetExpr().(*ExpressionContext)
			_, _, exprKey := GetTokenInfo(exprCtx)
			expr := gParser.nodes[exprKey].(*ast.Expression)
			args = append(args, *expr)
		}
	}

	// Create the invocation
	gParser.nodes[key] = ast.NewInvocation(funcName, args, token.NewToken(line, col))
}

// ExitExpression is called when exiting the expression production.
func (gParser *goliteParser) ExitExpression(c *ExpressionContext) {
	// Get the key for the expression
	line, col, key := GetTokenInfo(c)

	// Get the left expression
	leftExprCtx := c.GetBt()
	_, _, leftExprKey := GetTokenInfo(leftExprCtx)
	leftExpr := gParser.nodes[leftExprKey].(*ast.Expression)
	allRightExpr := c.AllExpressionPrime()
	if len(allRightExpr) != 0 {
		for _, rightExprCtx := range allRightExpr {
			// Get the right expression
			_, _, rightExprKey := GetTokenInfo(rightExprCtx)
			rightExpr := gParser.nodes[rightExprKey].(*ast.Expression)

			// Get the operator
			op := rightExprCtx.(*ExpressionPrimeContext).GetOp().GetText()
			opTerm := ast.StrToOp(op)

			// Create the expression
			mergedExpr := ast.NewBinOp(leftExpr, &opTerm, rightExpr, token.NewToken(line, col))
			leftExpr = mergedExpr.(*ast.Expression)
		}
	}
	gParser.nodes[key] = leftExpr

}

// ExitBoolTerm is called when exiting the boolTerm production.
func (gParser *goliteParser) ExitBoolTerm(c *BoolTermContext) {
	// Get the key for the bool term
	line, col, key := GetTokenInfo(c)

	// Get the first expression
	exprCtx := c.GetEq()
	_, _, exprKey := GetTokenInfo(exprCtx)
	expr := gParser.nodes[exprKey].(*ast.Expression)
	allOtherExpr := c.AllBoolTermPrime()
	if len(allOtherExpr) != 0 {
		// Get the rest of the expressions
		for _, exprWrapperCtx := range allOtherExpr {
			exprCtx := exprWrapperCtx.(*BoolTermPrimeContext).GetEq()
			_, _, exprKey := GetTokenInfo(exprCtx)
			otherExpr := gParser.nodes[exprKey].(*ast.Expression)

			// Get the operator
			op := exprWrapperCtx.(*BoolTermPrimeContext).GetOp().GetText()
			opTerm := ast.StrToOp(op)

			// Create the new expression
			mergedTerm := ast.NewBinOp(expr, &opTerm, otherExpr, token.NewToken(line, col))
			expr = mergedTerm.(*ast.Expression)
		}
	}
	gParser.nodes[key] = expr
}

// ExitEqualTerm is called when exiting the equalTerm production.
func (gParser *goliteParser) ExitEqualTerm(c *EqualTermContext) {
	// Get the key for the equal term
	line, col, key := GetTokenInfo(c)

	// Get the left expression
	leftExprCtx := c.GetRt()
	_, _, leftExprKey := GetTokenInfo(leftExprCtx)
	leftExpr := gParser.nodes[leftExprKey].(*ast.Expression)
	allRightExpr := c.AllEqualTermPrime()
	if len(allRightExpr) != 0 {
		for _, rightExprCtx := range allRightExpr {
			rightExprPrimeCtx := rightExprCtx.(*EqualTermPrimeContext)

			// Get the right expression
			_, _, rightExprKey := GetTokenInfo(rightExprPrimeCtx.GetRt())
			rightExpr := gParser.nodes[rightExprKey].(*ast.Expression)

			// Get the operator
			op := rightExprCtx.GetOp().GetText()
			opTerm := ast.StrToOp(op)

			// Create the equal term
			mergedTerm := ast.NewBinOp(leftExpr, &opTerm, rightExpr, token.NewToken(line, col))
			leftExpr = mergedTerm.(*ast.Expression)
		}
	}
	gParser.nodes[key] = leftExpr

}

// ExitRelationTerm is called when exiting the relationTerm production.
func (gParser *goliteParser) ExitRelationTerm(c *RelationTermContext) {
	// Get the key for the relation term
	line, col, key := GetTokenInfo(c)

	// Get the relationTerm
	relTermCtx := c.GetSt()
	_, _, relTermKey := GetTokenInfo(relTermCtx)
	relTerm := gParser.nodes[relTermKey].(*ast.Expression)
	allRelationTermCtx := c.AllRelationTermPrime()
	if len(allRelationTermCtx) != 0 {
		for _, relTermPrimeCtx := range allRelationTermCtx {
			relTermPrime := relTermPrimeCtx.(*RelationTermPrimeContext)
			// Get the operator
			op := relTermPrime.GetOp().GetText()
			// Get the relationTerm
			relTermCtx := relTermPrime.GetSt()
			_, _, relTermKey := GetTokenInfo(relTermCtx)
			otherRelTerm := gParser.nodes[relTermKey].(*ast.Expression)
			// Create the relation term
			opTerm := ast.StrToOp(op)
			mergedTerm := ast.NewBinOp(relTerm, &opTerm, otherRelTerm, token.NewToken(line, col))
			relTerm = mergedTerm.(*ast.Expression)
		}
	}
	gParser.nodes[key] = relTerm

}

// ExitSimpleTerm is called when exiting the simpleTerm production.
func (gParser *goliteParser) ExitSimpleTerm(c *SimpleTermContext) {
	// Get the key for the term
	line, col, key := GetTokenInfo(c)

	// Get the term
	termCtx := c.GetT().(*TermContext)
	_, _, termKey := GetTokenInfo(termCtx)
	term := gParser.nodes[termKey].(*ast.Expression)
	allTermPrimeCtx := c.AllSimpleTermPrime()
	if len(allTermPrimeCtx) != 0 {
		// Get the rest of the terms
		for _, termPrimeCtx := range allTermPrimeCtx {
			ctx := termPrimeCtx.(*SimpleTermPrimeContext)
			// Get the operator
			op := ctx.GetOp().GetText()

			_, _, termPrimeKey := GetTokenInfo(ctx.GetT().(*TermContext))
			termPrime := gParser.nodes[termPrimeKey].(*ast.Expression)
			opTerm := ast.StrToOp(op)
			mergedTerm := ast.NewBinOp(term, &opTerm, termPrime, token.NewToken(line, col))
			term = mergedTerm.(*ast.Expression)
		}
	}
	gParser.nodes[key] = term

}

// ExitTerm is called when exiting the term production.
func (gParser *goliteParser) ExitTerm(c *TermContext) {
	// Get the key for the term
	line, col, key := GetTokenInfo(c)
	// Get the unary term
	unaryTermCtx := c.GetUt().(*UnaryTermContext)
	// Get all other terms
	_, _, unaryTermKey := GetTokenInfo(unaryTermCtx)
	unaryTerm := gParser.nodes[unaryTermKey].(*ast.Expression)
	allTermPrimeCtx := c.AllTermPrime()
	if len(allTermPrimeCtx) != 0 {
		// Build nested binary terms
		for _, termPrimeWrapperCtx := range allTermPrimeCtx {
			termPrimeCtx := termPrimeWrapperCtx.(*TermPrimeContext)
			// Get the operator
			op := termPrimeCtx.GetOp().GetText()
			// Get the unary term
			unaryTermCtx := termPrimeCtx.GetUt().(*UnaryTermContext)
			_, _, unaryTermKey := GetTokenInfo(unaryTermCtx)
			unaryTerm2 := gParser.nodes[unaryTermKey].(*ast.Expression)
			// Cast to the correct
			// Create the binary term
			opTerm := ast.StrToOp(op)
			mergedTerm := ast.NewBinOp(unaryTerm, &opTerm, unaryTerm2, token.NewToken(line, col))
			unaryTerm = mergedTerm.(*ast.Expression)
		}
	}
	gParser.nodes[key] = unaryTerm

}

// ExitUnaryTerm is called when exiting the unary production.
func (gParser *goliteParser) ExitUnaryTerm(c *UnaryTermContext) {
	// Get the key for the unary expression
	line, col, key := GetTokenInfo(c)

	var op string
	var selector *ast.SelectorTerm
	if boolCtx := c.UnaryTermBool(); boolCtx != nil {
		op = boolCtx.GetOp().GetText()
		// Get the selector
		selectorCtx := boolCtx.GetSt().(*SelectorTermContext)
		_, _, selectorKey := GetTokenInfo(selectorCtx)
		selector = gParser.nodes[selectorKey].(*ast.SelectorTerm)
	} else if intCtx := c.UnaryTermInt(); intCtx != nil {
		op = intCtx.GetText()
		// Get the selector
		selectorCtx := intCtx.GetSt().(*SelectorTermContext)
		_, _, selectorKey := GetTokenInfo(selectorCtx)
		selector = gParser.nodes[selectorKey].(*ast.SelectorTerm)
	} else if selectCtx := c.SelectorTerm(); selectCtx != nil {
		op = ""
		// Get the selector
		selectorCtx := selectCtx.(*SelectorTermContext)
		_, _, selectorKey := GetTokenInfo(selectorCtx)
		selector = gParser.nodes[selectorKey].(*ast.SelectorTerm)
	} else {
		panic("Invalid unary term")
	}

	// Get the operator
	var optPtr *ast.Operator
	if op != "" {
		operator := ast.StrToOp(op)
		optPtr = &operator
	} else {
		// Create a nil pointer
		optPtr = nil
	}

	// Create the unary expression
	gParser.nodes[key] = ast.NewBinOp(nil, optPtr, selector, token.NewToken(line, col))
}

// ExitSelectorTerm is called when exiting the selectorTerm production.
func (gParser *goliteParser) ExitSelectorTerm(c *SelectorTermContext) {
	// Get the key for the selector
	line, col, key := GetTokenInfo(c)

	// Get the factor expression
	_, _, factorKey := GetTokenInfo(c.Factor())
	factor := gParser.nodes[factorKey].(*ast.Expression)

	// Get the fields
	var fields []string
	for _, fieldCtx := range c.AllSelectorTermPrime() {
		field := fieldCtx.(*SelectorTermPrimeContext).GetId().GetText()
		fields = append(fields, field)
	}

	// Create the selector term
	gParser.nodes[key] = ast.NewSelectorTerm(*factor, fields, types.StringToType("nil"), token.NewToken(line, col))
}

// ExitFactor is called when exiting the factor production.
func (gParser *goliteParser) ExitFactor(c *FactorContext) {
	line, col, key := GetTokenInfo(c)
	newToken := token.NewToken(line, col)

	if intFactor := c.INT_LIT(); intFactor != nil {
		// Check if it is an int literal
		// Convert to int
		intValue, err := strconv.Atoi(intFactor.GetText())
		// Check if the conversion was successful
		if err != nil {
			panic(err)
		}
		// Add the node to the ast
		gParser.nodes[key] = ast.NewIntLit(int64(intValue), newToken)
	} else if boolFactor := c.BOOL_LIT(); boolFactor != nil {
		// Check if it is a boolean literal
		// Convert to specific boolean value
		var boolVal bool
		if boolFactor.GetText() == "true" {
			boolVal = true
		} else if boolFactor.GetText() == "false" {
			boolVal = false
		} else {
			panic("Invalid boolean value")
		}
		// Add the node to the ast
		gParser.nodes[key] = ast.NewBoolLit(boolVal, newToken)
	} else if nilFactor := c.NIL_LIT(); nilFactor != nil {
		// Nil literal
		gParser.nodes[key] = ast.NewNilLit(newToken)
	} else if funcCallFactor := c.Functioncall(); funcCallFactor != nil {
		// Function call
		// Get the key for the function call
		_, _, funcCallKey := GetTokenInfo(funcCallFactor)
		// Get the function call
		funcCall := gParser.nodes[funcCallKey].(*ast.Invocation)
		// Add node to the ast
		gParser.nodes[key] = funcCall
	} else if allocationFactor := c.Allocation(); allocationFactor != nil {
		// Allocation
		// Add the node to the ast
		gParser.nodes[key] = ast.NewAllocate(allocationFactor.GetKey().GetText(), newToken)
	} else if subfactorFactor := c.Subfactor(); subfactorFactor != nil {
		// Subfactor
		// Get the expression
		// exprCtx := subfactorFactor.GetExpr().(*ExpressionContext)
		// _, _, exprKey := GetTokenInfo(exprCtx)
		// Add the node to the ast
		// gParser.nodes[key] =
	} else {
		panic("Invalid factor in factor production rule")
	}
}

// ExitFunctioncall is called when exiting the functionCall production.
func (gParser *goliteParser) ExitFunctioncall(c *FunctioncallContext) {
	// Get the key for the invocation
	line, col, key := GetTokenInfo(c)

	// Get the function name
	funcName := c.GetId().GetText()

	// Get the arguments
	var args []ast.Expression
	// If there are arguments
	if argsCtx := c.GetArgs().(*ArgumentsContext).GetArgs(); argsCtx != nil {
		// First argument
		exprCtx := argsCtx.GetExpr().(*ExpressionContext)
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr := gParser.nodes[exprKey].(*ast.Expression)
		args = append(args, *expr)

		// Get the rest of the arguments (if any)
		for _, exprWrapperCtx := range argsCtx.(*ArgumentsPrimeContext).AllArgumentsPrimePrime() {
			exprCtx := exprWrapperCtx.(*ArgumentsPrimePrimeContext).GetExpr().(*ExpressionContext)
			_, _, exprKey := GetTokenInfo(exprCtx)
			expr := gParser.nodes[exprKey].(*ast.Expression)
			args = append(args, *expr)
		}
	}

	// Create the invocation
	gParser.nodes[key] = ast.NewInvocation(funcName, args, token.NewToken(line, col))
}
