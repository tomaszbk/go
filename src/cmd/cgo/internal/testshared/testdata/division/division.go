package main

//go:noinline
func div(x, y uint32) uint32 {
	return x / y
}

func main() {
	a := div(97, 11)
	if a != 8 {
		panic("FAIL")
	}
}
