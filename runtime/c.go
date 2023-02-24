package runtime

type runtime struct {
	cFuncs map[string]string
}

// NewRuntime returns a new runtime
func NewRuntime() *runtime {
	cFuncs := make(map[string]string)
	cFuncs["malloc"] = "i8* @malloc(i32)"
	cFuncs["free"] = "void @free(i8*)"
	cFuncs["printf"] = "i32 @printf(i8*, ...)"
	cFuncs["scanf"] = "i32 @scanf(i8*, ...)"
	return &runtime{cFuncs: cFuncs}
}

// GetCFuncs returns a particular C function used by the runtime
func (r *runtime) GetCFunc(funcKey string) string {
	return r.cFuncs[funcKey]
}

// GetCFuncs returns the C functions used by the runtime
func (r *runtime) GetCFuncs() []string {
	fNames := make([]string, 0)
	for _, name := range r.cFuncs {
		fNames = append(fNames, name)
	}
	return fNames
}
