#include "textflag.h"

TEXT _rt0_386_linux(SB),NOSPLIT,$0
	JMP	_rt0_386(SB)

TEXT _rt0_386_linux_lib(SB),NOSPLIT,$0
	JMP	_rt0_386_lib(SB)

TEXT main(SB),NOSPLIT,$0
	// Remove the return address from the stack.
	// rt0_go doesn't expect it to be there.
	ADDL	$4, SP
	JMP	runtime·rt0_go(SB)
