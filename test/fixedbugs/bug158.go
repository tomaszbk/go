// compile


package main

func main() {
	x := 0;

	// this compiles
	switch x {
	case 0:
	}

	// this doesn't but should
	switch 0 {
	case 0:
	}
}


/*
bug158.go:14: fatal error: dowidth: unknown type: E-33
*/
