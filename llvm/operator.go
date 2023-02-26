package llvm

// Convert operator to LLVM IR op code
func OperatorToLLVM(op string) string {
	switch op {
	// Int operators
	case "+":
		return "add"
	case "-":
		return "sub"
	case "*":
		return "mul"
	case "/":
		return "sdiv"

	// Comparison operators
	case ">":
		return "icmp sgt"
	case ">=":
		return "icmp sge"
	case "<":
		return "icmp slt"
	case "<=":
		return "icmp sle"
	case "==":
		return "icmp eq"
	case "!=":
		return "icmp ne"

	// Bool operators
	case "&&":
		return "and"
	case "||":
		return "or"
	case "!":
		return "xor" // This is for because NOT (Input) = XOR (1 Input)
	}
	return ""
}
