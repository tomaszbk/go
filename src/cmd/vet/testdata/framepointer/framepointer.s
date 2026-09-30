//go:build amd64
// +build amd64

#include "textflag.h"

TEXT ·f(SB),0,$0-0
	MOVQ $0, BP // ERROR "frame pointer is clobbered before saving"
	RET
