//go:build !purego

#include "textflag.h"

DATA StaticData<>(SB)/4, $10
GLOBL StaticData<>(SB), NOPTR, $4

TEXT StaticText<>(SB), $0
	RET

TEXT ·PtrStaticData(SB), $0-4
	MOVL $StaticData<>(SB), AX
	MOVL AX, ret+0(FP)
	RET

TEXT ·PtrStaticText(SB), $0-4
	MOVL $StaticText<>(SB), AX
	MOVL AX, ret+0(FP)
	RET
