package llvm

import (
	"bytes"
	"strings"
)

// Representation of a global declaration
type Decl struct {
	// The name of the global declaration
	name    string
	ty      string
	initVal string
	isNil   bool
}

// NewglobalDecl returns a new global declaration
func NewDecl(name, ty, initVal string, isNil bool) *Decl {
	return &Decl{name, ty, initVal, isNil}
}

// String representation of the global declaration
func (s *Decl) String() string {
	var out bytes.Buffer
	// Format is @variable_name = common global type init_value
	// Common for all
	out.WriteString("@")
	// Global null value declaration for the struct
	if s.isNil {
		nonStructName := strings.Split(s.name, ".")[1]
		out.WriteString(".nil" + nonStructName)
	} else {
		out.WriteString(s.name)
	}
	// Common for all
	out.WriteString(" = common global ")
	// Global null or just struct
	if s.isNil || strings.Contains(s.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(s.ty)

	if s.isNil || strings.Contains(s.ty, "struct.") {
		out.WriteString("*")
	}
	out.WriteString(" " + s.initVal)

	return out.String()
}
