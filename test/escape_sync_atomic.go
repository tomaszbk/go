// errorcheck -0 -m -l


// Test escape analysis for sync/atomic.

package escape

import (
	"sync/atomic"
	"unsafe"
)

// BAD: should be "leaking param: addr to result ~r1 level=1$".
func LoadPointer(addr *unsafe.Pointer) unsafe.Pointer { // ERROR "leaking param: addr$"
	return atomic.LoadPointer(addr)
}

var ptr unsafe.Pointer

func StorePointer() {
	var x int // ERROR "moved to heap: x"
	atomic.StorePointer(&ptr, unsafe.Pointer(&x))
}

func SwapPointer() {
	var x int // ERROR "moved to heap: x"
	atomic.SwapPointer(&ptr, unsafe.Pointer(&x))
}

func CompareAndSwapPointer() {
	// BAD: x doesn't need to be heap allocated
	var x int // ERROR "moved to heap: x"
	var y int // ERROR "moved to heap: y"
	atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(&x), unsafe.Pointer(&y))
}
