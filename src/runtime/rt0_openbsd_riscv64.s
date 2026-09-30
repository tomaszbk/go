#include "textflag.h"

TEXT _rt0_riscv64_openbsd(SB),NOSPLIT|NOFRAME,$0
	MOV	0(X2), A0	// argc
	ADD	$8, X2, A1	// argv
	JMP	main(SB)

TEXT main(SB),NOSPLIT|NOFRAME,$0
	MOV	$runtime·rt0_go(SB), T0
	JALR	ZERO, T0
