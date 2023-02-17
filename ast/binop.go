package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
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
			errors = append(errors, NewSemanticAnalysisError("Invalid expression", "mismatched types", binOp.Token))
		}
	} else {

	}

	return errors
}
