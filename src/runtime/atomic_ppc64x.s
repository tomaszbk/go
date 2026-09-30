//go:build ppc64 || ppc64le

#include "textflag.h"

TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	// LWSYNC is the "export" barrier recommended by Power ISA
	// v2.07 book II, appendix B.2.2.2.
	// LWSYNC is a load/load, load/store, and store/store barrier.
	LWSYNC
	RET
