package llvm

import (
	"bytes"
	"strings"
)

// Representation of a struct declaration
type StructDecl struct {
	// The name of the struct declaration
	name   string
	fields []string
}

// NewStructDecl returns a new struct declaration
func NewStructDecl(name string, fields []string) *StructDecl {
	return &StructDecl{name, fields}
}

// String representation of the struct declaration
func (s *StructDecl) String() string {
	var out bytes.Buffer
	// Struct declaration
	out.WriteString("%" + s.name)
	out.WriteString(" = type {")
	for i, field := range s.fields {
		if i != 0 {
			out.WriteString(", ")
		}
		if strings.Contains(field, "struct.") {
			out.WriteString("%")
		}
		out.WriteString(field)
		if strings.Contains(field, "struct.") {
			out.WriteString("*")
		}
	}
	out.WriteString("}")

	return out.String()
}
