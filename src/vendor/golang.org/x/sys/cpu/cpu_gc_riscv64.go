//go:build gc

package cpu

// Can only be called when the vector extension is present.
// Implemented in cpu_riscv64.s.
func readVLENB() uint
