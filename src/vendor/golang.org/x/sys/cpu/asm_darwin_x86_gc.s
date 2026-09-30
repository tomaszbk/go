//go:build darwin && amd64 && gc

#include "textflag.h"

TEXT libc_sysctl_trampoline<>(SB),NOSPLIT,$0-0
	JMP	libc_sysctl(SB)
GLOBL	·libc_sysctl_trampoline_addr(SB), RODATA, $8
DATA	·libc_sysctl_trampoline_addr(SB)/8, $libc_sysctl_trampoline<>(SB)

TEXT libc_sysctlbyname_trampoline<>(SB),NOSPLIT,$0-0
	JMP	libc_sysctlbyname(SB)
GLOBL	·libc_sysctlbyname_trampoline_addr(SB), RODATA, $8
DATA	·libc_sysctlbyname_trampoline_addr(SB)/8, $libc_sysctlbyname_trampoline<>(SB)
