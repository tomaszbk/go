#include "textflag.h"

TEXT ·StorepNoWB(SB), NOSPLIT, $0-16
	MOVD ptr+0(FP), R0
	MOVD val+8(FP), 0(R0)
	RET
