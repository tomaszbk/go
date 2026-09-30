#include "textflag.h"

TEXT _rt0_arm_openbsd(SB),NOSPLIT,$0
	B	_rt0_arm(SB)

TEXT _rt0_arm_openbsd_lib(SB),NOSPLIT,$0
	B	_rt0_arm_lib(SB)
