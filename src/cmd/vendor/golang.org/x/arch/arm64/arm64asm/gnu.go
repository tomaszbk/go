package arm64asm

import (
	"strings"
)

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This form typically matches the syntax defined in the ARM Reference Manual.
func GNUSyntax(inst Inst) string {
	switch inst.Op {
	case RET:
		if r, ok := inst.Args[0].(Reg); ok && r == X30 {
			return "ret"
		}
	case B:
		if _, ok := inst.Args[0].(Cond); ok {
			return strings.ToLower("b." + inst.Args[0].String() + " " + inst.Args[1].String())
		}
	case SYSL:
		result := strings.ToLower(inst.String())
		return strings.Replace(result, "c", "C", -1)
	case DCPS1, DCPS2, DCPS3, CLREX:
		return strings.ToLower(strings.TrimSpace(inst.String()))
	case ISB:
		if strings.Contains(inst.String(), "SY") {
			result := strings.TrimSuffix(inst.String(), " SY")
			return strings.ToLower(result)
		}
	}
	return strings.ToLower(inst.String())
}
