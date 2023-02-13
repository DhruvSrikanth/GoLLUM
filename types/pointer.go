package types

// PointerType is a type that represents a nil value
type PointerType struct {
	ptrType string
}

// Singleton instance of PointerType
var singletonPointerTypes = make(map[string]*PointerType)

// String representation of NilType
func (ptrTy *PointerType) String() string {
	return ptrTy.ptrType
}

// Get the singleton instance of PointerType (thread safe)
func getPointerInstance(ptrType string) *PointerType {
	if singletonPointerTypes[ptrType] == nil {
		// No thread has created the singleton yet
		lock.Lock()
		defer lock.Unlock()
		if singletonPointerTypes[ptrType] == nil {
			// This thread is the first to create the singleton
			singletonPointerTypes[ptrType] = &PointerType{ptrType}
		}
		// Now the singleton is created so other theads waiting can use it
	}
	// Singleton has already been created
	return singletonPointerTypes[ptrType]
}
