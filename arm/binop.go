package arm

import "bytes"

// Represents an ARM64 binary operation.
type BinOp struct {
	destinationRegister string
	operator            string
	leftRegister        string
	rightRegister       string
	blockLabel          string
}

// Returns a new ARM64 binary operation.
func NewBinOp(destinationRegister, operator, leftRegister, rightRegister string) *BinOp {
	return &BinOp{destinationRegister, operator, leftRegister, rightRegister, ""}
}

// Returns the string representation of the ARM64 binary operation.
func (b *BinOp) String() string {
	var out bytes.Buffer
	out.WriteString(b.operator)
	out.WriteString(" ")
	out.WriteString(b.destinationRegister)
	out.WriteString(", ")
	out.WriteString(b.leftRegister)
	out.WriteString(", ")
	out.WriteString(b.rightRegister)
	return out.String()
}

// Sets the block label of the ARM64 binary operation.
func (b *BinOp) SetLabel(label string) {
	b.blockLabel = label
}

// Returns the block label of the ARM64 binary operation.
func (b *BinOp) GetLabel() string {
	return b.blockLabel
}
