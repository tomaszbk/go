// errorcheck


package main

func main() {
	for x		// GCCGO_ERROR "undefined"
	{		// ERROR "unexpected {, expected for loop condition|expecting .*{.* after for clause"
		z	// GCCGO_ERROR "undefined"
