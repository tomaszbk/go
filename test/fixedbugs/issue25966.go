// compile -N


// Issue 25966: liveness code complains autotmp live on
// function entry.

package p

var F = []func(){
	func() func() { return (func())(nil) }(),
}

var A = []int{}

type ss struct {
	string
	float64
	i int
}

var V = A[ss{}.i]
