// compile


// Indexed export format must not crash when writing
// the anonymous parameter for m.

package p

var x interface {
	m(int)
}

var M = x.m
