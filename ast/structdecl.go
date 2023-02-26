package ast

import (
	"bytes"
	"golite/llvm"
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
func NewStructDecl(name string, decls []Decl, token *token.Token) *StructDecl {
	return &StructDecl{token, name, decls}
}

// Get the name of the struct
func (d *StructDecl) GetName() string {
	return d.name
}

// String representation of the struct declaration
func (s *StructDecl) String() string {
	var out bytes.Buffer

	out.WriteString("type ")
	out.WriteString(s.name)
	out.WriteString(" struct {\n")
	for _, decl := range s.decls {
		out.WriteString(decl.String())
		out.WriteString(";")
		out.WriteString("\n")
	}
	out.WriteString("}")

	return out.String()
}

// Build the symbol table for the struct declaration
func (s *StructDecl) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Create fields for the struct entry
	fields := make([]*st.FieldEntry, 0)
	for _, decl := range s.decls {
		fields = append(fields, &st.FieldEntry{Name: decl.variable, Ty: decl.ty})
	}

	// Add the struct to the symbol table
	if !tables.Structs.Insert(s.name, &st.StructEntry{Name: s.name, Fields: fields}) {
		errors = append(errors, NewSemanticAnalysisError("Struct '"+s.name+"'redeclared.", "redeclaration", s.Token))
	}

	return errors
}

// Type check the struct declaration
func (d *StructDecl) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Nothing to type check because only lvalues are allowed in struct declarations
	for _, decl := range d.decls {
		errors = decl.TypeCheck(errors, tables, funcEntry)
	}
	return errors
}

// Control flow analysis for the struct declaration
func (d *StructDecl) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the struct declaration node does not have any control flow
	return errors, false
}

// Translate the struct declaration to LLVM IR
func (d *StructDecl) ToLLVM(tables *st.SymbolTables) *llvm.StructDecl {
	// Get the struct entry
	structEntry := tables.Structs.Contains(d.name)

	// Get the LLVM representation of the struct fields
	fields := make([]string, 0)
	for _, field := range structEntry.Fields {
		fields = append(fields, field.LlvmType)
	}

	// Create the struct declaration
	return llvm.NewStructDecl("struct."+d.name, fields)
}

// Translate the structdecl node to LLVM IR
func (d *StructDecl) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry) []*llvm.BasicBlock {
	return blocks
}
