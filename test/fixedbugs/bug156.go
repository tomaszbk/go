// compile


package main

func f(a int64) int64 {
	const b int64 = 0;
	n := a &^ b;
	return n;
}

func main() {
	f(1)
}

/*
bug156.go:7: constant 18446744073709551615 overflows int64
*/
