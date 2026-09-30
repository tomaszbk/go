TEXT foo(SB), 0, $0
	RET
TEXT foo(SB), 0, $0 // ERROR "symbol foo redeclared"
	RET

GLOBL bar(SB), 0, $8
GLOBL bar(SB), 0, $8 // ERROR "symbol bar redeclared"

DATA bar+0(SB)/8, $0
DATA bar+0(SB)/8, $0 // ERROR "overlapping DATA entry for bar"
