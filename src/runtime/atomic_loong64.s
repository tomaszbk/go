#include "textflag.h"

TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	DBAR	$0x1A // StoreStore barrier
	RET
