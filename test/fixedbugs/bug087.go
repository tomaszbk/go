// compile


package main

const s string = "foo";

func main() {
	i := len(s);  // should be legal to take len() of a constant
	_ = i;
}

/*
uetli:~/Source/go1/test/bugs gri$ 6g bug087.go
bug087.go:6: illegal combination of literals LEN 9
bug087.go:6: illegal combination of literals LEN 9
*/
