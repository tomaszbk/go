// run

// When the function Store an Arg and also use it in another place,
// be sure not to generate duplicated OpArgXXXReg values, which confuses
// the register allocator.

package main

//go:noinline
//go:registerparams
func F(x, y float32) {
	if x < 0 {
		panic("FAIL")
	}
	g = [4]float32{x, y, x, y}
}

var g [4]float32

func main() {
	F(1, 2)
	if g[0] != 1 || g[1] != 2 || g[2] != 1 || g[3] != 2 {
		panic("FAIL")
	}
}
