package types

import (
	"strings"
	"sync"
)

// Type includes information about working on Types
type Type interface {
	String() string
}

// Lock for the singleton
var lock = &sync.Mutex{}

// Conversion from string representation to Type
func StringToType(ty string) Type {
	if ty == "int" {
		return getIntInstance()
	} else if ty == "bool" {
		return getBoolInstance()
	} else if ty == "string" {
		return getStringInstance()
	} else if ty == "void" {
		return getVoidInstance()
	} else if ty == "nil" {
		return getNilInstance()
	} else if strings.Contains(ty, "*") {
		return getPointerInstance(ty)
	}
	panic("Type not found.")
}

// Kind of type
const (
	PRIMITIVE = iota
	STRUCT
)

// Map from kind of type to type_kind
func TypeToKind(ty Type) int {
	if ty == getIntInstance() || ty == getBoolInstance() || ty == getStringInstance() || ty == getVoidInstance() || ty == getNilInstance() {
		return PRIMITIVE
	} else {
		return STRUCT
	}
}
