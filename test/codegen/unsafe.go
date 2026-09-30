// asmcheck

package codegen

import "unsafe"

func f(p unsafe.Pointer, x, y uintptr) int64 {
	p = unsafe.Pointer(uintptr(p) + x + y)
	// amd64:`MOVQ \(.*\)\(.*\*1\), `
	// arm64:`MOVD \(R[0-9]+\)\(R[0-9]+\), `
	return *(*int64)(p)
}
