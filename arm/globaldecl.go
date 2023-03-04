package arm

import (
	"bytes"
	"strconv"
)

// Representation of a global declaration
type GlobalDecl struct {
	// The name of the global declaration
	name      string
	size      int
	alignment int
}

// NewglobalDecl returns a new global declaration
func NewGlobalDecl(name string, size, alignment int) *GlobalDecl {
	return &GlobalDecl{name, size, alignment}
}

// String representation of the global declaration
func (s *GlobalDecl) String() string {
	var out bytes.Buffer
	// Format is .comm var_name,size,alignment
	out.WriteString(".comm ")
	out.WriteString(s.name)
	out.WriteString(",")
	out.WriteString(strconv.Itoa(s.size))
	out.WriteString(",")
	out.WriteString(strconv.Itoa(s.alignment))

	return out.String()
}
