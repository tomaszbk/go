//go:build mips64 || mips64le

#include "textflag.h"

#define SYNC	WORD $0xf

TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	SYNC
	RET
