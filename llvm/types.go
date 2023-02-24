package llvm

import "golite/types"

// Convert type to LLVM type
func TypeToLLVM(t types.Type) string {
	if types.TypeToKind(t) == types.PRIMITIVE {
		if t == types.StringToType("int") {
			return "i64"
		} else if t == types.StringToType("bool") {
			return "i64"
		} else if t == types.StringToType("nil") {
			return "i64"
		} else if t == types.StringToType("string") {
			return "i64"
		} else if t == types.StringToType("void") {
			return "void"
		} else {
			panic("Unknown primitive type")
		}
	} else {
		if types.TypeToKind(t) == types.STRUCT {
			return "struct." + t.String()[1:]
		} else {
			panic("Unknown type")
		}
	}
}
