#include "textflag.h"

// See memclrNoHeapPointers Go doc for important implementation constraints.

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
TEXT runtime·memclrNoHeapPointers(SB), NOSPLIT, $0-16
	MOVD ptr+0(FP), R0
	MOVD n+8(FP), R1

	Get R0
	I32WrapI64
	I32Const $0
	Get R1
	I32WrapI64
	MemoryFill
	RET
