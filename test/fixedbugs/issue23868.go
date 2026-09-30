// compile

// Crashed gccgo.

package p

var F func([0]int) int
var G func() [0]int

var V = make([]int, F(G()))
