#include "textflag.h"

//in plan 9 argc is at top of stack followed by ptrs to arguments

TEXT _rt0_arm_plan9(SB),NOSPLIT|NOFRAME,$0
	MOVW	R0, _tos(SB)
	MOVW	0(R13), R0
	MOVW	$4(R13), R1
	B	runtime·rt0_go(SB)

GLOBL _tos(SB), NOPTR, $4
