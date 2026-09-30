// errorcheck


package main

func main() {
	for x; y; z	// ERROR "expected .*{.* after for clause|undefined"
	{
		z	// GCCGO_ERROR "undefined"


