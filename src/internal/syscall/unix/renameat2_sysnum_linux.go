//go:build linux && (loong64 || riscv64)

package unix

import "syscall"

const (
	// loong64 and riscv64 only have renameat2.
	// renameat2 has an extra flags parameter.
	// When called with a 0 flags it is identical to renameat.
	renameatTrap uintptr = syscall.SYS_RENAMEAT2
)
