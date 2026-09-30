#include "textflag.h"

TEXT ·sysvicall6(SB),NOSPLIT,$0-88
	JMP	syscall·sysvicall6(SB)
