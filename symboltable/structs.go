package symboltable

import (
	"fmt"
	"golite/llvm"
	"golite/types"
)

// Struct field entry
type FieldEntry struct {
	Name     string
	Ty       types.Type
	LlvmType string
}

// String representation of a struct field entry in the symbol table
func (entry *FieldEntry) String() string {
	return fmt.Sprintf("%s (Type: %s [%s])", entry.Name, entry.Ty, entry.LlvmType)
}

// Get the type of variable
func (entry *FieldEntry) GetType() types.Type {
	return entry.Ty
}

// Translate the struct field entry to its LLVM representation
func (entry *FieldEntry) LLVMTypeConversion() {
	entry.LlvmType = llvm.TypeToLLVM(entry.Ty)
}

// Struct entry in the symbol table
type StructEntry struct {
	Name   string
	Fields []*FieldEntry
}

// String representation of a struct entry in the symbol table
func (entry *StructEntry) String() string {
	fields := ""
	for i, f := range entry.Fields {
		field := *f
		if i < len(entry.Fields)-1 && len(entry.Fields) > 1 {
			fields += fmt.Sprintf("%s, ", field.String())
		} else {
			fields += field.String()
		}
	}
	return fmt.Sprintf("%s\nFields -  %s", entry.Name, fields)
}

// Translate the struct entries to their LLVM representation
func (entry *StructEntry) LLVMTypeConversion() {
	for _, field := range entry.Fields {
		field.LLVMTypeConversion()
	}
}
