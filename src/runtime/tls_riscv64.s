#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// If !iscgo, this is a no-op.
//
// NOTE: mcall() assumes this clobbers only X31 (REG_TMP).
TEXT runtime·save_g(SB),NOSPLIT|NOFRAME,$0-0
#ifndef GOOS_openbsd
	MOVB	runtime·iscgo(SB), X31
	BEQZ	X31, nocgo
#endif
	MOV	g, runtime·tls_g(SB)
nocgo:
	RET

TEXT runtime·load_g(SB),NOSPLIT|NOFRAME,$0-0
	MOV	runtime·tls_g(SB), g
	RET

GLOBL runtime·tls_g(SB), TLSBSS, $8
