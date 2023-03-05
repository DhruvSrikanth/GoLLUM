package arm

// Sub represents the substraction instruction in ARM
type Sub struct {
	// The destination register
	Dest string
	// The first operand
	Op1 string
	// The second operand
	Op2 string

	blockLabel string
}

// NewSub returns a new Sub instruction
func NewSub(dest string, op1 string, op2 string) *Sub {
	return &Sub{dest, op1, op2, ""}
}

// String returns the string representation of the Sub instruction
func (s *Sub) String() string {
	return "sub " + s.Dest + ", " + s.Op1 + ", " + s.Op2
}

// Set the label of the instruction
func (s *Sub) SetLabel(label string) {
	s.blockLabel = label
}

// Get the label of the instruction
func (s *Sub) GetLabel() string {
	return s.blockLabel
}
