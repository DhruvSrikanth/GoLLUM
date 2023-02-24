package llvm

// Representation of a runtime declaration
type RuntimeDecl struct {
	// The name of the runtime function
	name string
	// The LLVM representation of the runtime function
	llvm string
}

// NewRuntimeDecl returns a new runtime declaration
func NewRuntimeDecl(name string, llvm string) *RuntimeDecl {
	return &RuntimeDecl{name, llvm}
}

// String representation of the runtime declaration
func (r *RuntimeDecl) String() string {
	return "declare " + r.name
}
