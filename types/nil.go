package types

// NilType is a type that represents a nil value
type NilType struct{}

// Singleton instance of NilType
var singletonNil *NilType

// String representation of NilType
func (nilTy *NilType) String() string {
	return "nil"
}

// Get the singleton instance of NilType (thread safe)
func getNilInstance() *NilType {
	if singletonNil == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonNil == nil {
			// This thread is the first to create the singleton
			singletonNil = &NilType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonNil
}
