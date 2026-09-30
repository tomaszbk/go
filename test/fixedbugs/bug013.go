// run


package main

func main() {
	var cu0 uint16 = '\u1234';
	var cU1 uint32 = '\U00101234';
	_, _ = cu0, cU1;
}
/*
bug13.go:4: missing '
bug13.go:4: syntax error
bug13.go:5: newline in string
bug13.go:5: missing '
bug13.go:6: newline in string
*/
