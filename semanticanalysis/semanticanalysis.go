package semanticanalysis

import (
	"fmt"
	"golite/ast"
	st "golite/symboltable"
)

// Check if there are any errors in the error list
func hasErrors(errors []*ast.SemanticAnalysisError) bool {
	return len(errors) > 0
}

// Report the errors to the user
func reportErrors(errors []*ast.SemanticAnalysisError) {
	for _, err := range errors {
		fmt.Println(err)
	}
}

func Execute(program *ast.Program) *st.SymbolTables {

	//Define the compiler symbol tables
	tables := st.NewSymbolTables()

	// Collect all the errors
	errors := make([]ast.SemanticAnalysisError, 0)

	//First Build the Symbol Table(s) for all global declarations
	errors = program.BuildSymbolTable(tables, errors)

	//Second perform type checking.
	errors = program.TypeCheck(errors, tables)
	if hasErrors(errors) {
		reportErrors(errors)
		return nil
	}

	return tables
}
