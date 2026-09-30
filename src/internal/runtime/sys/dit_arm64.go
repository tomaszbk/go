//go:build arm64

package sys

import (
	"internal/cpu"
	"unsafe"
)

var DITSupported = cpu.ARM64.HasDIT

const offsetARM64HasSB = unsafe.Offsetof(cpu.ARM64.HasSB)

func EnableDIT() bool
func DITEnabled() bool
func DisableDIT()
