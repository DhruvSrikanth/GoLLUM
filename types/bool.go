package types

// BoolType is a type that represents a boolean
type BoolType struct{}

// Singleton instance of BoolType
var singletonBool *BoolType

// String representation of BoolType
func (boolTy *BoolType) String() string {
	return "bool"
}

// Get the singleton instance of BoolType (thread safe)
func getBoolInstance() *BoolType {
	if singletonBool == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonBool == nil {
			// This thread is the first to create the singleton
			singletonBool = &BoolType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonBool
}
