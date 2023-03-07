package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
	"strconv"
	"strings"
)

// Representation of a function declaration
type FunctionDecl struct {
	// The name of the function declaration
	name       string
	ty         string
	params     []string
	paramTypes []string
	blocks     []*BasicBlock
}

// NewFunctionDecl returns a new function declaration
func NewFunctionDecl(name string, ty string, params, paramTypes []string, blocks []*BasicBlock) *FunctionDecl {
	return &FunctionDecl{name, ty, params, paramTypes, blocks}
}

// String representation of the function declaration
func (f *FunctionDecl) String() string {
	var out bytes.Buffer
	// Format is define <return type> @<function name>(<ty> %<param_name1>, <ty> %<param_name2>) \n{\n <blocks> \n}\n
	out.WriteString("define ")
	if strings.Contains(f.ty, "struct.") {
		out.WriteString("%")
	}
	out.WriteString(f.ty)
	if strings.Contains(f.ty, "struct.") {
		out.WriteString("*")
	}
	out.WriteString(" @")
	out.WriteString(f.name)
	out.WriteString("(")
	for i, param := range f.params {
		paramType := f.paramTypes[i]
		if strings.Contains(paramType, "struct.") {
			out.WriteString("%")
		}
		out.WriteString(paramType)
		if strings.Contains(paramType, "struct.") {
			out.WriteString("*")
		}
		out.WriteString(" %")
		out.WriteString(param)
		if i < len(f.params)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n{\n")
	for _, block := range f.blocks {
		out.WriteString(block.String())
	}
	out.WriteString("\n}\n")

	return out.String()
}

// Build the stack table for each function
func (f *FunctionDecl) BuildStackTable(stack *stack.Stack) {
	// Add stack frame for the function
	stack.AddFrame(f.name)
	// Add fp to the stack
	stack.AddEntry(f.name, arm.SP, "0")
	// Add lr to the stack
	stack.AddEntry(f.name, arm.LR, "8")
	for _, block := range f.blocks {
		block.BuildStackTable(f.name, stack)
	}
	// Add the parameter register to the stack
	for _, param := range f.params {
		stack.AddEntry(f.name, "P_"+param, strconv.Itoa(stack.GetFrame(f.name).GetLargestOffset()+8))
	}
	// Add the parameter itself if it is greater than the number of avaliable registers
	nAvailableRegisters := 8
	for i, param := range f.params {
		if i >= nAvailableRegisters {
			stack.AddEntry(f.name, param, strconv.Itoa(stack.GetFrame(f.name).GetLargestOffset()+8))
		} else {
			stack.AddEntry(f.name, param, "x"+strconv.Itoa(i))
		}
	}
}

// Convert the LLVM IR to ARM assembly
func (f *FunctionDecl) ToARM(stack *stack.Stack) *arm.FunctionDecl {
	blocks := make([]*arm.BasicBlock, 0)
	isFirstBlock := true

	for _, block := range f.blocks {
		blocks = append(blocks, block.ToARM(f.name, stack, isFirstBlock))
		if isFirstBlock {
			isFirstBlock = false
		}
	}

	stackFrame := stack.GetFrame(f.name)

	// Add the conditional branches to the end of the block if they are cfg entry blocks
	for i, armBlock := range blocks {
		llvmBlock := f.blocks[i]
		availableRegNum := stackFrame.GetNextRegister()
		if llvmBlock.IsIfEntry() {
			var op1R, op2R string

			availableReg := "x" + strconv.Itoa(availableRegNum)
			availableRegNum += 1

			// Add an instruction to branch
			condBranchLLVMInst := llvmBlock.GetLastInstruction().(*BranchConditional)

			condReg := condBranchLLVMInst.GetConditionRegister()
			condAddr := stackFrame.GetLocation(condReg[1:])
			if strings.Contains(condAddr, "x") {
				// Already in register
				op1R = condAddr
			} else {
				// In an address on the stack
				ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+condAddr)
				ldrInst.SetLabel(llvmBlock.GetLabel())
				armBlock.AddInstruction(ldrInst)

				op1R = availableReg

			}

			availableReg = "x" + strconv.Itoa(availableRegNum)
			availableRegNum += 1
			// Move #0 to a register because we want to check if the condition is false
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(llvmBlock.GetLabel())
			armBlock.AddInstruction(movInst)

			op2R = availableReg

			// Compare the two registers
			cmpInst := arm.NewCmp(op1R, op2R)
			cmpInst.SetLabel(llvmBlock.GetLabel())
			armBlock.AddInstruction(cmpInst)

			// Branch to the false label if the condition is true i.e. icmp condition is false == 0
			falseLabel := "." + condBranchLLVMInst.GetFalseLabel()
			bInst := arm.NewBranch(falseLabel)
			bInst.SetLabel(llvmBlock.GetLabel())
			bInst.SetBranchType("conditional")
			bInst.SetConditionFlag("eq")
			armBlock.AddInstruction(bInst)
		} else if llvmBlock.IsForEntry() {
			var op1R, op2R string

			availableReg := "x" + strconv.Itoa(availableRegNum)
			availableRegNum += 1

			// Add an instruction to branch
			condBranchLLVMInst := llvmBlock.GetLastInstruction().(*BranchConditional)

			condReg := condBranchLLVMInst.GetConditionRegister()
			condAddr := stackFrame.GetLocation(condReg[1:])
			if strings.Contains(condAddr, "x") {
				// Already in register
				op1R = condAddr
			} else {
				// In an address on the stack
				ldrInst := arm.NewLdr(availableReg, arm.SP+", #"+condAddr)
				ldrInst.SetLabel(llvmBlock.GetLabel())
				armBlock.AddInstruction(ldrInst)

				op1R = availableReg

			}

			availableReg = "x" + strconv.Itoa(availableRegNum)
			availableRegNum += 1
			// Move #0 to a register because we want to check if the condition is false
			movInst := arm.NewMov(availableReg, "#0")
			movInst.SetLabel(llvmBlock.GetLabel())
			armBlock.AddInstruction(movInst)

			op2R = availableReg

			// Compare the two registers
			cmpInst := arm.NewCmp(op1R, op2R)
			cmpInst.SetLabel(llvmBlock.GetLabel())
			armBlock.AddInstruction(cmpInst)

			// Branch to the false label if the condition is true i.e. icmp condition is false == 0
			falseLabel := "." + condBranchLLVMInst.GetFalseLabel()
			bInst := arm.NewBranch(falseLabel)
			bInst.SetLabel(llvmBlock.GetLabel())
			bInst.SetBranchType("conditional")
			bInst.SetConditionFlag("eq")
			armBlock.AddInstruction(bInst)
		}
	}

	return arm.NewFunctionDecl(f.name, blocks)
}
