//go:build gc

#include "textflag.h"

//
// System calls for ppc64, AIX are implemented in runtime/syscall_aix.go
//

TEXT ·syscall6(SB),NOSPLIT,$0-88
	JMP	syscall·syscall6(SB)

TEXT ·rawSyscall6(SB),NOSPLIT,$0-88
	JMP	syscall·rawSyscall6(SB)
