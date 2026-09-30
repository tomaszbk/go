// errorcheck


package main

func atom(s string) {
	if s == nil {	// ERROR "nil|incompatible"
		return;
	}
}

func main() {}

/*
bug047.go:4: fatal error: stringpool: not string
*/
