package ssa

import "cmd/compile/internal/ssa/ssaop"

func (v *Value) ClobbersFlags() bool {
	if ssaop.OpcodeTable[v.Op].ClobberFlags {
		return true
	}
	if v.Type.IsTuple() && (v.Type.FieldType(0).IsFlags() || v.Type.FieldType(1).IsFlags()) {
		// This case handles the possibility where a flag value is generated but never used.
		// In that case, there's no corresponding Select to overwrite the flags value,
		// so we must consider flags clobbered by the tuple-generating instruction.
		return true
	}
	return false
}
