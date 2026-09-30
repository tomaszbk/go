#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

TEXT _rt0_amd64_windows(SB),NOSPLIT,$0
	JMP	_rt0_amd64(SB)

// When building with -buildmode=(c-shared or c-archive), this
// symbol is called.
TEXT _rt0_amd64_windows_lib(SB),NOSPLIT,$0
	JMP	_rt0_amd64_lib(SB)
