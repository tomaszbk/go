// compile

// Issue 4614: slicing of nil slices confuses the compiler
// with a uintptr(nil) node.

package p

import "unsafe"

var n int

var _ = []int(nil)[1:]
var _ = []int(nil)[n:]

var _ = uintptr(unsafe.Pointer(nil))
var _ = unsafe.Pointer(uintptr(0))
