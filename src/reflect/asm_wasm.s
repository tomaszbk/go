#include "textflag.h"
#include "funcdata.h"

// makeFuncStub is the code half of the function returned by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go
// for more details.
// No arg size here; runtime pulls arg map out of the func value.
TEXT ·makeFuncStub(SB),(NOSPLIT|WRAPPER),$40
	NO_LOCAL_POINTERS

	MOVD CTXT, 0(SP)

	Get SP
	Get SP
	I64ExtendI32U
	I64Const $argframe+0(FP)
	I64Add
	I64Store $8

	MOVB $0, 32(SP)
	MOVD $32(SP), 16(SP)
	MOVD $0, 24(SP)

	CALL ·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by makeMethodValue.
// See the comment on the declaration of methodValueCall in makefunc.go
// for more details.
// No arg size here; runtime pulls arg map out of the func value.
TEXT ·methodValueCall(SB),(NOSPLIT|WRAPPER),$40
	NO_LOCAL_POINTERS

	MOVD CTXT, 0(SP)

	Get SP
	Get SP
	I64ExtendI32U
	I64Const $argframe+0(FP)
	I64Add
	I64Store $8

	MOVB $0, 32(SP)
	MOVD $32(SP), 16(SP)
	MOVD $0, 24(SP)

	CALL ·callMethod(SB)
	RET
