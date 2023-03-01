package ast

import (
	"bytes"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// Assignment node struct for the AST

type Assignment struct {
	*token.Token
	variable LValue     // The lvalue for the assignment statement
	right    Expression // The value assigned to the variable of this statements
}

// New Assignment node
func NewAssignment(variable LValue, right Expression, token *token.Token) *Assignment {
	return &Assignment{token, variable, right}
}

// String representation of the assignment node
func (a *Assignment) String() string {
	var out bytes.Buffer

	out.WriteString(a.variable.String())
	out.WriteString(" = ")
	out.WriteString(a.right.String())
	out.WriteString(";")

	return out.String()
}

// Build the symbol table for the assignment
func (a *Assignment) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since an assignment is not added to the symbol table
	return errors
}

// Type checking for the assignment
func (a *Assignment) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the expression on the right hand side of the assignment
	// Also ensures that GetType() on expression is the correct predictive type
	errors = a.right.TypeCheck(errors, tables, funcEntry)

	// Type check the lvalue on the left hand side of the assignment
	errors = a.variable.TypeCheck(errors, tables, funcEntry)

	if a.variable.GetType() == nil {
		// Use inferred data type from rhs to continue using predictive type checking
		// This ensures that the compiler can continue type checking to find more errors
		// Error for this nil type will be added in the type checking of the lvalue so all errors can be accounted for
		a.variable.ty = a.right.GetType()
		errors = append(errors, NewSemanticAnalysisError("type mismatch in assignment", "type mistmatch", a.GetToken()))
	}

	// Check that the type of the lvalue is the same as the type of the expression
	// Check if it is a struct type
	if a.variable.GetType() != a.right.GetType() {
		// If the types are not the same, check if the lvalue is a struct type
		// If it is, the right hand side must also be a struct literal or nil
		// In this case since it is not a struct literal or atleast not the same struct literal, we can check for nil
		if (types.TypeToKind(a.variable.GetType()) == types.STRUCT) && (a.right.GetType() == types.StringToType("nil")) {
			return errors
		}
		errors = append(errors, NewSemanticAnalysisError("type mismatch in assignment", "type mistmatch", a.GetToken()))
	}

	return errors
}

// Get the token for the assignment
func (a *Assignment) GetToken() *token.Token {
	return a.Token
}

// Control flow analysis for the assignment node
func (a *Assignment) GetControlFlow(errors []*SemanticAnalysisError, funcEntry *st.FuncEntry) ([]*SemanticAnalysisError, bool) {
	// Nothing to do here since the assignment node does not have any control flow
	return errors, false
}

// Translate the assignment node to LLVM IR
func (a *Assignment) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl) {
	// Stay in the same block
	var mostRecentOperand string

	var leftReg string
	useNamedReg := false
	// Check if it us a local variable, global variable, or a parameter
	// local - %name
	// global - @name
	// parameter - %P_name

	varEntry := funcEntry.Variables.Contains(a.variable.String())
	if varEntry != nil {
		// This is a local variable or global variable
		useNamedReg = true
		// Get the appropriate scope
		if varEntry.Scope == st.GLOBAL {
			leftReg = "@" + a.variable.String()
		} else {
			leftReg = "%" + a.variable.String()
		}
	} else {
		for _, param := range funcEntry.Parameters {
			if param.Name == a.variable.String() {
				useNamedReg = true
				leftReg = "%P_" + param.Name
				break
			}
		}
	}
	if !useNamedReg {
		// If we are not using a named register, we need to evaluate the lhs value and get the register that the result is stored at
		// Get the address of the lvalue on the left hand side of the assignment
		blocks, constDecls, mostRecentOperand = a.variable.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		leftReg = mostRecentOperand
	}

	// Get the value of the expression on the right hand side of the assignment
	var rightReg string
	if a.right.GetType() == types.StringToType("nil") {
		// If the right hand side is nil, we need to set the value to nullptr for the particular type
		// We can do this by performing a load from the address of the default value for the type
		// Get the type of the lvalue
		ty := a.variable.GetType()
		nilReg := "@.nil" + ty.String()[1:]
		llvmTy := llvm.TypeToLLVM(ty)
		if strings.Contains(llvmTy, "struct.") {
			llvmTy += "*"
		}

		// Create the load instruction
		loadInst := llvm.NewLoad(nilReg, llvmTy)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)
		mostRecentOperand = llvm.GetPreviousRegister()

		// Get the register that the default value was loaded into
		rightReg = mostRecentOperand
	} else {
		// If we are not using a named register, we need to evaluate the rhs expression and get the register that the result is stored at
		blocks, constDecls, mostRecentOperand = a.right.ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		rightReg = mostRecentOperand
	}

	ty := a.variable.GetType()
	llvmTy := llvm.TypeToLLVM(ty)
	if strings.Contains(llvmTy, "struct.") {
		llvmTy += "*"
	}

	// Store the value of the expression into the address of the lvalue
	storeInst := llvm.NewStore(rightReg, leftReg, llvmTy)
	storeInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(storeInst)

	return blocks, constDecls
}
