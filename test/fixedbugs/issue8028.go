// compile


// Issue 8028. Used to fail in -race mode with "non-orig name" error.

package p

var (
	t2 = T{F, "s1"}
	t1 = T{F, "s2"}

	tt = [...]T{t1, t2}
)

type I interface{}

type T struct {
	F func() I
	S string
}

type E struct{}

func F() I { return new(E) }
