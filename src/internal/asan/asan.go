//go:build asan

package asan

import (
	"unsafe"
)

const Enabled = true

//go:linkname Read runtime.asanread
func Read(addr unsafe.Pointer, len uintptr)

//go:linkname Write runtime.asanwrite
func Write(addr unsafe.Pointer, len uintptr)
