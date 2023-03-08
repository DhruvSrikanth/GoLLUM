package semanticanalysis

import (
	"fmt"
	"golite/ast"
	st "golite/symboltable"
	"golite/utils"
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

// Check for redeclaration errors
func checkRedeclaration(errors []*ast.SemanticAnalysisError, tables *st.SymbolTables) []*ast.SemanticAnalysisError {
	nameSpace := make([]string, 0)
	for _, v := range tables.Structs.GetTable() {
		nameSpace = append(nameSpace, v.Name)
	}
	for _, v := range tables.Funcs.GetTable() {
		if utils.Contains(nameSpace, v.Name) {
			err := ast.NewSemanticAnalysisError("function name used before", "redeclaration error", nil)
			errors = append(errors, err)
		}
		for _, p := range v.Parameters {
			if utils.Contains(nameSpace, p.Name) {
				err := ast.NewSemanticAnalysisError("function parameter name used before", "redeclaration error", nil)
				errors = append(errors, err)
			}
		}
		for _, varEntry := range v.Variables.GetTable() {
			if utils.Contains(nameSpace, varEntry.Name) {
				err := ast.NewSemanticAnalysisError("function variable name used before", "redeclaration error", nil)
				errors = append(errors, err)
			}
		}
		nameSpace = append(nameSpace, v.Name)
	}
	return errors
}

// Check for the existance of a main function
func checkForMain(errors []*ast.SemanticAnalysisError, tables *st.SymbolTables) []*ast.SemanticAnalysisError {
	funcNames := make([]string, 0)
	for _, v := range tables.Funcs.GetTable() {
		funcNames = append(funcNames, v.Name)
	}
	if !utils.Contains(funcNames, "main") {
		err := ast.NewSemanticAnalysisError("main function not found", "main function error", nil)
		errors = append(errors, err)
	}
	return errors
}

func Execute(program *ast.Program) *st.SymbolTables {

	//Define the compiler symbol tables
	tables := st.NewSymbolTables()

	// Collect all the errors
	errors := make([]*ast.SemanticAnalysisError, 0)

	// First Build the Symbol Table(s) for all global declarations
	errors = program.BuildSymbolTable(tables, errors)

	// Check for redeclaration errors
	errors = checkRedeclaration(errors, tables)

	// Check for the existance of a main function
	errors = checkForMain(errors, tables)

	// Second perform type checking
	// Pass the funcEntry as nil as it will be updated at the function nodes in the AST during the traverse
	errors = program.TypeCheck(errors, tables, nil)

	// Third perform control flow analysis
	errors = program.ControlFlowCheck(errors, tables, nil)

	// If there are any errors, report them to the user
	if hasErrors(errors) {
		reportErrors(errors)
		return nil
	}

	return tables
}
