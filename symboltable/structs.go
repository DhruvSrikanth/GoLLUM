package symboltable

import "fmt"

// Struct entry in the symbol table
type StructEntry struct {
	Name   string
	Fields []*VarEntry
}

// String representation of a struct entry in the symbol table
func (entry *StructEntry) String() string {
	fields := ""
	for _, f := range entry.Fields {
		field := *f
		fields += fmt.Sprintf("%s, ", field.String())
	}
	fields = fields[:len(fields)-2]
	return fmt.Sprintf("%s {Fields: %s}", entry.Name, fields)
}
