// run


package main

var G func(int) int

//go:noinline
func callclo(q, r int) int {
	p := func(z int) int {
		G = func(int) int { return 1 }
		return z + 1
	}
	res := p(q) ^ p(r) // These calls to "p" will be inlined
	G = p
	return res
}

func main() {
	callclo(1, 2)
}
