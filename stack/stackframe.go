package stack

type StackFrame struct {
	// The name of the stack table (corresponds to function name)
	name string
	// The stack table
	table map[string]string
}

// Create a new stack table
func NewStackFrame(name string) *StackFrame {
	return &StackFrame{name, make(map[string]string)}
}

// Add a new entry to the stack table
func (s *StackFrame) AddEntry(name string, location string) {
	s.table[name] = location
}

// String representation of the stack table
func (s *StackFrame) String() string {
	var out string
	out += "Stack Frame: " + s.name + "\n"
	for name, location := range s.table {
		out += name + "->" + location + "\n"
	}
	return out
}
