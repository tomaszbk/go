//go:build !purego

#include "textflag.h"

DATA StaticData<>(SB)/4, $10
GLOBL StaticData<>(SB), NOPTR, $4

TEXT StaticText<>(SB), $0
	RET

TEXT ·PtrStaticData(SB), $0-8
	MOVV $StaticData<>(SB), R4
	MOVV R4, ret+0(FP)
	RET

TEXT ·PtrStaticText(SB), $0-8
	MOVV $StaticText<>(SB), R4
	MOVV R4, ret+0(FP)
	RET
