package arm

import (
	"bytes"
)

// Representation of a function declaration
type FunctionDecl struct {
	// The name of the function declaration
	name   string
	blocks []*BasicBlock
}

// NewFunctionDecl returns a new function declaration
func NewFunctionDecl(name string, blocks []*BasicBlock) *FunctionDecl {
	return &FunctionDecl{name, blocks}
}

// String representation of the function declaration
func (f *FunctionDecl) String() string {
	var out bytes.Buffer
	// Format is .type <name>, %function
	//           .global <name>
	// 		 	 .p2align 2
	//           <name>:
	//           <blocks>
	// 		     .size <name>, (.-<name>)
	out.WriteString("\t.type " + f.name + ", %" + "function\n")
	out.WriteString("\t.global " + f.name + "\n")
	out.WriteString("\t.p2align 2\n")
	out.WriteString(f.name + ":\n")
	// for _, block := range f.blocks {
	// 	out.WriteString(block.String())
	// out.WritrString("\n")
	// }
	out.WriteString("\t.size " + f.name + ", (.-" + f.name + ")\n")

	return out.String()
}
