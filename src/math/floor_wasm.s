#include "textflag.h"

TEXT ·archFloor(SB),NOSPLIT,$0
	Get SP
	F64Load x+0(FP)
	F64Floor
	F64Store ret+8(FP)
	RET

TEXT ·archCeil(SB),NOSPLIT,$0
	Get SP
	F64Load x+0(FP)
	F64Ceil
	F64Store ret+8(FP)
	RET

TEXT ·archTrunc(SB),NOSPLIT,$0
	Get SP
	F64Load x+0(FP)
	F64Trunc
	F64Store ret+8(FP)
	RET
