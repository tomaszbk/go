TEXT ·trampoline(SB),0,$0
	CALL	libc_getpid(SB)
	CALL	libc_kill(SB)
	CALL	libc_open(SB)
	CALL	libc_close(SB)
	RET
