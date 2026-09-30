//go:build linux && !(loong64 || riscv64)

package unix

import "syscall"

const (
	renameatTrap uintptr = syscall.SYS_RENAMEAT
)
