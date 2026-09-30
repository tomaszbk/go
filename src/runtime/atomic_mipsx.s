//go:build mips || mipsle

#include "textflag.h"

TEXT ·publicationBarrier(SB),NOSPLIT,$0
	SYNC
	RET
