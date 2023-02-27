package llvm

import (
	"bytes"
	"strconv"
	"strings"
)

// Representation of a function call
type FunctionCall struct {
	fName           string
	retTy           string
	sourceTypes     []string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewFunctionCall returns a new function call
func NewFunctionCall(fName string, retTy string, srcReg []string, srcTy []string) *FunctionCall {
	srcR := make([]int, 0)
	// Remove the %r from the register names
	for _, r := range srcReg {
		r = strings.Replace(r, "%r", "", -1)
		i, _ := strconv.Atoi(r)
		srcR = append(srcR, i)
	}

	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &FunctionCall{fName, retTy, srcTy, "", srcR, tgtR}
}

// String representation of the FunctionCall
func (f *FunctionCall) String() string {
	var out bytes.Buffer
	// Format is - %result = call type @fName(ty %reg_var)
	// Common for all
	out.WriteString("%r" + strconv.Itoa(f.targetRegisters[len(f.targetRegisters)-1]))
	out.WriteString(" = call ")

	// Check if its a struct type
	if strings.Contains(f.retTy, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(f.retTy)

	out.WriteString(" @" + f.fName + "(")
	for i := 0; i < len(f.sourceTypes); i++ {
		// Check if its a struct type
		if strings.Contains(f.sourceTypes[i], "struct.") {
			out.WriteString("%")
		}
		out.WriteString(f.sourceTypes[i])
		out.WriteString(" %r" + strconv.Itoa(f.sourceRegisters[i]))
		if i < len(f.sourceTypes)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String()
}

// Get the registers targeted by the instruction.
func (f *FunctionCall) GetTargets() []int {
	return f.targetRegisters
}

// Get the source registers of the instruction.
func (f *FunctionCall) GetSources() []int {
	return f.sourceRegisters
}

// Get the immediate value of the instruction.
func (f *FunctionCall) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (f *FunctionCall) GetLabel() string {
	return f.blockLabel
}

// Set the label that marks this instruction in code.
func (f *FunctionCall) SetLabel(newLabel string) {
	f.blockLabel = newLabel
}
