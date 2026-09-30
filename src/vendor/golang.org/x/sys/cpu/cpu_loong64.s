#include "textflag.h"

// func get_cpucfg(reg uint32) uint32
TEXT ·get_cpucfg(SB), NOSPLIT|NOFRAME, $0
	MOVW	reg+0(FP), R5
	// CPUCFG R5, R4 = 0x00006ca4
	WORD	$0x00006ca4
	MOVW	R4, ret+8(FP)
	RET
