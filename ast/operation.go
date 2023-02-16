package ast

import "fmt"

// Alias for the operator type
type Operator int

// Values for the operator type
const (
	// Int operators
	ADD Operator = iota
	SUB
	MUL
	DIV

	// Comparison operators
	GT
	GE
	LT
	LE
	EQ
	NE

	// Bool operators
	AND
	OR
	NOT
)

// Convert a string representation to an operator
func StrToOp(op string) Operator {
	switch op {

	// Int operators
	case "+":
		return ADD
	case "-":
		return SUB
	case "*":
		return MUL
	case "/":
		return DIV

	// Comparison operators
	case ">":
		return GT
	case ">=":
		return GE
	case "<":
		return LT
	case "<=":
		return LE
	case "==":
		return EQ
	case "!=":
		return NE

	// Bool operators
	case "&&":
		return AND
	case "||":
		return OR
	case "!":
		return NOT
	}

	fmt.Println(op)
	panic("Error: Could not determine operator")

}

// Convert an operator to a string representation
func OpToStr(op Operator) string {
	switch op {
	// Int operators
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"

	// Comparison operators
	case GT:
		return ">"
	case GE:
		return ">="
	case LT:
		return "<"
	case LE:
		return "<="
	case EQ:
		return "=="
	case NE:
		return "!="

	// Bool operators
	case AND:
		return "&&"
	case OR:
		return "||"
	case NOT:
		return "!"
	}
	panic("Error: Could not determine operator")
}
