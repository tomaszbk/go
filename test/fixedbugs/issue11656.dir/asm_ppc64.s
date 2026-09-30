#include "textflag.h"

// func syncIcache(p uintptr)
TEXT main·syncIcache(SB), NOSPLIT|NOFRAME, $0-0
	SYNC
	MOVD (R3), R3
	ICBI (R3)
	ISYNC
	RET
