//go:build goexperiment.runtimesecret

package secret

import (
	"internal/cpu"
	"unsafe"
)

// exports for assembly testing functions
const (
	offsetX86HasAVX      = unsafe.Offsetof(cpu.X86.HasAVX)
	offsetX86HasAVX512   = unsafe.Offsetof(cpu.X86.HasAVX512)
	offsetLOONG64HasLSX  = unsafe.Offsetof(cpu.Loong64.HasLSX)
	offsetLOONG64HasLASX = unsafe.Offsetof(cpu.Loong64.HasLASX)
)
