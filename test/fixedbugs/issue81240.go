// compile


package p

const W = 32 << (^uintptr(0) >> 63) // 32 or 64

type T struct {
	a [1<<(W-30) - 1]byte
}

func f(x, y *T) {
	*x = *y
}
