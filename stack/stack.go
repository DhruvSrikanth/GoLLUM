package stack

type Stack struct {
	// A collection of stack frames
	frames map[string]*StackFrame
}

// Create a new stack
func NewStack() *Stack {
	return &Stack{make(map[string]*StackFrame)}
}

// Add a new stack frame to the stack
func (s *Stack) AddFrame(name string) {
	s.frames[name] = NewStackFrame(name)
}

// Add a new entry to the stack frame
func (s *Stack) AddEntry(frameName string, name string, location string, ty string) {
	s.frames[frameName].AddEntry(name, location, ty)
}

// Get the stack frame
func (s *Stack) GetFrame(name string) *StackFrame {
	return s.frames[name]
}

// String representation of the stack
func (s *Stack) String() string {
	var out string
	for _, frame := range s.frames {
		out += "------------------\n"
		out += frame.String()
		out += "------------------\n"
	}
	return out
}
