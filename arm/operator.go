package arm

// Convert operator from LLVM IR to ARM64 assembly.
func LLVMOpToARMOp(op string) string {
	switch op {
	// Int operators
	case "add":
		return "add"
	case "sub":
		return "sub"
	case "mul":
		return "mul"
	case "sdiv":
		return "sdiv"

	// Comparison operators
	case "icmp sgt":
		return "gt"
	case "icmp sge":
		return "ge"
	case "icmp slt":
		return "lt"
	case "icmp sle":
		return "le"
	case "icmp eq":
		return "eq"
	case "icmp ne":
		return "ne"

	// Bool operators
	case "and":
		return "and"
	case "or":
		return "or"
	case "xor":
		return "eor" // This is for because NOT (Input) = XOR (1 Input)
	}
	return ""
}

// Is the operator a comparison operator?
func IsComparisonOp(op string) bool {
	switch op {
	case "icmp sgt":
		return true
	case "icmp sge":
		return true
	case "icmp slt":
		return true
	case "icmp sle":
		return true
	case "icmp eq":
		return true
	case "icmp ne":
		return true
	}
	return false
}
