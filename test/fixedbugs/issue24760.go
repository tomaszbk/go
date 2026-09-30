// compile


package p

import "unsafe"

var _ = string([]byte(nil))[0]
var _ = uintptr(unsafe.Pointer(uintptr(1))) << 100
