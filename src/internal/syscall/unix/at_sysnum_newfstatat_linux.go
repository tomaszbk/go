//go:build amd64 || mips64 || mips64le || ppc64 || ppc64le || s390x

package unix

import "syscall"

const fstatatTrap uintptr = syscall.SYS_NEWFSTATAT
