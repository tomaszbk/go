// func ReadX15() uint64
TEXT ·ReadX15(SB), $0-8
	MOVQ	X15, ret+0(FP)
	RET
