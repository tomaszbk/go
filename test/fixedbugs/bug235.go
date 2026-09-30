// compile


// used to crash the compiler

package bug235

type T struct {
	x [4]byte
}

var p *T
var v = *p

