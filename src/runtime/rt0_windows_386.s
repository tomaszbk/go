#include "textflag.h"

TEXT _rt0_386_windows(SB),NOSPLIT,$0
	JMP	_rt0_386(SB)

// When building with -buildmode=(c-shared or c-archive), this
// symbol is called. For dynamic libraries it is called when the
// library is loaded. For static libraries it is called when the
// final executable starts, during the C runtime initialization
// phase.
TEXT _rt0_386_windows_lib(SB),NOSPLIT,$0
	JMP	_rt0_386_lib(SB)

TEXT _main(SB),NOSPLIT,$0
	// Remove the return address from the stack.
	// rt0_go doesn't expect it to be there.
	ADDL	$4, SP
	JMP	runtime·rt0_go(SB)
