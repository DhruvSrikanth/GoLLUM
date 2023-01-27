// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseGoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type BaseGoliteParserListener struct{}

var _ GoliteParserListener = &BaseGoliteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoliteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoliteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoliteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoliteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoliteParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoliteParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseGoliteParserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseGoliteParserListener) ExitTypes(ctx *TypesContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseGoliteParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseGoliteParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseGoliteParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseGoliteParserListener) ExitFields(ctx *FieldsContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseGoliteParserListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseGoliteParserListener) ExitDecl(ctx *DeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoliteParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoliteParserListener) ExitType(ctx *TypeContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseGoliteParserListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseGoliteParserListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGoliteParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGoliteParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterIds is called when production ids is entered.
func (s *BaseGoliteParserListener) EnterIds(ctx *IdsContext) {}

// ExitIds is called when production ids is exited.
func (s *BaseGoliteParserListener) ExitIds(ctx *IdsContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseGoliteParserListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseGoliteParserListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseGoliteParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseGoliteParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGoliteParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGoliteParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseGoliteParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseGoliteParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseGoliteParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseGoliteParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoliteParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoliteParserListener) ExitStatement(ctx *StatementContext) {}

// EnterRead is called when production read is entered.
func (s *BaseGoliteParserListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseGoliteParserListener) ExitRead(ctx *ReadContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoliteParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoliteParserListener) ExitBlock(ctx *BlockContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseGoliteParserListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseGoliteParserListener) ExitDelete(ctx *DeleteContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoliteParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoliteParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseGoliteParserListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseGoliteParserListener) ExitPrint(ctx *PrintContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseGoliteParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseGoliteParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseGoliteParserListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseGoliteParserListener) ExitLoop(ctx *LoopContext) {}

// EnterReturnRule is called when production returnRule is entered.
func (s *BaseGoliteParserListener) EnterReturnRule(ctx *ReturnRuleContext) {}

// ExitReturnRule is called when production returnRule is exited.
func (s *BaseGoliteParserListener) ExitReturnRule(ctx *ReturnRuleContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseGoliteParserListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseGoliteParserListener) ExitInvocation(ctx *InvocationContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoliteParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoliteParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLValue is called when production lValue is entered.
func (s *BaseGoliteParserListener) EnterLValue(ctx *LValueContext) {}

// ExitLValue is called when production lValue is exited.
func (s *BaseGoliteParserListener) ExitLValue(ctx *LValueContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoliteParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoliteParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBoolTerm is called when production boolTerm is entered.
func (s *BaseGoliteParserListener) EnterBoolTerm(ctx *BoolTermContext) {}

// ExitBoolTerm is called when production boolTerm is exited.
func (s *BaseGoliteParserListener) ExitBoolTerm(ctx *BoolTermContext) {}

// EnterEqualTerm is called when production equalTerm is entered.
func (s *BaseGoliteParserListener) EnterEqualTerm(ctx *EqualTermContext) {}

// ExitEqualTerm is called when production equalTerm is exited.
func (s *BaseGoliteParserListener) ExitEqualTerm(ctx *EqualTermContext) {}

// EnterRelationTerm is called when production relationTerm is entered.
func (s *BaseGoliteParserListener) EnterRelationTerm(ctx *RelationTermContext) {}

// ExitRelationTerm is called when production relationTerm is exited.
func (s *BaseGoliteParserListener) ExitRelationTerm(ctx *RelationTermContext) {}

// EnterSimpleTerm is called when production simpleTerm is entered.
func (s *BaseGoliteParserListener) EnterSimpleTerm(ctx *SimpleTermContext) {}

// ExitSimpleTerm is called when production simpleTerm is exited.
func (s *BaseGoliteParserListener) ExitSimpleTerm(ctx *SimpleTermContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGoliteParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGoliteParserListener) ExitTerm(ctx *TermContext) {}

// EnterUnaryTerm is called when production unaryTerm is entered.
func (s *BaseGoliteParserListener) EnterUnaryTerm(ctx *UnaryTermContext) {}

// ExitUnaryTerm is called when production unaryTerm is exited.
func (s *BaseGoliteParserListener) ExitUnaryTerm(ctx *UnaryTermContext) {}

// EnterSelectorTerm is called when production selectorTerm is entered.
func (s *BaseGoliteParserListener) EnterSelectorTerm(ctx *SelectorTermContext) {}

// ExitSelectorTerm is called when production selectorTerm is exited.
func (s *BaseGoliteParserListener) ExitSelectorTerm(ctx *SelectorTermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseGoliteParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseGoliteParserListener) ExitFactor(ctx *FactorContext) {}
