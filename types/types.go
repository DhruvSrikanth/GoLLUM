package types

import (
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
	}
	panic("Type not found.")
}
