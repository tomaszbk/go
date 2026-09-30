// errorcheck


package main

func main() {
	if x { }	// GCCGO_ERROR "undefined"
	else { }	// ERROR "unexpected semicolon or newline before .?else.?|unexpected keyword else"
}


