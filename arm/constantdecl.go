package arm

import (
	"strconv"
	"strings"
)

// Representation of a constant declaration (global scope)
type ConstantDecl struct {
	constLabel string
	constSize  int
	constValue string
}

// NewConstantDecl returns a new constant declaration
func NewConstantDecl(constLabel string, constSize int, constValue string) *ConstantDecl {
	constLabel = "." + strings.ToUpper(constLabel)
	return &ConstantDecl{constLabel, constSize, constValue}
}

// String representation of the constant declaration in ARM assembly
func (c *ConstantDecl) String() string {
	var out string
	out += c.constLabel + ":\n"
	out += "\t.asciz\t\"" + c.constValue + "\"" + "\n"
	out += "\t.size\t" + c.constLabel + ", " + strconv.Itoa(c.constSize) + "\n"
	return out
}
