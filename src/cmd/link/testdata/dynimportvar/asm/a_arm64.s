//go:build darwin

TEXT ·Mach_task_self(SB),0,$0-4
	MOVD	$libc_mach_task_self_(SB), R0
	MOVD	(R0), R0
	MOVW	R0, ret+0(FP)
	RET
