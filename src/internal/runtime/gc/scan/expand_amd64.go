package scan

import "internal/runtime/gc"

// ExpandAVX512 expands each bit in packed into f consecutive bits in unpacked,
// where f is the word size of objects in sizeClass.
//
// This is a testing entrypoint to the expanders used by scanSpanPacked*.
//
//go:noescape
func ExpandAVX512(sizeClass int, packed *gc.ObjMask, unpacked *gc.PtrMask)

// gcExpandersAVX512 is the PCs of expander functions. These cannot be called directly
// as they don't follow the Go ABI, but you can use this to check if a given
// expander PC is 0.
//
// It is defined in assembly.
var gcExpandersAVX512 [len(gc.SizeClassToSize)]uintptr
