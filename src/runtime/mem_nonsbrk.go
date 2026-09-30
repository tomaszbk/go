//go:build !plan9 && !wasm

package runtime

import "unsafe"

const isSbrkPlatform = false

func sysReserveAlignedSbrk(size, align uintptr) (unsafe.Pointer, uintptr) {
	panic("unreachable")
}
