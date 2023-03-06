package arm

// STR ARM64 instruction
type Str struct {
	// The source register
	source string
	// The address of the desintation
	destAddress string

	// block label
	blockLabel string
}

// NewSTR returns a new STR instruction
func NewStr(source string, destAddress string) *Str {
	return &Str{source, destAddress, ""}
}

// String returns the string representation of the STR instruction
func (s *Str) String() string {
	return "str " + s.source + ", [" + s.destAddress + "]"
}

// Set the label of the instruction
func (s *Str) SetLabel(label string) {
	s.blockLabel = label
}

// Get the label of the instruction
func (s *Str) GetLabel() string {
	return s.blockLabel
}
