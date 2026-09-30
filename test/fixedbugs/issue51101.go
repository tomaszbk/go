// run

// Issue 51101: on RISCV64, difference of two pointers
// was marked as pointer and crashes GC.

package main

var a, b int

func main() {
	F(&b, &a)
}

//go:noinline
func F(a, b *int) bool {
	x := a == b
	G(x)
	y := a != b
	return y
}

//go:noinline
func G(bool) {
	grow([1000]int{20})
}

func grow(x [1000]int) {
	if x[0] != 0 {
		x[0]--
		grow(x)
	}
}
