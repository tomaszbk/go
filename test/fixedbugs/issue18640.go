// compile


package p

type (
	a = b
	b struct {
		*a
	}
)

type (
	c struct {
		*d
	}
	d = c
)

// The compiler cannot handle these cases. Disabled for now.
// See issue #25838.
/*
type (
	e = f
	f = g
	g = []h
	h i
	i = j
	j = e
)

type (
	a1 struct{ *b1 }
	b1 = c1
	c1 struct{ *b1 }
)

type (
	a2 struct{ b2 }
	b2 = c2
	c2 struct{ *b2 }
)
*/
