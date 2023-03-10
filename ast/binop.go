package ast

import (
	"bytes"
	"fmt"
	"golite/llvm"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

// Binary operation struct for rhs evaluation
type BinOpExpr struct {
	*token.Token             // The token information
	left         *Expression // The left operand expression
	operator     *Operator   // The operator for the binary expression
	right        *Expression // The right operand expression
	ty           types.Type  // The inferred type of the expression
}

// Is the binary operation an integer operation
func isIntOp(op Operator) bool {
	return op == ADD ||
		op == SUB ||
		op == MUL ||
		op == DIV
}

// Is the binary operation a comparison operation
func isCompareOp(op Operator) bool {
	return op == GT ||
		op == GE ||
		op == LT ||
		op == LE ||
		op == EQ ||
		op == NE
}

// Is the binary operation a boolean operation
func isBoolOp(op Operator) bool {
	return op == AND ||
		op == OR ||
		op == NOT
}

// Create nested binary operation expressions in a tree form
func MergeBinOps(terms []Expression, ops []Operator, tok *token.Token) Expression {
	var finalTerm Expression
	if len(terms) == 1 {
		finalTerm = terms[0]
	} else {
		op_idx := 0
		term_idx := 0
		for {
			if term_idx == len(terms)-1 {
				break
			}
			term1 := terms[term_idx]
			term2 := terms[term_idx+1]
			op := ops[op_idx]
			term := NewBinOp(&term1, &op, &term2, tok)
			terms[term_idx+1] = term
			term_idx++
			op_idx++
		}
		finalTerm = terms[len(terms)-1]
	}
	return finalTerm
}

// Return a new binary operation expression
func NewBinOp(left *Expression, op *Operator, right *Expression, token *token.Token) *BinOpExpr {
	return &BinOpExpr{token, left, op, right, nil}
}

// String representation of the binary operation
func (binOp *BinOpExpr) String() string {
	var out bytes.Buffer

	if binOp.operator != nil {
		out.WriteString("(")
	}
	if binOp.left != nil {
		out.WriteString((*binOp.left).String())
	}
	if binOp.operator != nil {
		out.WriteString(" " + OpToStr(*binOp.operator) + " ")
	}
	out.WriteString((*binOp.right).String())
	if binOp.operator != nil {
		out.WriteString(")")
	}

	return out.String()
}

// Infer the type of the binary operation
func (binOp *BinOpExpr) GetType() types.Type {
	return binOp.ty
}

// Build the symbol table for the binary operation
func (binOp *BinOpExpr) BuildSymbolTable(tables *st.SymbolTables, errors []*SemanticAnalysisError) []*SemanticAnalysisError {
	// Nothing to do here since the expression does not introduce any new symbols
	return errors
}

// Type check the binary operation
func (binOp *BinOpExpr) TypeCheck(errors []*SemanticAnalysisError, tables *st.SymbolTables, funcEntry *st.FuncEntry) []*SemanticAnalysisError {
	// Type check the left and right expressions
	if binOp.left != nil {
		errors = (*binOp.left).TypeCheck(errors, tables, funcEntry)
	}
	if binOp.right != nil {
		errors = (*binOp.right).TypeCheck(errors, tables, funcEntry)
	}

	// First check if it is a unary term
	if binOp.operator == nil && binOp.left == nil {
		binOp.ty = (*binOp.right).GetType()
	} else if binOp.operator != nil && binOp.left == nil {
		// Valid unary operators are - and !
		op := *binOp.operator
		rightType := (*binOp.right).GetType()
		if op == NOT && rightType == types.StringToType("bool") {
			binOp.ty = rightType
		} else if op == SUB && rightType == types.StringToType("int") {
			binOp.ty = rightType
		} else {
			var expectedType types.Type
			if op == NOT {
				expectedType = types.StringToType("bool")
			} else if op == SUB {
				expectedType = types.StringToType("int")
			}
			// Set the expected type for predictive type checking
			binOp.ty = expectedType

			// Record the error
			errors = append(errors, NewSemanticAnalysisError("invalid expression", "mismatched types", binOp.Token))
		}
	} else {

		leftType := (*binOp.left).GetType()
		rightType := (*binOp.right).GetType()
		// Check for the term
		if isIntOp(*binOp.operator) {
			// Could be a term or a simple term
			// Only ints allowed
			// Result is an int
			if leftType == types.StringToType("int") && rightType == types.StringToType("int") {
				binOp.ty = types.StringToType("int")
			} else {
				// Set the expected type for predictive type checking
				binOp.ty = types.StringToType("int")

				// Record the error
				errors = append(errors, NewSemanticAnalysisError("invalid expression", "mismatched types", binOp.Token))
			}
		} else if isCompareOp(*binOp.operator) {
			// Could be a relation term, equal term
			// Only ints allowed for relation terms
			// Ints, structs, bool and nil allowed for equal terms
			// Result is a bool
			if leftType == types.StringToType("int") && rightType == types.StringToType("int") {
				binOp.ty = types.StringToType("bool")
			} else if leftType == types.StringToType("bool") && rightType == types.StringToType("bool") {
				binOp.ty = types.StringToType("bool")
			} else if types.TypeToKind(leftType) == types.STRUCT && types.TypeToKind(rightType) == types.STRUCT {
				if *binOp.operator == EQ || *binOp.operator == NE {
					binOp.ty = types.StringToType("bool")
					// Struct comparison is only allowed between the same struct
					// struct equivalence is through name equivalence
					// this can be acomplished by comparing the pointers to the struct types
					if leftType != rightType {
						// Record the error
						errors = append(errors, NewSemanticAnalysisError("invalid struct comparison", "mismatched types", binOp.Token))
					}
				} else {
					// Set the expected type for predictive type checking
					binOp.ty = types.StringToType("bool")

					// Record the error
					errors = append(errors, NewSemanticAnalysisError("only struct to struct or struct to nil equivalence can be evaluated", "mismatched types", binOp.Token))
				}
			} else if types.TypeToKind(leftType) == types.STRUCT && rightType == types.StringToType("nil") || leftType == types.StringToType("nil") && types.TypeToKind(rightType) == types.STRUCT {
				// Comparing struct to nil
				if *binOp.operator == EQ || *binOp.operator == NE {
					binOp.ty = types.StringToType("bool")
				} else {
					// Set the expected type for predictive type checking
					binOp.ty = types.StringToType("bool")

					// Record the error
					errors = append(errors, NewSemanticAnalysisError("only struct to struct or struct to nil equivalence can be evaluated", "mismatched types", binOp.Token))
				}
			} else {
				// Set the expected type for predictive type checking
				binOp.ty = types.StringToType("bool")

				// Record the error
				errors = append(errors, NewSemanticAnalysisError("invalid expression", "mismatched types", binOp.Token))
			}
		} else if isBoolOp(*binOp.operator) {
			// Could be a bool term or an expression term
			// Only bools allowed
			// Result is a bool
			if leftType == types.StringToType("bool") && rightType == types.StringToType("bool") {
				binOp.ty = types.StringToType("bool")
			} else {
				// Set the expected type for predictive type checking
				binOp.ty = types.StringToType("bool")

				// Record the error
				errors = append(errors, NewSemanticAnalysisError("invalid expression", "mismatched types", binOp.Token))
			}
		} else {
			// Invalid operator
			// This should never happen
			errors = append(errors, NewSemanticAnalysisError("invalid operator provided in expression", "invalid operator", binOp.Token))
		}
	}

	return errors
}

// Translation of the expression to LLVM IR
func (b *BinOpExpr) ToLLVMCFG(tables *st.SymbolTables, blocks []*llvm.BasicBlock, funcEntry *st.FuncEntry, constDecls []*llvm.ConstantDecl) ([]*llvm.BasicBlock, []*llvm.ConstantDecl, string) {
	var op string
	var lastUsedRegLeft, lastUsedRegRight string
	var mostRecentOperand string
	if b.operator != nil && b.left != nil {
		// Convert the operator to its LLVM equivalent
		op = llvm.OperatorToLLVM(OpToStr(*b.operator))
	} else if b.operator != nil && b.left == nil {
		// In this case we create the left operand
		// Could be the negation or the not operator
		// If negative, we multiple by -1
		// If not, we perform a xor with 1
		if *b.operator == SUB {
			op = "mul"
			// Create the constant -1
			mostRecentOperand = "-1"
			lastUsedRegLeft = mostRecentOperand
		} else if *b.operator == NOT {
			op = "xor"
			// Create the constant 1
			mostRecentOperand = "1"
			lastUsedRegLeft = mostRecentOperand
		} else {
			// This should never happen
			fmt.Println("Invalid operator provided in expression")
		}
	} else if b.operator == nil && b.left != nil {
		// Impossible because the right is guaranteed and the left cannot exists without an operator for the right
	} else {
		blocks, constDecls, mostRecentOperand = (*b.right).ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		// Nothing to do here since left will also be nil and the right has been evaluated
		return blocks, constDecls, mostRecentOperand
	}

	// Call the ToLLVMCFG method on the left and right expressions
	if b.left != nil {
		blocks, constDecls, mostRecentOperand = (*b.left).ToLLVMCFG(tables, blocks, funcEntry, constDecls)
		// Get the last register used if there is a left expression
		lastUsedRegLeft = mostRecentOperand
	}

	blocks, constDecls, mostRecentOperand = (*b.right).ToLLVMCFG(tables, blocks, funcEntry, constDecls)
	lastUsedRegRight = mostRecentOperand

	// If either operand is not a number, we need to load it into a register
	if isNamed(funcEntry, lastUsedRegLeft[1:]) {
		// Create the load instruction
		ty := llvm.TypeToLLVM((*b.left).GetType())
		if strings.Contains(ty, "struct.") {
			ty += "*"
		}
		loadInst := llvm.NewLoad(lastUsedRegLeft, ty)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)
		mostRecentOperand = llvm.GetPreviousRegister()
		lastUsedRegLeft = mostRecentOperand
	} else if b.left != nil && (blocks[len(blocks)-1].Size() > 0) {
		if strings.Contains(blocks[len(blocks)-1].GetLastInstruction().String(), "getelementptr") && strings.Contains(blocks[len(blocks)-1].GetLastInstruction().String(), "struct.") && strings.Contains((*b.left).String(), ".") && !strings.Contains((*b.left).String(), " ") {
			// Create the load instruction
			ty := llvm.TypeToLLVM((*b.left).GetType())
			if strings.Contains(ty, "struct.") {
				ty += "*"
			}
			loadInst := llvm.NewLoad(lastUsedRegLeft, ty)
			loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
			blocks[len(blocks)-1].AddInstruction(loadInst)
			mostRecentOperand = llvm.GetPreviousRegister()
			lastUsedRegLeft = mostRecentOperand
		}
	}
	if isNamed(funcEntry, lastUsedRegRight[1:]) || (strings.Contains((*b.right).String(), ".") && !strings.Contains((*b.right).String(), " ") && !strings.Contains((*b.right).String(), " ")) {
		// Create the load instruction
		ty := llvm.TypeToLLVM((*b.right).GetType())
		if strings.Contains(ty, "struct.") {
			ty += "*"
		}
		loadInst := llvm.NewLoad(lastUsedRegRight, ty)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)
		mostRecentOperand = llvm.GetPreviousRegister()
		lastUsedRegRight = mostRecentOperand
	}

	// Figure out the kind type of operands we have
	var opType string
	// Perform a nil check
	var nilExprEq Expression
	if lastUsedRegLeft == "nil" {
		nilExprEq = *b.right
	} else if lastUsedRegRight == "nil" {
		nilExprEq = *b.left
	}
	// If any is nil, we need to set the value to nullptr for the particular type
	if lastUsedRegLeft == "nil" || lastUsedRegRight == "nil" {
		nilReg := "@.nil" + nilExprEq.GetType().String()[1:]
		llvmTy := llvm.TypeToLLVM(nilExprEq.GetType())
		if strings.Contains(llvmTy, "struct.") {
			llvmTy += "*"
		}
		opType = llvmTy

		// Create the load instruction
		loadInst := llvm.NewLoad(nilReg, llvmTy)
		loadInst.SetLabel(blocks[len(blocks)-1].GetLabel())
		blocks[len(blocks)-1].AddInstruction(loadInst)
		mostRecentOperand = llvm.GetPreviousRegister()

	} else {
		// Set the operand type
		// Use right as left may not exist
		opType = llvm.TypeToLLVM((*b.right).GetType())
		if strings.Contains(opType, "struct.") {
			opType += "*"
		}
	}

	if lastUsedRegLeft == "nil" {
		lastUsedRegLeft = mostRecentOperand
	} else if lastUsedRegRight == "nil" {
		lastUsedRegRight = mostRecentOperand
	}

	// Perform the operation
	operationInst := llvm.NewBinOp(lastUsedRegLeft, op, opType, lastUsedRegRight)
	operationInst.SetLabel(blocks[len(blocks)-1].GetLabel())
	blocks[len(blocks)-1].AddInstruction(operationInst)
	mostRecentOperand = llvm.GetPreviousRegister()

	return blocks, constDecls, mostRecentOperand
}
