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

// IntType is a type that represents an integer
type IntType struct{}

// Singleton instance of IntType
var singletonInt *IntType

// String representation of IntType
func (intTy *IntType) String() string {
	return "int"
}

// Get the singleton instance of IntType (thread safe)
func getIntInstance() *IntType {
	if singletonInt == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonInt == nil {
			// This thread is the first to create the singleton
			singletonInt = &IntType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonInt
}

// BoolType is a type that represents a boolean
type BoolType struct{}

// Singleton instance of BoolType
var singletonBool *BoolType

// String representation of BoolType
func (boolTy *BoolType) String() string {
	return "bool"
}

// Get the singleton instance of BoolType (thread safe)
func getBoolInstance() *BoolType {
	if singletonBool == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonBool == nil {
			// This thread is the first to create the singleton
			singletonBool = &BoolType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonBool
}

// StringType is a type that represents a string
type StringType struct{}

// Singleton instance of StringType
var singletonString *StringType

// String representation of StringType
func (stringTy *StringType) String() string {
	return "string"
}

// Get the singleton instance of StringType (thread safe)
func getStringInstance() *StringType {
	if singletonString == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonString == nil {
			// This thread is the first to create the singleton
			singletonString = &StringType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonString
}

// NilType is a type that represents a nil value
type NilType struct{}

// Singleton instance of NilType
var singletonNil *NilType

// String representation of NilType
func (nilTy *NilType) String() string {
	return "nil"
}

// Get the singleton instance of NilType (thread safe)
func getNilInstance() *NilType {
	if singletonNil == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonNil == nil {
			// This thread is the first to create the singleton
			singletonNil = &NilType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonNil
}

// Conversion from string representation to Type
func StringToType(ty string) Type {
	if ty == "int" {
		return getIntInstance()
	} else if ty == "bool" {
		return getBoolInstance()
	} else if ty == "string" {
		return getStringInstance()
	} else if ty == "nil" {
		return getNilInstance()
	}
	panic("Type not found.")
}
