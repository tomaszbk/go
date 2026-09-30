#include "textflag.h"

TEXT	·checkAVX(SB), NOSPLIT|NOFRAME, $0-0
	VXORPS	X1, X2, X3
	RET
