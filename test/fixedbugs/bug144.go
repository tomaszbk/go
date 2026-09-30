// compile


package main

const c = 1;

func main() {
	c := 0;
	_ = c;
}

/*
bug144.go:8: left side of := must be a name
bug144.go:8: operation LITERAL not allowed in assignment context
bug144.go:8: illegal types for operand: AS
	ideal
	int
*/
