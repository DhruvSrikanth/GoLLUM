// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// GoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type GoliteParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterMorefields is called when entering the morefields production.
	EnterMorefields(c *MorefieldsContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterIds is called when entering the ids production.
	EnterIds(c *IdsContext)

	// EnterOtherids is called when entering the otherids production.
	EnterOtherids(c *OtheridsContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParametersPrime is called when entering the parametersPrime production.
	EnterParametersPrime(c *ParametersPrimeContext)

	// EnterParametersPPrime is called when entering the parametersPPrime production.
	EnterParametersPPrime(c *ParametersPPrimeContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPrintPrime is called when entering the printPrime production.
	EnterPrintPrime(c *PrintPrimeContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterConditionalPrime is called when entering the conditionalPrime production.
	EnterConditionalPrime(c *ConditionalPrimeContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterReturnRule is called when entering the returnRule production.
	EnterReturnRule(c *ReturnRuleContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentsPrime is called when entering the argumentsPrime production.
	EnterArgumentsPrime(c *ArgumentsPrimeContext)

	// EnterArgumentsPrimePrime is called when entering the argumentsPrimePrime production.
	EnterArgumentsPrimePrime(c *ArgumentsPrimePrimeContext)

	// EnterLValue is called when entering the lValue production.
	EnterLValue(c *LValueContext)

	// EnterLValuePrime is called when entering the lValuePrime production.
	EnterLValuePrime(c *LValuePrimeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionPrime is called when entering the expressionPrime production.
	EnterExpressionPrime(c *ExpressionPrimeContext)

	// EnterBoolTerm is called when entering the boolTerm production.
	EnterBoolTerm(c *BoolTermContext)

	// EnterBoolTermPrime is called when entering the boolTermPrime production.
	EnterBoolTermPrime(c *BoolTermPrimeContext)

	// EnterEqualTerm is called when entering the equalTerm production.
	EnterEqualTerm(c *EqualTermContext)

	// EnterEqualTermPrime is called when entering the equalTermPrime production.
	EnterEqualTermPrime(c *EqualTermPrimeContext)

	// EnterRelationTerm is called when entering the relationTerm production.
	EnterRelationTerm(c *RelationTermContext)

	// EnterRelationTermPrime is called when entering the relationTermPrime production.
	EnterRelationTermPrime(c *RelationTermPrimeContext)

	// EnterSimpleTerm is called when entering the simpleTerm production.
	EnterSimpleTerm(c *SimpleTermContext)

	// EnterSimpleTermPrime is called when entering the simpleTermPrime production.
	EnterSimpleTermPrime(c *SimpleTermPrimeContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTermPrime is called when entering the termPrime production.
	EnterTermPrime(c *TermPrimeContext)

	// EnterUnaryTerm is called when entering the unaryTerm production.
	EnterUnaryTerm(c *UnaryTermContext)

	// EnterUnaryTermBool is called when entering the unaryTermBool production.
	EnterUnaryTermBool(c *UnaryTermBoolContext)

	// EnterUnaryTermInt is called when entering the unaryTermInt production.
	EnterUnaryTermInt(c *UnaryTermIntContext)

	// EnterSelectorTerm is called when entering the selectorTerm production.
	EnterSelectorTerm(c *SelectorTermContext)

	// EnterSelectorTermPrime is called when entering the selectorTermPrime production.
	EnterSelectorTermPrime(c *SelectorTermPrimeContext)

	// EnterSubfactor is called when entering the subfactor production.
	EnterSubfactor(c *SubfactorContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterAllocation is called when entering the allocation production.
	EnterAllocation(c *AllocationContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitMorefields is called when exiting the morefields production.
	ExitMorefields(c *MorefieldsContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitIds is called when exiting the ids production.
	ExitIds(c *IdsContext)

	// ExitOtherids is called when exiting the otherids production.
	ExitOtherids(c *OtheridsContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParametersPrime is called when exiting the parametersPrime production.
	ExitParametersPrime(c *ParametersPrimeContext)

	// ExitParametersPPrime is called when exiting the parametersPPrime production.
	ExitParametersPPrime(c *ParametersPPrimeContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPrintPrime is called when exiting the printPrime production.
	ExitPrintPrime(c *PrintPrimeContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitConditionalPrime is called when exiting the conditionalPrime production.
	ExitConditionalPrime(c *ConditionalPrimeContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitReturnRule is called when exiting the returnRule production.
	ExitReturnRule(c *ReturnRuleContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentsPrime is called when exiting the argumentsPrime production.
	ExitArgumentsPrime(c *ArgumentsPrimeContext)

	// ExitArgumentsPrimePrime is called when exiting the argumentsPrimePrime production.
	ExitArgumentsPrimePrime(c *ArgumentsPrimePrimeContext)

	// ExitLValue is called when exiting the lValue production.
	ExitLValue(c *LValueContext)

	// ExitLValuePrime is called when exiting the lValuePrime production.
	ExitLValuePrime(c *LValuePrimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionPrime is called when exiting the expressionPrime production.
	ExitExpressionPrime(c *ExpressionPrimeContext)

	// ExitBoolTerm is called when exiting the boolTerm production.
	ExitBoolTerm(c *BoolTermContext)

	// ExitBoolTermPrime is called when exiting the boolTermPrime production.
	ExitBoolTermPrime(c *BoolTermPrimeContext)

	// ExitEqualTerm is called when exiting the equalTerm production.
	ExitEqualTerm(c *EqualTermContext)

	// ExitEqualTermPrime is called when exiting the equalTermPrime production.
	ExitEqualTermPrime(c *EqualTermPrimeContext)

	// ExitRelationTerm is called when exiting the relationTerm production.
	ExitRelationTerm(c *RelationTermContext)

	// ExitRelationTermPrime is called when exiting the relationTermPrime production.
	ExitRelationTermPrime(c *RelationTermPrimeContext)

	// ExitSimpleTerm is called when exiting the simpleTerm production.
	ExitSimpleTerm(c *SimpleTermContext)

	// ExitSimpleTermPrime is called when exiting the simpleTermPrime production.
	ExitSimpleTermPrime(c *SimpleTermPrimeContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTermPrime is called when exiting the termPrime production.
	ExitTermPrime(c *TermPrimeContext)

	// ExitUnaryTerm is called when exiting the unaryTerm production.
	ExitUnaryTerm(c *UnaryTermContext)

	// ExitUnaryTermBool is called when exiting the unaryTermBool production.
	ExitUnaryTermBool(c *UnaryTermBoolContext)

	// ExitUnaryTermInt is called when exiting the unaryTermInt production.
	ExitUnaryTermInt(c *UnaryTermIntContext)

	// ExitSelectorTerm is called when exiting the selectorTerm production.
	ExitSelectorTerm(c *SelectorTermContext)

	// ExitSelectorTermPrime is called when exiting the selectorTermPrime production.
	ExitSelectorTermPrime(c *SelectorTermPrimeContext)

	// ExitSubfactor is called when exiting the subfactor production.
	ExitSubfactor(c *SubfactorContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitAllocation is called when exiting the allocation production.
	ExitAllocation(c *AllocationContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
