package llvm

import (
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
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
	var out string
	// Format is - call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr1, i32 0, i32 0), i64 %r15, i64 %r18, i64 %r21)
	// Common for all
	out += "call i32 (i8*, ...) @printf(i8* getelementptr inbounds (["
	out += strconv.Itoa(p.size)
	out += " x i8], ["
	out += strconv.Itoa(p.size)
	out += " x i8]* @."
	out += p.varName
	out += ", i32 0, i32 0)"
	for _, r := range p.sourceRegString {
		out += ", "
		out += "i64 "
		out += r
	}
	out += ")"
	return out
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

// Convert LLVM IR to ARM assembly.
func (p *Printf) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)

	// Compute the address of all the source registers and store them in registers
	availableRegNum := stackFrame.GetNextRegister()
	tempRegs := make([]string, 0)
	for _, r := range p.sourceRegisters {
		reg := "r" + strconv.Itoa(r)
		argumentAddress := stackFrame.GetLocation(reg)
		availableReg := "x" + strconv.Itoa(availableRegNum)

		// If the argument is in a register, we move it, otherwise we load it
		if strings.Contains(argumentAddress, "x") {
			movInst := arm.NewMov(availableReg, argumentAddress)
			movInst.SetLabel(p.blockLabel)
			insts = append(insts, movInst)
		} else {
			ldrInst := arm.NewLdr(availableReg, "sp, #"+argumentAddress)
			ldrInst.SetLabel(p.blockLabel)
			insts = append(insts, ldrInst)
		}

		tempRegs = append(tempRegs, availableReg)
		availableRegNum += 1
	}

	// Move the arguments from temporary registers to the correct registers i.e. x1, x2, x3 ...
	for i, r := range tempRegs {
		movInst := arm.NewMov("x"+strconv.Itoa(i+1), r)
		movInst.SetLabel(p.blockLabel)
		insts = append(insts, movInst)
	}

	// Compute the address of the string using the global label
	// and store it in the next available register
	availableReg := "x" + strconv.Itoa(availableRegNum)
	adrpInst := arm.NewAdrp(availableReg, "."+strings.ToUpper(p.varName))
	adrpInst.SetLabel(p.blockLabel)
	insts = append(insts, adrpInst)
	// Add the offset to the address
	addInst := arm.NewAdd(availableReg, availableReg, ":lo12:."+strings.ToUpper(p.varName))
	addInst.SetLabel(p.blockLabel)
	insts = append(insts, addInst)

	// Move the address of the string to the first argument register
	movInst := arm.NewMov("x0", availableReg)
	movInst.SetLabel(p.blockLabel)
	insts = append(insts, movInst)

	// Call the printf function
	blInst := arm.NewBranch("printf")
	blInst.SetLabel(p.blockLabel)
	insts = append(insts, blInst)

	return insts
}

// Build the stack table for the instruction.
func (p *Printf) BuildStackTable(fName string, stack *stack.Stack) {
}
