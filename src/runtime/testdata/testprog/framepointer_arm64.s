#include "textflag.h"

TEXT	·getFP(SB), NOSPLIT|NOFRAME, $0-8
	MOVD	R29, ret+0(FP)
	RET
