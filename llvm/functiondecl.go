package llvm

import (
	"bytes"
	"golite/arm"
	"golite/stack"
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
	for _, block := range f.blocks {
		block.BuildStackTable(f.name, stack)
	}
}

// Convert the LLVM IR to ARM assembly
func (f *FunctionDecl) ToARM() *arm.FunctionDecl {
	blocks := make([]*arm.BasicBlock, 0)
	for _, block := range f.blocks {
		blocks = append(blocks, block.ToARM())
	}
	return arm.NewFunctionDecl(f.name, blocks)
}
