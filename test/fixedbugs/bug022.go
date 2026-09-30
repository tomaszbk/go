// errorcheck


package main

func putint(digits *string) {
	var i byte;
	i = (*digits)[7];  // compiles
	i = digits[7];  // ERROR "illegal|is not|cannot index"
	_ = i;
}

func main() {
	s := "asdfasdfasdfasdf";
	putint(&s);
}

/*
bug022.go:8: illegal types for operand
	(*<string>*STRING) INDEXPTR (<int32>INT32)
bug022.go:8: illegal types for operand
	(<uint8>UINT8) AS
*/
