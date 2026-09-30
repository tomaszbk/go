#include "textflag.h"

TEXT _rt0_arm64_darwin(SB),NOSPLIT,$0
	// Darwin puts argc and argv in R0 and R1,
	// so there is no need to go through _rt0_arm64.
	JMP	runtime·rt0_go(SB)

// When linking with -buildmode=c-archive or -buildmode=c-shared,
// this symbol is called from a global initialization function.
TEXT _rt0_arm64_darwin_lib(SB),NOSPLIT,$0
	JMP	_rt0_arm64_lib(SB)
