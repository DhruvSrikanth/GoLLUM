package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

// Struct declaration
type StructDecl struct {
	*token.Token
	name  string // Name of the struct
	decls []Decl // Declarations in the struct
}

// Returns an instance to a new struct declaration
func NewStructDecl(token *token.Token, name string, decls []Decl) *StructDecl {
	return &StructDecl{token, name, decls}
}

// String representation of the struct declaration
func (s *StructDecl) String() string {
	var out bytes.Buffer

	out.WriteString("struct ")
	out.WriteString(s.name)
	out.WriteString(" {\n")
	for _, decl := range s.decls {
		out.WriteString(decl.String())
		out.WriteString("\n")
	}
	out.WriteString("}")

	return out.String()
}

// Build the symbol table for the struct declaration
func (s *StructDecl) BuildSymbolTable(tables *st.SymbolTables) {
	// Create fields for the struct entry
	fields := make([]*st.VarEntry, 0)
	for _, decl := range s.decls {
		fields = append(fields, &st.VarEntry{decl.variable, decl.ty, st.GLOBAL})
	}

	// Add the struct to the symbol table
	tables.Structs.Insert(s.name, &st.StructEntry{s.name, fields})
}

// Type check the struct declaration
func (d *StructDecl) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	// Nothing to type check because only lvalues are allowed in struct declarations
	return errors
}
