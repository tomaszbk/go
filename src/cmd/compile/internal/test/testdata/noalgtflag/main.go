package main

import (
	"cmd/compile/internal/test/testdata/noalgtflag/weak"
	"internal/abi"
)

var strong any = [8]uint64{}

func main() {
	if weak.FoldedFlags() != abi.TypeOf(strong).TFlag {
		panic("folded flags disagree with linked descriptor")
	}
}
