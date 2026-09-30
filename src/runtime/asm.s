#include "textflag.h"

#ifndef GOARCH_amd64
TEXT ·sigpanic0(SB),NOSPLIT,$0-0
	JMP	·sigpanic<ABIInternal>(SB)
#endif

// See map.go comment on the need for this routine.
TEXT ·mapinitnoop<ABIInternal>(SB),NOSPLIT,$0-0
	RET

