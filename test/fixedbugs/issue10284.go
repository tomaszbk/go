// compile


// Issue 10284: gccgo failed to allow converting a user-defined
// type whose underlying type is uintptr to unsafe.Pointer.

package p

import "unsafe"

type T uintptr

var _ unsafe.Pointer = unsafe.Pointer(T(0))
