package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Program production rule for the AST
type Program struct {
	*token.Token
	structTypes  []StructTypeList
	declarations []DeclarationList
	funcs        []FunctionList
}

// Create a new program node
func NewProgram(structTypes []StructTypeList, declarations []DeclarationList, funcs []FunctionList, token *token.Token) *Program {
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
func (p *Program) BuildSymbolTable(tables *st.SymbolTables) {
	// Struct types
	for _, strct := range p.structTypes {
		strct.BuildSymbolTable(tables)
	}

	// Declarations
	for _, decl := range p.declarations {
		decl.BuildSymbolTable(tables)
	}

	// Functions
	for _, fn := range p.funcs {
		fn.BuildSymbolTable(tables)
	}
}

// Type Check using the symbol tables
func (p *Program) TypeCheck(errors []string, tables *st.SymbolTables) []string {
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
