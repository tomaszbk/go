// run


package main

//go:noinline
func f(a, b uint) int {
	return int(a-b) / 8
}

func main() {
	if x := f(1, 2); x != 0 {
		panic(x)
	}
}
