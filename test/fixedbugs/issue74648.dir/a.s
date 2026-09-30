TEXT ·F(SB), $0
	JMP prealigned
	INT $3 // should never be reached
prealigned:
	PCALIGN $0x10
aligned:
	RET
