// compile


package main

func main() {
	x := false;
	func () { if x          { println(1); } }();  // this does not compile
	func () { if x == false { println(2); } }();  // this works as expected
}

/*
bug139.go:7: fatal error: naddr: ONAME class x 5
*/
