package types

const BADWIDTH = -1000000000

type bitset8 uint8

func (f *bitset8) set(mask uint8, b bool) {
	if b {
		*(*uint8)(f) |= mask
	} else {
		*(*uint8)(f) &^= mask
	}
}

type bitset16 uint16

func (f *bitset16) set(mask uint16, b bool) {
	if b {
		*(*uint16)(f) |= mask
	} else {
		*(*uint16)(f) &^= mask
	}
}
