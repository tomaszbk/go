// run


package main

func frexp() (a int, b float64) {
	return 1, 2.0
}

func main() {
	a, b := frexp();
	_, _ = a, b;
}

/*
bug056.go:8: illegal types for operand: AS
	(<int32>INT32)
	(<int32>INT32)
*/
