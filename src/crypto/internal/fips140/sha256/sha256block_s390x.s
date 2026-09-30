//go:build !purego

#include "textflag.h"

// func blockS390X(dig *Digest, p []byte)
TEXT ·blockS390X(SB), NOSPLIT|NOFRAME, $0-32
	LMG    dig+0(FP), R1, R3            // R2 = &p[0], R3 = len(p)
	MOVBZ  $2, R0                       // SHA-256 function code

loop:
	KIMD R0, R2      // compute intermediate message digest (KIMD)
	BVS  loop        // continue if interrupted
	RET
