//go:build arm || mips || mipsle || 386

package unix

import "syscall"

const fstatatTrap uintptr = syscall.SYS_FSTATAT64
