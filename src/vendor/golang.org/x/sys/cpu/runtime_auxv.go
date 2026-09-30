package cpu

// getAuxvFn is non-nil on Go 1.21+ (via runtime_auxv_go121.go init)
// on platforms that use auxv.
var getAuxvFn func() []uintptr

func getAuxv() []uintptr {
	if getAuxvFn == nil {
		return nil
	}
	return getAuxvFn()
}
