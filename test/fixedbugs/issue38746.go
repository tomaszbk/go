// compile


package main

var g *uint64

func main() {
	var v uint64
	g = &v
	v &^= (1 << 31)
	v |= 1 << 63
	v &^= (1 << 63)
}
