package ast

import (
	"bytes"
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
func (p *Program) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	// Struct types
	for _, strct := range p.structTypes {
		errors = strct.TypeCheck(errors, tables)
	}

	// Declarations
	for _, decl := range p.declarations {
		errors = decl.TypeCheck(errors, tables)
	}

	// Functions
	for _, fn := range p.funcs {
		errors = fn.TypeCheck(errors, tables)
	}

	return errors
}
