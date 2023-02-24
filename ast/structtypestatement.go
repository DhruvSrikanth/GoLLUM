package ast

import (
	"golite/llvm"
	st "golite/symboltable"
)

// StructTypeStatement is an interface that all struct types implement
type StructTypeStatement interface {
	Statement
	ToLLVM(*st.SymbolTables) (*st.SymbolTables, *llvm.StructDecl)
}
