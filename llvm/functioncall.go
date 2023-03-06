package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a function call
type FunctionCall struct {
	fName           string
	retTy           string
	registerNames   []string
	sourceTypes     []string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// NewFunctionCall returns a new function call
func NewFunctionCall(fName string, retTy string, srcReg []string, srcTy []string) *FunctionCall {
	registerNames := srcReg
	srcR := make([]int, 0)
	// Remove the %r from the register names
	for _, r := range srcReg {
		if strings.Contains(r, "%r") {
			r = strings.Replace(r, "%r", "", -1)
			i, _ := strconv.Atoi(r)
			srcR = append(srcR, i)
		}
	}

	tgtR := make([]int, 0)
	nextR, _ := strconv.Atoi(GetNextRegister()[2:])
	tgtR = append(tgtR, nextR)
	return &FunctionCall{fName, retTy, registerNames, srcTy, "", srcR, tgtR}
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
		out.WriteString(" " + f.registerNames[i])
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

// Convert to ARM assembly.
func (f *FunctionCall) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	// Push all the arguments onto the stack but putting them in registers
	// Call the function
	// Place the return value which is in x0 into the destination register
	insts := make([]arm.Instruction, 0)

	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	tempRegs := make([]string, 0)
	// Move the arguments into temporary registers
	for i := 0; i < len(f.registerNames); i++ {
		regName := f.registerNames[i]
		availableReg := "x" + strconv.Itoa(availableRegNum)
		if strings.Contains(regName, "@") {
			// Global variable
			adrpInst := arm.NewAdrp(availableReg, regName[1:])
			adrpInst.SetLabel(f.blockLabel)
			insts = append(insts, adrpInst)

			addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+regName[1:])
			addInst.SetLabel(f.blockLabel)
			insts = append(insts, addInst)

			ldrInst := arm.NewLdr(availableReg, availableReg)
			ldrInst.SetLabel(f.blockLabel)
			insts = append(insts, ldrInst)

		} else if strings.Contains(regName, "%") {
			srcAddr := stackFrame.GetLocation(regName[1:])
			if strings.Contains(srcAddr, "x") {
				// Register
				movInst := arm.NewMov(availableReg, srcAddr)
				movInst.SetLabel(f.blockLabel)
				insts = append(insts, movInst)
			} else {
				// Address on the stack
				ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+srcAddr)
				ldrInst.SetLabel(f.blockLabel)
				insts = append(insts, ldrInst)
			}
		} else {
			// Immediate
			// Check if its negative
			if strings.Contains(regName, "-") {
				movInst := arm.NewMov(availableReg, "#0")
				movInst.SetLabel(f.blockLabel)
				insts = append(insts, movInst)

				subInst := arm.NewSub(availableReg, availableReg, "#"+regName[1:])
				subInst.SetLabel(f.blockLabel)
				insts = append(insts, subInst)
			} else {
				movInst := arm.NewMov(availableReg, "#"+regName)
				movInst.SetLabel(f.blockLabel)
				insts = append(insts, movInst)
			}
		}

		tempRegs = append(tempRegs, availableReg)
		availableRegNum += 1
	}

	// Move the arguments into the correct registers
	for i := 0; i < len(tempRegs); i++ {
		movInst := arm.NewMov("x"+strconv.Itoa(i), tempRegs[i])
		movInst.SetLabel(f.blockLabel)
		insts = append(insts, movInst)
	}

	// Call the function
	callInst := arm.NewBranch(f.fName)
	callInst.SetLabel(f.blockLabel)
	insts = append(insts, callInst)

	// Move the return value into the destination register
	retReg := "x0"
	destReg := "r" + strconv.Itoa(f.targetRegisters[len(f.targetRegisters)-1])
	destAddr := stackFrame.GetLocation(destReg)
	if strings.Contains(destAddr, "x") {
		// Register
		movInst := arm.NewMov(destAddr, retReg)
		movInst.SetLabel(f.blockLabel)
		insts = append(insts, movInst)
	} else {
		// Address on the stack
		strInst := arm.NewStr(retReg, arm.SP+", #"+destAddr)
		strInst.SetLabel(f.blockLabel)
		insts = append(insts, strInst)
	}

	return insts
}

// Build the stack table for the instruction.
func (f *FunctionCall) BuildStackTable(funcName string, stack *stack.Stack) {
	destinationReg := "r" + strconv.Itoa(f.targetRegisters[len(f.targetRegisters)-1])
	stack.AddEntry(funcName, destinationReg, strconv.Itoa(stack.GetFrame(funcName).GetLargestOffset()+8))
}
