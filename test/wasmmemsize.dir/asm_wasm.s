#include "textflag.h"

TEXT ·currentMemory(SB), NOSPLIT, $0
	Get SP
	CurrentMemory
	I32Store ret+0(FP)
	RET
