//go:build gc

#include "textflag.h"

//
// System calls for amd64, Solaris are implemented in runtime/syscall_solaris.go
//

TEXT ·sysvicall6(SB),NOSPLIT,$0-88
	JMP	syscall·sysvicall6(SB)

TEXT ·rawSysvicall6(SB),NOSPLIT,$0-88
	JMP	syscall·rawSysvicall6(SB)
