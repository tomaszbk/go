#include "textflag.h"
#include "funcdata.h"

// makeFuncStub is the code half of the function returned by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go
// for more details.
// No argsize here, gc generates argsize info at call site.
TEXT ·makeFuncStub(SB),(NOSPLIT|WRAPPER),$20
	NO_LOCAL_POINTERS
	MOVL	DX, 0(SP)
	LEAL	argframe+0(FP), CX
	MOVL	CX, 4(SP)
	MOVB	$0, 16(SP)
	LEAL	16(SP), AX
	MOVL	AX, 8(SP)
	MOVL	$0, 12(SP)
	CALL	·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by makeMethodValue.
// See the comment on the declaration of methodValueCall in makefunc.go
// for more details.
// No argsize here, gc generates argsize info at call site.
TEXT ·methodValueCall(SB),(NOSPLIT|WRAPPER),$20
	NO_LOCAL_POINTERS
	MOVL	DX, 0(SP)
	LEAL	argframe+0(FP), CX
	MOVL	CX, 4(SP)
	MOVB	$0, 16(SP)
	LEAL	16(SP), AX
	MOVL	AX, 8(SP)
	MOVL	$0, 12(SP)
	CALL	·callMethod(SB)
	RET
