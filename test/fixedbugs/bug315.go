// compile


// Issue 1368.

package main

func main() {
	a := complex(2, 2)
	a /= 2
}

/*
bug315.go:13: internal compiler error: optoas: no entry DIV-complex
*/
