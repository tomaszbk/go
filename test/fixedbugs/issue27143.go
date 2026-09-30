// compile

// Issue 27143: cmd/compile: erroneous application of walkinrange
// optimization for const over 2**63

package p

var c uint64

var b1 bool = 0x7fffffffffffffff < c && c < 0x8000000000000000
var b2 bool = c < 0x8000000000000000 && 0x7fffffffffffffff < c
var b3 bool = 0x8000000000000000 < c && c < 0x8000000000000001
var b4 bool = c < 0x8000000000000001 && 0x8000000000000000 < c
