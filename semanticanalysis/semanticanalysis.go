package semanticanalysis

import (
	"golite/ast"
	st "golite/symboltable"
)

// Check if there are any errors in the error list
func hasErrors(errors []string) bool {
	return len(errors) > 0
}

// Report the errors to the user
func reportErrors(errors []string) {
	for _, err := range errors {
		println(err)
	}
}

func Execute(program *ast.Program) *st.SymbolTables {

	//Define the compiler symbol tables
	tables := st.NewSymbolTables()

	//First Build the Symbol Table(s) for all global declarations
	program.BuildSymbolTable(tables)

	//Second perform type checking.
	errors := make([]string, 0)
	errors = program.TypeCheck(errors, tables)
	if hasErrors(errors) {
		reportErrors(errors)
		return nil
	}

	return tables
}
