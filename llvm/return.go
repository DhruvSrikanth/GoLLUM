package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"golite/utils"
	"strconv"
	"strings"
)

// Represents a return instruction in LLVM IR
type Return struct {
	sourceRegister  string
	ty              string
	blockLabel      string
	sourceRegisters []int
	targetRegisters []int
}

// Create a new return instruction
func NewReturn(sourceRegister, ty string) *Return {
	srcR := make([]int, 0)
	targetR := make([]int, 0)
	return &Return{sourceRegister, ty, "", srcR, targetR}
}

// String representation of the return instruction
func (r *Return) String() string {
	// Output format is - ret <type> %<register>

	var out bytes.Buffer

	out.WriteString("ret ")

	// struct type
	if strings.Contains(r.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(r.ty)

	out.WriteString(" ")
	out.WriteString(r.sourceRegister)

	return out.String()
}

// Get the registers targeted by the instruction.
func (r *Return) GetTargets() []int {
	return r.targetRegisters
}

// Get the source registers of the instruction.
func (r *Return) GetSources() []int {
	return r.sourceRegisters
}

// Get the immediate value of the instruction.
func (r *Return) GetImmediate() int {
	return 0
}

// Get the label that marks this instruction in code.
func (r *Return) GetLabel() string {
	return r.blockLabel
}

// Set the label that marks this instruction in code.
func (r *Return) SetLabel(newLabel string) {
	r.blockLabel = newLabel
}

// Convert LLVM IR to ARM assembly.
func (r *Return) ToARM(fnName string, stack *stack.Stack) []arm.Instruction {
	insts := make([]arm.Instruction, 0)

	// Store the return value in x0
	stackFrame := stack.GetFrame(fnName)
	availableRegNum := stackFrame.GetNextRegister()
	availableReg := "x" + strconv.Itoa(availableRegNum)
	if strings.Contains(r.sourceRegister, "@") {
		adrpInst := arm.NewAdrp(availableReg, r.sourceRegister[1:])
		adrpInst.SetLabel(r.blockLabel)
		insts = append(insts, adrpInst)

		addInst := arm.NewAdd(availableReg, availableReg, ":lo12:"+r.sourceRegister[1:])
		addInst.SetLabel(r.blockLabel)
		insts = append(insts, addInst)

		ldrInst := arm.NewLdr(availableReg, availableReg)
		ldrInst.SetLabel(r.blockLabel)
		insts = append(insts, ldrInst)
	} else if strings.Contains(r.sourceRegister, "%") {
		srcAddr := stackFrame.GetLocation(r.sourceRegister[1:])
		if strings.Contains(srcAddr, "x") {
			availableReg = srcAddr
		} else {
			ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+srcAddr)
			ldrInst.SetLabel(r.blockLabel)
			insts = append(insts, ldrInst)
		}
	} else {
		// If the immediate value is negative, move 0 into a register and then subtract the value from it
		if strings.Contains(r.sourceRegister, "-") {
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(r.blockLabel)
			insts = append(insts, movInst)

			// Max value is 4095, so we can subtract the value from 0
			num, _ := strconv.Atoi(r.sourceRegister[1:])

			availableRegNum += 1
			nextReg := "x" + strconv.Itoa(availableRegNum)
			hexNum := utils.IntToHex(int64(num))
			hexSplit := utils.SplitHex(hexNum)
			// Move the upper 16 bits into the register and shift left by 48 bits
			movzInst := arm.NewMovz(nextReg, "#"+hexSplit[0]+", lsl #48")
			movzInst.SetLabel(r.blockLabel)
			insts = append(insts, movzInst)

			// Move the next 16 bits into the register and shift left by 32 bits
			movkInst := arm.NewMovk(nextReg, "#"+hexSplit[1]+", lsl #32")
			movkInst.SetLabel(r.blockLabel)
			insts = append(insts, movkInst)

			// Move the next 16 bits into the register and shift left by 16 bits
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[2]+", lsl #16")
			movkInst.SetLabel(r.blockLabel)
			insts = append(insts, movkInst)

			// Move the last 16 bits into the register
			movkInst = arm.NewMovk(nextReg, "#"+hexSplit[3]+", lsl #0")
			movkInst.SetLabel(r.blockLabel)
			insts = append(insts, movkInst)

			subInst := arm.NewSub(availableReg, availableReg, nextReg)
			subInst.SetLabel(r.blockLabel)
			insts = append(insts, subInst)
		} else {
			if num, _ := strconv.Atoi(r.sourceRegister); num > 4095 {

				hexNum := utils.IntToHex(int64(num))
				hexSplit := utils.SplitHex(hexNum)
				// Move the upper 16 bits into the register and shift left by 48 bits
				movzInst := arm.NewMovz(availableReg, "#"+hexSplit[0]+", lsl #48")
				movzInst.SetLabel(r.blockLabel)
				insts = append(insts, movzInst)

				// Move the next 16 bits into the register and shift left by 32 bits
				movkInst := arm.NewMovk(availableReg, "#"+hexSplit[1]+", lsl #32")
				movkInst.SetLabel(r.blockLabel)
				insts = append(insts, movkInst)

				// Move the next 16 bits into the register and shift left by 16 bits
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[2]+", lsl #16")
				movkInst.SetLabel(r.blockLabel)
				insts = append(insts, movkInst)

				// Move the last 16 bits into the register
				movkInst = arm.NewMovk(availableReg, "#"+hexSplit[3]+", lsl #0")
				movkInst.SetLabel(r.blockLabel)
				insts = append(insts, movkInst)
			} else {
				// If the value is less than 4095, we can just move it into x0
				movInst := arm.NewMov(availableReg, "#"+r.sourceRegister)
				movInst.SetLabel(r.blockLabel)
				insts = append(insts, movInst)
			}
		}
	}
	movInst := arm.NewMov("x0", availableReg)
	movInst.SetLabel(r.blockLabel)
	insts = append(insts, movInst)

	// Any time there is a return instruction, we need to have the epilogue
	// Restore the frame pointer and link register
	ldpInst := arm.NewLdp("x29", "x30", arm.SP)
	ldpInst.SetLabel(r.blockLabel)
	insts = append(insts, ldpInst)

	// Restore the stack pointer
	addInst := arm.NewAdd(arm.SP, arm.SP, "#16")
	addInst.SetLabel(r.blockLabel)
	insts = append(insts, addInst)

	// Add the amount that we grew the stack by previously
	growBy := (stackFrame.Size() - 2) * 8
	// Round to the nearest multiple of 16 for alignment.
	if growBy%16 != 0 {
		growBy += 16 - (growBy % 16)
	}

	addInst = arm.NewAdd(arm.SP, arm.SP, "#"+strconv.Itoa(growBy))
	addInst.SetLabel(r.blockLabel)
	insts = append(insts, addInst)

	// Return
	retInst := arm.NewRet()
	retInst.SetLabel(r.blockLabel)
	insts = append(insts, retInst)

	return insts
}

// Build the stack table for the instruction.
func (r *Return) BuildStackTable(fName string, stack *stack.Stack) {
}
