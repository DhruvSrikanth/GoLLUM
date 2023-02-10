package types

// Void type
type VoidType struct{}

// Singleton instance of VoidType
var singletonVoid *VoidType

// String representation of VoidType
func (voidTy *VoidType) String() string {
	return "void"
}

// Get the singleton instance of VoidType (thread safe)
func getVoidInstance() *VoidType {
	if singletonVoid == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonVoid == nil {
			// This thread created the singleton
			singletonVoid = &VoidType{}
		}
		// Now the singleton is created, other threads can use it
	}
	// Singleton has already been created
	return singletonVoid
}
