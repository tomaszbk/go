#include "textflag.h"

TEXT _rt0_amd64_plan9(SB),NOSPLIT,$24
	MOVQ	AX, _tos(SB)
	LEAQ	16(SP), AX
	MOVQ	AX, _privates(SB)
	MOVL	$1, _nprivates(SB)
	MOVL	inargc-8(FP), DI
	LEAQ	inargv+0(FP), SI
	MOVQ	$runtime·rt0_go(SB), AX
	JMP	AX

GLOBL _tos(SB), NOPTR, $8
GLOBL _privates(SB), NOPTR, $8
GLOBL _nprivates(SB), NOPTR, $4
