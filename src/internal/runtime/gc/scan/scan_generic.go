//go:build !amd64

package scan

import (
	"internal/runtime/gc"
	"unsafe"
)

func HasFastScanSpanPacked() bool {
	// N.B. ScanSpanPackedGeneric isn't actually fast enough to serve as a general-purpose implementation.
	// The runtime's alternative of jumping between each object is still substantially better, even at
	// relatively high object densities.
	return false
}

func ScanSpanPacked(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32) {
	return ScanSpanPackedGo(mem, bufp, objMarks, sizeClass, ptrMask)
}
