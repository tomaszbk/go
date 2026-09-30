// run


package main

import "fmt"

//go:noinline
func f(x uint64) uint64 {
	s := "\x04"
	c := s[0]
	return x << c << 4
}
func main() {
	if want, got := uint64(1<<8), f(1); want != got {
		panic(fmt.Sprintf("want %x got %x", want, got))
	}
}
