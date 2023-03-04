package arm

// The ARM program node
type Program struct {
	globalDecls   []GlobalDecl
	functionDecls []FunctionDecl
	constantDecls []ConstantDecl
}

// Create a new ARM program node
func NewProgram(globalDecls []GlobalDecl, functionDecls []FunctionDecl, constantDecls []ConstantDecl) *Program {
	return &Program{globalDecls, functionDecls, constantDecls}
}

// String representation of the ARM program node
func (p *Program) String() string {
	// Print the struct declarations
	var out string

	out += "\t.arch armv8-a\n"
	out += "\n"

	// Print the global declarations
	for _, global := range p.globalDecls {
		out += "\t" + global.String() + "\n"
	}
	out += "\n"

	out += "\t.text\n"
	out += "\n"

	// Get the symbol table for the function
	// Print the function declarations
	for _, fn := range p.functionDecls {
		out += fn.String()
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
