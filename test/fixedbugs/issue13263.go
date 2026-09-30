// compile


package b

var (
	x uint
	y = x
	z = uintptr(y)
	a = uint32(y)
	b = uint64(y)
)
