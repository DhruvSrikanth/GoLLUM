package llvm

import "golite/types"

// Get the default value for a type.
func GetTypeDefault(t types.Type) string {
	if types.TypeToKind(t) == types.PRIMITIVE {
		if t == types.StringToType("int") {
			return "0"
		} else if t == types.StringToType("bool") {
			return "0"
		} else if t == types.StringToType("nil") {
			return "null"
		} else if t == types.StringToType("string") {
			return "null"
		} else if t == types.StringToType("void") {
			return "null"
		} else {
			panic("Unknown primitive type")
		}
	} else {
		if types.TypeToKind(t) == types.STRUCT {
			return "null"
		} else {
			panic("Unknown type")
		}
	}
}
