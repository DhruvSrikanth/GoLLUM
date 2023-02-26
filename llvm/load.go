package llvm

import (
	"bytes"
	"strconv"
	"strings"
)

// Representation of a load operation
type Load struct {
	variable        string
	ty              string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewLoad returns a new Load operation
func NewLoad(varName, ty string) *Load {
	srcR := make([]int, 0)
	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &Load{varName, ty, "", srcR, tgtR}
}

// String representation of the Local declaration
func (l *Load) String() string {
	var out bytes.Buffer
	// Format is %result = load type, type* %var_name
	// Common for all
	out.WriteString("%")
	out.WriteString("r" + strconv.Itoa(l.targetRegisters[len(l.targetRegisters)-1]))
	out.WriteString(" = load ")
	// struct type
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(l.ty)
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("*")
	}
	// Common for all
	out.WriteString(", ")
	// struct type
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("%")
	}
	// Common for all
	out.WriteString(l.ty)
	if strings.Contains(l.ty, "struct.") {
		out.WriteString("*")
	}
	// Common for all
	out.WriteString("*")
	out.WriteString(" %")
	out.WriteString(l.variable)

	return out.String()
}

// Get the registers targeted by the instruction.
func (l *Load) GetTargets() []int {
	return l.targetRegisters
}

// Get the source registers of the instruction.
func (l *Load) GetSources() []int {
	return l.sourceRegisters
}

// Get the immediate value of the instruction.
func (l *Load) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (l *Load) GetLabel() string {
	return l.blockLabel
}

// Set the label that marks this instruction in code.
func (l *Load) SetLabel(newLabel string) {
	l.blockLabel = newLabel
}
