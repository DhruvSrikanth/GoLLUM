package ast

import st "golite/symboltable"

// FunctionStatement is an interface for all functions
type FunctionStatement interface {
	Statement
	ControlFlowCheck([]*SemanticAnalysisError, *st.SymbolTables, *st.FuncEntry) []*SemanticAnalysisError
}
