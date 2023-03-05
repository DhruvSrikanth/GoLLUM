package arm

// Add represents the addition instruction in ARM
type Add struct {
	// The destination register
	Dest string
	// The first operand
	Op1 string
	// The second operand
	Op2 string

	blockLabel string
}

// NewAdd returns a new Add instruction
func NewAdd(dest string, op1 string, op2 string) *Add {
	return &Add{dest, op1, op2, ""}
}

// String returns the string representation of the Add instruction
func (a *Add) String() string {
	return "add " + a.Dest + ", " + a.Op1 + ", " + a.Op2
}

// Set the label of the instruction
func (a *Add) SetLabel(label string) {
	a.blockLabel = label
}

// Get the label of the instruction
func (a *Add) GetLabel() string {
	return a.blockLabel
}
