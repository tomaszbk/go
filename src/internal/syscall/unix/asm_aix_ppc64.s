#include "textflag.h"

//
// System calls for aix/ppc64 are implemented in syscall/syscall_aix.go
//

TEXT ·syscall6(SB),NOSPLIT,$0
	JMP	syscall·syscall6(SB)
