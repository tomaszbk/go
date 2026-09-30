// run


package main

func shift(x int) int { return 1 << (1 << (1 << (uint(x)))) }

func main() {
	if n := shift(2); n != 1<<(1<<(1<<2)) {
		println("bad shift", n)
		panic("fail")
	}
}
