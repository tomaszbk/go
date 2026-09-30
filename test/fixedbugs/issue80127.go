// run

package main

type T [2][2]int64

//go:noinline
func f(i, j int) int64 {
	if i < 0 || i > 1 || j < 0 || j > 1 {
		return 3
	}
	var x T
	x[i][i] = 33

	var y T
	y[i][i] = 44
	r := y[j][j]

	return r + x[i][i]
}

func main() {
	if x := f(1, 1); x != 77 {
		println("bad", x)
	}
}
