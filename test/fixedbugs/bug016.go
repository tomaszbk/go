// errorcheck


package main

func main() {
	var i int = 100
	i = i << -3 // ERROR "overflows|negative"
}

/*
ixedbugs/bug016.go:7: overflow converting constant to <uint32>UINT32
fixedbugs/bug016.go:7: illegal types for operand: AS
	(<int32>INT32)
*/
