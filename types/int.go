package types

// IntType is a type that represents an integer
type IntType struct{}

// Singleton instance of IntType
var singletonInt *IntType

// String representation of IntType
func (intTy *IntType) String() string {
	return "int"
}

// Get the singleton instance of IntType (thread safe)
func getIntInstance() *IntType {
	if singletonInt == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonInt == nil {
			// This thread is the first to create the singleton
			singletonInt = &IntType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonInt
}
