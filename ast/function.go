package ast

import (
	"bytes"
	"golite/llvm"
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
	funcEntry    *st.FuncEntry
}

// New Function node
func NewFunction(name string, parameters []Decl, declarations []Declaration, statements []Statement, returnType types.Type, token *token.Token) *Function {
	return &Function{token, name, parameters, returnType, declarations, statements, nil}
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

	// Set the funcEntry for the function
	f.funcEntry = &st.FuncEntry{Name: f.name, RetTy: f.returnType, Parameters: params, Variables: localTable}

	// Add the function to the symbol table
	if !tables.Funcs.Insert(f.name, f.funcEntry) {
		errors = append(errors, NewSemanticAnalysisError("Function '"+f.name+"' redeclared.", "redeclaration", f.Token))
	}

	return errors
}

// Type checking for the function
func (f *Function) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Pass funcEntry formed in the function AST node to the children nodes to provide context
	// for type checking

	// Check the parameters
	for _, param := range f.parameters {
		errors = param.TypeCheck(errors, tables, f.funcEntry)
	}

	// Check the declarations
	for _, decl := range f.declarations {
		errors = decl.TypeCheck(errors, tables, f.funcEntry)
	}

	// Check the statements
	for _, stmt := range f.statements {
		errors = stmt.TypeCheck(errors, tables, f.funcEntry)
	}

	return errors
}

// Get the return type of the function
func (f *Function) GetReturnType() types.Type {
	return f.returnType
}

// Perform a control flow check on the function
func (f *Function) ControlFlowCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Check the control flow of the entire function
	var flow bool
	errors, flow = f.GetControlFlow(errors, f.funcEntry)
	// If function return is void and no control flow, its accepted
	// If function return is void and control flow, its accepted
	// If function return is not void and control flow, its accepted

	// If function return is not void and no control flow, its not accepted

	// If the flow is false, then we check to see if the function is a void function
	// If the function is a not a void function and the control flow is false, then we
	// have an error
	if f.GetReturnType() != types.StringToType("void") && !flow {
		errors = append(errors, NewSemanticAnalysisError("control flow does not reach return statement", "invalid control flow", f.Token))
	}

	return errors
}

// Get the control flow
func (f *Function) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Check the statements
	flow := false
	cFlow := false
	for _, stmt := range f.statements {
		errors, cFlow = stmt.GetControlFlow(errors, f.funcEntry)
		flow = flow || cFlow
	}
	return errors, flow
}

// Translate the declaration to LLVM IR
func (f *Function) ToLLVM(tables *st.SymbolTables) (*st.SymbolTables, *llvm.FunctionDecl) {
	// Declarations are translated in the program and function
	// Get the function entry
	funcEntry := tables.Funcs.Contains(f.name)

	// Get the parameters
	params := make([]string, 0)
	paramTypes := make([]string, 0)
	for _, param := range funcEntry.Parameters {
		params = append(params, param.Name)
		paramTypes = append(paramTypes, param.LlvmTy)
	}

	// Get the basic blocks
	blocks := make([]*llvm.BasicBlock, 0)

	return tables, llvm.NewFunctionDecl(funcEntry.Name, funcEntry.LlvmRetTy, params, paramTypes, blocks)
}
