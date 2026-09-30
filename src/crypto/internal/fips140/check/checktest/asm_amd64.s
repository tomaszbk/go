//go:build !purego

#include "textflag.h"

DATA StaticData<>(SB)/4, $10
GLOBL StaticData<>(SB), NOPTR, $4

TEXT StaticText<>(SB), $0
	RET

TEXT ·PtrStaticData(SB), $0-8
	MOVQ $StaticData<>(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·PtrStaticText(SB), $0-8
	MOVQ $StaticText<>(SB), AX
	MOVQ AX, ret+0(FP)
	RET
