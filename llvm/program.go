package llvm

import "bytes"

// The LLVM program node
type Program struct {
	structDecls   []StructDecl
	globalDecls   []GlobalDecl
	functionDecls []FunctionDecl
	runtimeDecls  []RuntimeDecl
	constantDecls []ConstantDecl
}

// Create a new LLVM program node
func NewProgram(structDecls []StructDecl, globalDecls []GlobalDecl, functionDecls []FunctionDecl, runtimeDecls []RuntimeDecl, constantDecls []ConstantDecl) *Program {
	return &Program{structDecls, globalDecls, functionDecls, runtimeDecls, constantDecls}
}

// String representation of the LLVM program node
func (p *Program) String() string {
	// Print the struct declarations
	var out string
	for _, strct := range p.structDecls {
		out += strct.String()
		out += "\n"
	}

	out += "\n"

	// Print the global declarations
	for _, global := range p.globalDecls {
		out += global.String()
		out += "\n"
	}

	out += "\n"

	// Print the function declarations
	for _, fn := range p.functionDecls {
		out += fn.String()
		out += "\n"
	}

	out += "\n"

	// Print the runtime declarations
	for _, runtime := range p.runtimeDecls {
		out += runtime.String()
		out += "\n"
	}

	out += "\n"

	// Print the constant declarations
	for _, constant := range p.constantDecls {
		out += constant.String()
		out += "\n"
	}

	return out
}

// Combines the target information and the LLVM program into a single object
func GetLLVMSource(targetInfo *TargetInformation, program *Program) string {
	var llvmRepr bytes.Buffer
	llvmRepr.WriteString(targetInfo.String())
	llvmRepr.WriteString(program.String())
	llvmRepr.WriteString("\n")
	return llvmRepr.String()
}
