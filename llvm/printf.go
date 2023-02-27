package llvm

import (
	"bytes"
	"strconv"
)

// Representation of a Printf runtime call
type Printf struct {
	varName         string
	sourceRegString []string
	size            int
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewPrintf returns a new Printf runtime call
func NewPrintf(varName string, sourceR []string, size int) *Printf {
	srcR := make([]int, 0)
	for _, r := range sourceR {
		rInt, _ := strconv.Atoi(r[2:])
		srcR = append(srcR, rInt)
	}

	tgtR := make([]int, 0)

	return &Printf{
		varName:         varName,
		sourceRegString: sourceR,
		size:            size,
		blockLabel:      "",
		sourceRegisters: srcR,
		targetRegisters: tgtR,
	}
}

// String representation of the Printf call
func (p *Printf) String() string {
	var out bytes.Buffer
	// Format is - call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr1, i32 0, i32 0), i64 %r15, i64 %r18, i64 %r21)
	// Common for all
	out.WriteString("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([")
	out.WriteString(strconv.Itoa(p.size))
	out.WriteString(" x i8], [")
	out.WriteString(strconv.Itoa(p.size))
	out.WriteString(" x i8]* @.")
	out.WriteString(p.varName)
	out.WriteString(", i32 0, i32 0)")
	for _, r := range p.sourceRegString {
		out.WriteString(", ")
		out.WriteString("i64 ")
		out.WriteString(r)
	}
	out.WriteString(")")
	return out.String()
}

// Get the registers targeted by the instruction.
func (p *Printf) GetTargets() []int {
	return p.targetRegisters
}

// Get the source registers of the instruction.
func (p *Printf) GetSources() []int {
	return p.sourceRegisters
}

// Get the immediate value of the instruction.
func (p *Printf) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (p *Printf) GetLabel() string {
	return p.blockLabel
}

// Set the label that marks this instruction in code.
func (p *Printf) SetLabel(newLabel string) {
	p.blockLabel = newLabel
}
