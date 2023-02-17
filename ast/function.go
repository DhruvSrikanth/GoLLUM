package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Function node for the AST
type Function struct {
	*token.Token
	name         string
	parameters   []Decl
	returnType   types.Type
	declarations []Declaration
	statements   []Statement
}

// New Function node
func NewFunction(name string, parameters []Decl, declarations []Declaration, statements []Statement, returnType types.Type, token *token.Token) *Function {
	return &Function{token, name, parameters, returnType, declarations, statements}
}

// String representation of the function node
func (f *Function) String() string {
	var out bytes.Buffer
	out.WriteString("func ")
	out.WriteString(f.name)

	out.WriteString("(")
	for i, param := range f.parameters {
		out.WriteString(param.String())
		if i < len(f.parameters)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") ")

	out.WriteString(f.returnType.String())

	out.WriteString(" {\n")

	for _, decl := range f.declarations {
		out.WriteString(decl.String())
		out.WriteString("\n")
	}

	for _, stmt := range f.statements {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	out.WriteString("}")

	return out.String()
}

// Build the symbol table for the function
func (f *Function) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Build up the entries for the parameters
	params := make([]*st.VarEntry, 0)
	for _, param := range f.parameters {
		params = append(params, &st.VarEntry{Name: param.variable, Ty: param.ty, Scope: st.LOCAL})
	}

	// Function local symbol table
	localTable := st.NewSymbolTable(tables.Globals)

	// Build up the entries for the variables
	for _, decl := range f.declarations {
		if !localTable.Insert(decl.variable, &st.VarEntry{Name: decl.variable, Ty: decl.ty, Scope: st.LOCAL}) {
			errors = append(errors, NewSemanticAnalysisError("Variable '"+decl.variable+"' redeclared.", "redeclaration", decl.Token))
		}
	}

	// Add the function to the symbol table
	if !tables.Funcs.Insert(f.name, &st.FuncEntry{Name: f.name, RetTy: f.returnType, Parameters: params, Variables: localTable}) {
		errors = append(errors, NewSemanticAnalysisError("Function '"+f.name+"' redeclared.", "redeclaration", f.Token))
	}
	return errors
}

// Type checking for the function
func (f *Function) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables) []*SemanticAnalysisError {
	// // Check the parameters
	// for _, param := range f.parameters {
	// 	errors = param.TypeCheck(errors, tables)
	// }

	// // Check the declarations
	// for _, decl := range f.declarations {
	// 	errors = decl.TypeCheck(errors, tables)
	// }

	// // Check the statements
	// for _, stmt := range f.statements {
	// 	errors = stmt.TypeCheck(errors, tables)
	// }

	return errors
}
