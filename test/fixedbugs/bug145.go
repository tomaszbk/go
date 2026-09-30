// compile


package main

type t int

func main() {
	t := 0;
	_ = t;
}

/*
bug145.go:8: t is type, not var
*/
