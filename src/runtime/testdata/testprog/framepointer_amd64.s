#include "textflag.h"

TEXT	·getFP(SB), NOSPLIT|NOFRAME, $0-8
	MOVQ	BP, ret+0(FP)
	RET
