package ast

import (
	"golite/llvm"
	st "golite/symboltable"
)

// FunctionStatement is an interface for all functions
type FunctionStatement interface {
	Statement
	ControlFlowCheck([]*SemanticAnalysisError, *st.SymbolTables, *st.FuncEntry) []*SemanticAnalysisError
	ToLLVM(*st.SymbolTables, []llvm.ConstantDecl) (*st.SymbolTables, *llvm.FunctionDecl, []llvm.ConstantDecl)
}
