// asmcheck


package codegen

//go:noinline
func use16(b []byte, x uint64) {
	b[0] = byte(x)
}

//go:noinline
func resultSlotArray(x uint64) [16]byte {
	var ret [16]byte
	// amd64:-`.*\.ret\+`
	use16(ret[:8], x)
	// amd64:-`.*\.ret\+`
	use16(ret[8:], x)
	// amd64:-`.*\.ret\+` -`.*MOVUPS.*~r0`
	return ret
}

//go:noinline
func resultSlotArrayMultipleReturns(x uint64, c bool) [16]byte {
	var ret [16]byte
	// amd64:-`.*\.ret\+`
	use16(ret[:8], x)
	if c {
		// amd64:-`.*\.ret\+` -`.*MOVUPS.*~r0`
		return ret
	}
	// amd64:-`.*\.ret\+`
	use16(ret[8:], x)
	// amd64:-`.*\.ret\+` -`.*MOVUPS.*~r0`
	return ret
}

//go:noinline
func resultSlotPartialReturn(x uint64, c bool) [16]byte {
	var ret [16]byte
	// amd64:-`.*\.ret\+`
	use16(ret[:8], x)
	if c {
		// amd64:-`.*\.ret\+` -`.*MOVUPS.*~r0`
		return ret
	}
	return [16]byte{1}
}

//go:noinline
func resultSlotPartialReturnLocal(x uint64, c bool) [16]byte {
	var ret [16]byte
	// amd64:-`.*\.ret\+`
	use16(ret[:8], x)
	if c {
		// amd64:-`.*\.ret\+` -`.*MOVUPS.*~r0`
		return ret
	}
	other := [16]byte{1}
	return other
}
