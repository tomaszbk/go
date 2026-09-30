//go:build !purego

#include "textflag.h"

DATA StaticData<>(SB)/4, $10
GLOBL StaticData<>(SB), NOPTR, $4

TEXT StaticText<>(SB), $0
	RET

TEXT ·PtrStaticData(SB), $0-8
	MOVD $StaticData<>(SB), R1
	MOVD R1, ret+0(FP)
	RET

TEXT ·PtrStaticText(SB), $0-8
	MOVD $StaticText<>(SB), R1
	MOVD R1, ret+0(FP)
	RET
