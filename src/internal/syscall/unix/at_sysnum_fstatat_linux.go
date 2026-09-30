//go:build arm64 || riscv64

package unix

import "syscall"

const fstatatTrap uintptr = syscall.SYS_FSTATAT
