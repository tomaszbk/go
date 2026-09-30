// run


package main

func main() {
	var i int;
	i = '\'';
	i = '\\';
	var s string;
	s = "\"";
	_, _ = i, s;
}
/*
bug.go:5: unknown escape sequence: '
bug.go:6: unknown escape sequence: \
bug.go:8: unknown escape sequence: "
*/
