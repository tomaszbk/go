package race_test

import "unsafe"

// golang.org/issue/12225
// The test is that this compiles at all.

//go:noinline
func convert(s string) []byte {
	return []byte(s)
}

func issue12225() {
	println(*(*int)(unsafe.Pointer(&convert("")[0])))
	println(*(*int)(unsafe.Pointer(&[]byte("")[0])))
}
