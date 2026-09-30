// run


package main

func f(x int) int {
        if x == 1 {
                return 7
        }
        return 7 / (x - 1)
}

//go:noinline
func g(x int) int {
	r := 0
	for range 5 {
		r += f(x)
	}
	return r
}

func main() {
	g(1)
}
