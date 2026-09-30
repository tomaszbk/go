package scan

// FilterNilAVX512 is the simd version of FilterNil,
// it is implemented in assembly.
func FilterNilAVX512(bufp *uintptr, n int32) int32
