package arm

// CSET ARM64 instruction
type Cset struct {
	// The target register
	targetRegister string
	// The condition
	cond string
	// The block label
	blockLabel string
}

// NewCset returns a new CSET instruction
func NewCset(targetRegister string, cond string) *Cset {
	return &Cset{targetRegister, cond, ""}
}

// String representation of the CSET instruction
func (c *Cset) String() string {
	// Format is cset <target>, <cond>
	return "cset " + c.targetRegister + ", " + c.cond
}

// Set the block label
func (c *Cset) SetLabel(label string) {
	c.blockLabel = label
}

// Get the block label
func (c *Cset) GetLabel() string {
	return c.blockLabel
}
