//go:build msan

package msan

import (
	"unsafe"
)

const Enabled = true

//go:linkname Read runtime.msanread
func Read(addr unsafe.Pointer, sz uintptr)

//go:linkname Write runtime.msanwrite
func Write(addr unsafe.Pointer, sz uintptr)

//go:linkname Malloc runtime.msanmalloc
func Malloc(addr unsafe.Pointer, sz uintptr)

//go:linkname Free runtime.msanfree
func Free(addr unsafe.Pointer, sz uintptr)

//go:linkname Move runtime.msanmove
func Move(dst, src unsafe.Pointer, sz uintptr)
