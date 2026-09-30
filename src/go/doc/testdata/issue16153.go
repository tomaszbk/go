package issue16153

// original test case
const (
	x1 uint8 = 255
	Y1       = 256
)

// variations
const (
	x2 uint8 = 255
	Y2
)

const (
	X3 int64 = iota
	Y3       = 1
)

const (
	X4 int64 = iota
	Y4
)
