// compile

package p

type S struct {
	a [4]struct{}
	f chan int
}

func f(p *S) {
	var s S

	// Memory write that requires a write barrier should work
	// with structs having zero-sized arrays of non-zero elements.
	*p = s
}
