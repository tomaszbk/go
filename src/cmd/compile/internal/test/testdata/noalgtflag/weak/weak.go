package weak

import "internal/abi"

type M map[uint64]uint64

// Emit a noalg descriptor for [8]uint64.
var Sink any = M{}

//go:noinline
func FoldedFlags() abi.TFlag {
	return abi.TypeFor[[8]uint64]().TFlag
}
