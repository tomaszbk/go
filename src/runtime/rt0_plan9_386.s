#include "textflag.h"

TEXT _rt0_386_plan9(SB),NOSPLIT,$12
	MOVL	AX, _tos(SB)
	LEAL	8(SP), AX
	MOVL	AX, _privates(SB)
	MOVL	$1, _nprivates(SB)
	CALL	runtime·asminit(SB)
	MOVL	inargc-4(FP), AX
	MOVL	AX, 0(SP)
	LEAL	inargv+0(FP), AX
	MOVL	AX, 4(SP)
	JMP	runtime·rt0_go(SB)

GLOBL _tos(SB), NOPTR, $4
GLOBL _privates(SB), NOPTR, $4
GLOBL _nprivates(SB), NOPTR, $4
