//go:build darwin

TEXT ·Mach_task_self(SB),0,$0-4
	MOVQ	$libc_mach_task_self_(SB), AX
	MOVQ	(AX), AX
	MOVL	AX, ret+0(FP)
	RET
