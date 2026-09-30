// run


package main

func main() {
	Foo: {
		return;
	}
	goto Foo;
}
/*
bug5.go:4: Foo undefined
bug5.go:4: fatal error: walktype: switch 1 unknown op GOTO l(4)
*/
