//go:build linux && !(386 || arm || mips || mipsle || (gccgo && (ppc || s390)))

package runtime

import (
	"unsafe"
)

func configure64bitsTimeOn32BitsArchitectures() {}

//go:noescape
func futex(addr unsafe.Pointer, op int32, val uint32, ts *timespec, addr2 unsafe.Pointer, val3 uint32) int32

//go:noescape
func timer_settime(timerid int32, flags int32, new, old *itimerspec) int32
