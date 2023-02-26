package llvm

import (
	"fmt"
	"golite/types"
)

// Get the default value for a type.
func GetTypeDefault(t types.Type) string {
	if types.TypeToKind(t) == types.PRIMITIVE {
		if t == types.StringToType("int") {
			return "0"
		} else if t == types.StringToType("bool") {
			return "0" // Since we are representing bool as i64, 0 is false and 1 is true
		} else if t == types.StringToType("nil") {
			return "null"
		} else if t == types.StringToType("string") {
			return "null"
		} else if t == types.StringToType("void") {
			return "0" // Since we are representing void as i64
		} else {
			fmt.Println("Unknown primitive type")
			return "0"
		}
	} else {
		if types.TypeToKind(t) == types.STRUCT {
			return "null"
		} else {
			fmt.Println("Unknown type")
			return "null"
		}
	}
}
