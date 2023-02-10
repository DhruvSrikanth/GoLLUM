package types

// StringType is a type that represents a string
type StringType struct{}

// Singleton instance of StringType
var singletonString *StringType

// String representation of StringType
func (stringTy *StringType) String() string {
	return "string"
}

// Get the singleton instance of StringType (thread safe)
func getStringInstance() *StringType {
	if singletonString == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonString == nil {
			// This thread is the first to create the singleton
			singletonString = &StringType{}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonString
}
