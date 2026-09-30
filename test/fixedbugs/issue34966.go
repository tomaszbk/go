// compile -d=checkptr

package p

import "unsafe"

type ptr unsafe.Pointer

func f(p ptr) *int { return (*int)(p) }
func g(p ptr) ptr  { return ptr(uintptr(p) + 1) }
