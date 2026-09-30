#include "textflag.h"

TEXT ·noframe(SB), NOSPLIT|NOFRAME, $0-0
	XORL	BP, BP
	RET
