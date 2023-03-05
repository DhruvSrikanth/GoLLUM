package arm

// STR ARM64 instruction
type Str struct {
	// The source register
	Source string
	// The offset of the desintation register from sp
	SPOffset string

	// block label
	blockLabel string
}

// NewSTR returns a new STR instruction
func NewStr(source string, spOffset string) *Str {
	return &Str{source, spOffset, ""}
}

// String returns the string representation of the STR instruction
func (s *Str) String() string {
	return "str " + s.Source + ", [sp, #" + s.SPOffset + "]"
}

// Set the label of the instruction
func (s *Str) SetLabel(label string) {
	s.blockLabel = label
}

// Get the label of the instruction
func (s *Str) GetLabel() string {
	return s.blockLabel
}
