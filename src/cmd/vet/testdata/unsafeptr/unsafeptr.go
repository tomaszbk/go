package unsafeptr

import "unsafe"

func _() {
	var x unsafe.Pointer
	var y uintptr
	x = unsafe.Pointer(y) // ERROR "possible misuse of unsafe.Pointer"
	_ = x
}
