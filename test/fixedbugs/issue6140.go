// compile


// Issue 6140: compiler incorrectly rejects method values
// whose receiver has an unnamed interface type.

package p

type T *interface {
	m() int
}

var x T

var _ = (*x).m

var y interface {
	m() int
}

var _ = y.m

type I interface {
	String() string
}

var z *struct{ I }
var _ = z.String
