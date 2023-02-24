package llvm

import "bytes"

// The LLVM program node
type Program struct {
	structDecls   []StructDecl
	globalDecls   []GlobalDecl
	functionDecls []FunctionDecl
	runtimeDecls  []RuntimeDecl
}

// Create a new LLVM program node
func NewProgram(structDecls []StructDecl, globalDecls []GlobalDecl, functionDecls []FunctionDecl, runtimeDecls []RuntimeDecl) *Program {
	return &Program{structDecls, globalDecls, functionDecls, runtimeDecls}
}

// String representation of the LLVM program node
func (p *Program) String() string {
	// Print the struct declarations
	var out bytes.Buffer
	for _, strct := range p.structDecls {
		out.WriteString(strct.String())
		out.WriteString("\n")
	}

	// Print the global declarations
	for _, global := range p.globalDecls {
		out.WriteString(global.String())
		out.WriteString("\n")
	}

	// Print the function declarations
	for _, fn := range p.functionDecls {
		out.WriteString(fn.String())
		out.WriteString("\n")
	}

	// Print the runtime declarations
	for _, runtime := range p.runtimeDecls {
		out.WriteString(runtime.String())
		out.WriteString("\n")
	}

	return out.String()
}

// Combines the target information and the LLVM program into a single object
func GetLLVMSource(targetInfo *TargetInformation, program *Program) string {
	var llvmRepr bytes.Buffer
	llvmRepr.WriteString(targetInfo.String())
	llvmRepr.WriteString(program.String())
	llvmRepr.WriteString("\n")
	return llvmRepr.String()
}
