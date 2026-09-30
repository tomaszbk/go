// run


package main

func main() {
	var s string;
	s = "0000000000000000000000000000000000000000000000000000000000"[0:7];
	_ = s;
}

/*
uetli:~/Source/go1/test/bugs gri$ 6g bug061.go
Bus error
*/
