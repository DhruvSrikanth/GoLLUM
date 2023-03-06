package arm

// Represents the CMP ARM64 instruction
type Cmp struct {
	// The first register
	register1 string
	// The second register
	register2 string
	// The block label
	blockLabel string
}

// NewCmp returns a new CMP instruction
func NewCmp(register1 string, register2 string) *Cmp {
	return &Cmp{register1, register2, ""}
}

// String representation of the CMP instruction
func (c *Cmp) String() string {
	// Format is cmp <register1>, <register2>
	return "cmp " + c.register1 + ", " + c.register2
}

// Set the block label
func (c *Cmp) SetLabel(label string) {
	c.blockLabel = label
}

// Get the block label
func (c *Cmp) GetLabel() string {
	return c.blockLabel
}
