package stack

// Represents a an entry in the stack
type StackEntry struct {
	// The location of the entry
	location string
	// The type of the entry ("value" or "pointer")
	ty string
}

// Create a new stack entry
func NewStackEntry(location string, ty string) *StackEntry {
	return &StackEntry{location, ty}
}

// Get the location of the entry
func (s *StackEntry) GetLocation() string {
	return s.location
}

// Get the type of the entry
func (s *StackEntry) GetType() string {
	return s.ty
}
