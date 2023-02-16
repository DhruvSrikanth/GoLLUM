package symboltable

import (
	"fmt"
	"golite/types"
)

// Struct field entry
type FieldEntry struct {
	Name string
	Ty   types.Type
}

// String representation of a struct field entry in the symbol table
func (entry *FieldEntry) String() string {
	return fmt.Sprintf("[%s (Type: %s)]", entry.Name, entry.Ty)
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
		if i < len(entry.Fields)-1 {
			fields += fmt.Sprintf("%s, ", field.String())
		} else {
			fields += field.String()
		}
	}
	return fmt.Sprintf("%s {Fields: %s}", entry.Name, fields)
}
