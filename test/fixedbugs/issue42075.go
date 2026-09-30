// errorcheck

package p

import "unsafe"

type T struct { // ERROR "recursive type"
	x int
	p unsafe.Pointer

	f T
}
