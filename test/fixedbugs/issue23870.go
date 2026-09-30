// compile


// Crashed gccgo.

package p

var F func() [0]struct{
	A int
}

var i int
var V = (F()[i]).A
