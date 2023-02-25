package ast

import (
	"bytes"
	"golite/llvm"
	"golite/runtime"
	st "golite/symboltable"
	"golite/token"
)

// Program production rule for the AST
type Program struct {
	*token.Token
	structTypes  []StructTypeStatement
	declarations []DeclarationStatement
	funcs        []FunctionStatement
}

// Create a new program node
func NewProgram(structTypes []StructTypeStatement, declarations []DeclarationStatement, funcs []FunctionStatement, token *token.Token) *Program {
	return &Program{token, structTypes, declarations, funcs}
}

// String representation of the program node
func (p *Program) String() string {
	var out bytes.Buffer

	// Struct types
	for _, strct := range p.structTypes {
		out.WriteString(strct.String())
		out.WriteString("\n")
	}

	// Declarations
	for _, decl := range p.declarations {
		out.WriteString(decl.String())
		out.WriteString("\n")
	}

	// Functions
	for _, fn := range p.funcs {
		out.WriteString(fn.String())
		out.WriteString("\n")
	}

	return out.String()
}

// Build the symbol table for the program
func (p *Program) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Struct types
	for _, strct := range p.structTypes {
		errors = strct.BuildSymbolTable(tables, errors)
	}

	// Declarations
	for _, decl := range p.declarations {
		if !tables.Globals.Insert(decl.GetVariable(), &st.VarEntry{Name: decl.GetVariable(), Ty: decl.GetType(), Scope: st.GLOBAL}) {
			errors = append(errors, NewSemanticAnalysisError("Variable '"+decl.GetVariable()+"' redeclared.", "redeclaration", decl.GetToken()))
		}
	}

	// Functions
	for _, fn := range p.funcs {
		errors = fn.BuildSymbolTable(tables, errors)
	}

	return errors
}

// Type Check using the symbol tables
func (p *Program) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Struct types
	for _, strct := range p.structTypes {
		errors = strct.TypeCheck(errors, tables, funcEntry)
	}

	// Declarations
	for _, decl := range p.declarations {
		errors = decl.TypeCheck(errors, tables, funcEntry)
	}

	// Functions
	for _, fn := range p.funcs {
		errors = fn.TypeCheck(errors, tables, funcEntry)
	}

	return errors
}

// Control flow check
func (p *Program) ControlFlowCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Functions
	for _, fn := range p.funcs {
		errors = fn.ControlFlowCheck(errors, tables, funcEntry)
	}

	return errors
}

// Tralsate the program to LLVM IR
func (p *Program) ToLLVM(tables *st.SymbolTables) *llvm.Program {
	var globalDecls []llvm.Decl
	var globalDecl *llvm.Decl

	// Translate the struct definitions to LLVM IR
	var structDecls []llvm.StructDecl
	var structDecl *llvm.StructDecl
	for _, strct := range p.structTypes {
		structDecl = strct.ToLLVM(tables)
		structDecls = append(structDecls, *structDecl)

		// Add the global null value declarations for the structs
		globalDecl = llvm.NewDecl("struct."+strct.GetName(), "struct."+strct.GetName(), "null", true)
		globalDecls = append(globalDecls, *globalDecl)
	}

	// Translate the declarations to LLVM IR
	// Add the global declarations for the variables
	for _, decl := range p.declarations {
		globalDecl = decl.ToLLVM(tables)
		globalDecls = append(globalDecls, *globalDecl)
	}

	// Translate the functions to LLVM IR
	var functionDecls []llvm.FunctionDecl
	var functionDecl *llvm.FunctionDecl
	for _, fn := range p.funcs {
		tables, functionDecl = fn.ToLLVM(tables)
		functionDecls = append(functionDecls, *functionDecl)
	}

	// Get the C runtime functions
	runtimeDecls := make([]llvm.RuntimeDecl, 0)
	cRuntime := runtime.NewRuntime()

	for _, cFuncName := range cRuntime.GetCFuncs() {
		runtimeDecls = append(runtimeDecls, *llvm.NewRuntimeDecl(cFuncName, cRuntime.GetCFunc(cFuncName)))
	}

	// Create the LLVM program
	llvm := llvm.NewProgram(structDecls, globalDecls, functionDecls, runtimeDecls)

	return llvm
}
