package llvm

import (
	"fmt"
	"golite/types"
)

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
			return "i64"
		} else {
			fmt.Println("Unknown primitive type")
			return ""
		}
	} else {
		if types.TypeToKind(t) == types.STRUCT {
			return "struct." + t.String()[1:]
		} else {
			fmt.Println("Unknown type")
			return ""
		}
	}
}
