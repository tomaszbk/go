// errorcheck


package main

func main() {
	if x; y		// ERROR "expected .*{.* after if clause|undefined"
	{
		z	// GCCGO_ERROR "undefined"


