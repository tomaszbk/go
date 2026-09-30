// compile


// Gccgo crashed compiling this code due to failing to finalize
// interfaces in the right order.

package p

type s1 struct {
	f m
	I
}

type m interface {
	Mm(*s2)
}

type s2 struct {
	*s1
}

type I interface {
	MI()
}
