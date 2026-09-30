#include "textflag.h"

TEXT _rt0_arm64_linux(SB),NOSPLIT,$0
	JMP	_rt0_arm64(SB)

// When building with -buildmode=c-shared, this symbol is called when the shared
// library is loaded.
TEXT _rt0_arm64_linux_lib(SB),NOSPLIT,$0
	JMP	_rt0_arm64_lib(SB)
