// compile


package main

func main() {
	var s int = 0
	var x int = 0
	x = x << s // as of 1.13, these are ok
	x = x >> s // as of 1.13, these are ok
}
