// run


package main

//go:noinline
func f(c bool) int {
	b := true
	x := 0
	y := 1
	for b {
		b = false
		y = x
		x = 2
		if c {
			return 3
		}
	}
	return y
}

func main() {
	if got := f(false); got != 0 {
		panic(got)
	}
}
