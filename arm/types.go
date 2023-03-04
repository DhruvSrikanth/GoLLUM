package arm

import "golite/types"

// Get the default size of golite types
func GetSizeGolite(t types.Type) int {
	switch t {
	case types.StringToType("int"):
		return 8
	case types.StringToType("bool"):
		return 8
	case types.StringToType("void"):
		return 8
	case types.StringToType("string"):
		return 8
	case types.StringToType("nil"):
		return 8
	default:
		// This coverts pointers which will also have a size of 8
		return 8
	}
}

// Get the alignment of golite types
func GetAlignmentGolite(t types.Type) int {
	switch t {
	case types.StringToType("int"):
		return 8
	case types.StringToType("bool"):
		return 8
	case types.StringToType("void"):
		return 8
	case types.StringToType("string"):
		return 8
	case types.StringToType("nil"):
		return 8
	default:
		// This coverts pointers which will also have a size of 8
		return 8
	}
}

// Get the default size of llvm types
func GetSizeLlvm(t string) int {
	switch t {
	case "i64":
		return 8
	default:
		// This coverts pointers which will also have a size of 8
		return 8
	}
}

// Get the alignment of llvm types
func GetAlignmentLlvm(t string) int {
	switch t {
	case "i64":
		return 8
	default:
		// This coverts pointers which will also have a size of 8
		return 8
	}
}
