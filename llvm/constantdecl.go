package llvm

import "strconv"

// Representation of a constant declaration
type ConstantDecl struct {
	varName string
	length  int
	value   string
}

// NewConstantDecl returns a new constant declaration
func NewConstantDecl(varName string, length int, value string) *ConstantDecl {
	return &ConstantDecl{varName, length, value}
}

// String representation of the constant declaration
func (c *ConstantDecl) String() string {
	return "@." + c.varName + " = private unnamed_addr constant [" + strconv.Itoa(c.length) + " x i8] c\"" + c.value + "\\00\", align 1"
}
