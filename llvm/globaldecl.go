package llvm

import (
	"bytes"
	"golite/arm"
	"strings"
)

// Representation of a global declaration
type GlobalDecl struct {
	// The name of the global declaration
	name    string
	ty      string
	initVal string
	isNil   bool
}

// NewglobalDecl returns a new global declaration
func NewGlobalDecl(name, ty, initVal string, isNil bool) *GlobalDecl {
	return &GlobalDecl{name, ty, initVal, isNil}
}

// String representation of the global declaration
func (s *GlobalDecl) String() string {
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

// Returns if it is a struct type declaration
func (s *GlobalDecl) IsStructType() bool {
	return strings.Contains(s.name, "struct.")
}

// Is this a nil declaration
func (s *GlobalDecl) IsNil() bool {
	return s.isNil
}

// Get the name of the global declaration
func (s *GlobalDecl) GetName() string {
	return s.name
}

// Set the name of the global declaration
func (s *GlobalDecl) SetName(name string) {
	s.name = name
}

// Convert to ARM
func (s *GlobalDecl) ToARM() *arm.GlobalDecl {
	size := arm.GetSizeLlvm(s.ty)
	alignment := arm.GetAlignmentLlvm(s.ty)
	return arm.NewGlobalDecl(s.name, size, alignment)
}
