#include "textflag.h"

TEXT _rt0_amd64_illumos(SB),NOSPLIT,$-8
	JMP	_rt0_amd64(SB)

TEXT _rt0_amd64_illumos_lib(SB),NOSPLIT,$0
	JMP	_rt0_amd64_lib(SB)
