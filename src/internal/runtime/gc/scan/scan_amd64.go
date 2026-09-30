package scan

import (
	"internal/cpu"
	"internal/runtime/gc"
	"unsafe"
)

func ScanSpanPacked(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32) {
	if CanAVX512() {
		return ScanSpanPackedAVX512(mem, bufp, objMarks, sizeClass, ptrMask)
	}
	panic("not implemented")
}

func HasFastScanSpanPacked() bool {
	return avx512ScanPackedReqsMet
}

// -- AVX512 --

func CanAVX512() bool {
	return avx512ScanPackedReqsMet
}

func ScanSpanPackedAVX512(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32) {
	return FilterNilAVX512(bufp, scanSpanPackedAVX512(mem, bufp, objMarks, sizeClass, ptrMask))
}

//go:noescape
func scanSpanPackedAVX512(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32)

var avx512ScanPackedReqsMet = cpu.X86.HasAVX512VL &&
	cpu.X86.HasAVX512BW &&
	cpu.X86.HasGFNI &&
	cpu.X86.HasAVX512BITALG &&
	cpu.X86.HasAVX512DQ && // for kmovb, see #79871
	cpu.X86.HasAVX512VBMI &&
	cpu.X86.HasPOPCNT
