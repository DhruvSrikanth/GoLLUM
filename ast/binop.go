package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Binary operation struct for rhs evaluation
type BinOpExpr struct {
	*token.Token            // The token information
	operator     Operator   // The operator for the binary expression
	right        Expression // The right operand expression
	left         Expression // The left operand expression
	ty           types.Type // The inferred type of the expression
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

// Return a new binary operation expression
func NewBinOp(op Operator, right, left Expression, token *token.Token) *BinOpExpr {
	return &BinOpExpr{token, op, right, left, nil}
}

// String representation of the binary operation
func (binOp *BinOpExpr) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(binOp.left.String())
	out.WriteString(" " + OpToStr(binOp.operator) + " ")
	out.WriteString(binOp.right.String())
	out.WriteString(")")

	return out.String()
}

// Infer the type of the binary operation
func (binOp *BinOpExpr) GetType() types.Type {
	return binOp.ty
}

// Build the symbol table for the binary operation
func (binOp *BinOpExpr) BuildSymbolTable(tables *st.SymbolTables) {
	// Nothing to do here since the expression does not introduce any new symbols
}

// Type check the binary operation
func (binOp *BinOpExpr) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	return errors
}
