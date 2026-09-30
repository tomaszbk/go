//go:build gc && !arm

#include "textflag.h"

TEXT ·useVFPv1(SB),NOSPLIT,$0
	RET

TEXT ·useVFPv3(SB),NOSPLIT,$0
	RET

TEXT ·useARMv6K(SB),NOSPLIT,$0
	RET
