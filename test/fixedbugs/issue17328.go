// errorcheck


package main

func main() {
	i := 0
	for ; ; i++) { // ERROR "unexpected \), expecting { after for clause|expected .*{.*|expected .*;.*"
	}
} // GCCGO_ERROR "expected declaration"
