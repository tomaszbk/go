#include "textflag.h"

TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	DMB	$0xe	// DMB ST
	RET
