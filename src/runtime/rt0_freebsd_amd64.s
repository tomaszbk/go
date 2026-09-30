#include "textflag.h"

// On FreeBSD argc/argv are passed in DI, not SP, so we can't use _rt0_amd64.
TEXT _rt0_amd64_freebsd(SB),NOSPLIT,$-8
	LEAQ	8(DI), SI // argv
	MOVQ	0(DI), DI // argc
	JMP	runtime·rt0_go(SB)

TEXT _rt0_amd64_freebsd_lib(SB),NOSPLIT,$0
	JMP	_rt0_amd64_lib(SB)
