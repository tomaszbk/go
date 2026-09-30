// run


package main

type T struct {
	i int
}

func main() {
	var ta []*T;

	ta = new([1]*T)[0:];
	ta[0] = nil;
}
/*
bug045.go:13: fatal error: goc: exit 1
*/
