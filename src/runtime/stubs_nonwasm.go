//go:build !wasm

package runtime

// pause is only used on wasm.
func pause(newsp uintptr) { panic("unreachable") }
