package llvm

import (
	"golite/arm"
	"strconv"
	"strings"
)

// Representation of a constant declaration (global scope)
type ConstantDecl struct {
	varName       string
	length        int
	value         string
	originalValue string
}

// NewConstantDecl returns a new constant declaration
func NewConstantDecl(varName string, length int, value string, originalValue string) *ConstantDecl {
	return &ConstantDecl{varName, length, value, originalValue}
}

// String representation of the constant declaration
func (c *ConstantDecl) String() string {
	return "@." + c.varName + " = private unnamed_addr constant [" + strconv.Itoa(c.length) + " x i8] c\"" + c.value + "\\00\", align 1"
}

// Get the variable name of the constant declaration.
func (c *ConstantDecl) GetVarName() string {
	return c.varName
}

// Convert the constant declaration to ARM assembly.
func (c *ConstantDecl) ToARM() *arm.ConstantDecl {
	label := c.varName
	value := c.originalValue
	// Replace all %d with %ld
	value = strings.Replace(value, "%d", "%ld", -1)
	// Add 1 for the null terminator
	size := len(value) + 1
	return arm.NewConstantDecl(label, size, value)

}
