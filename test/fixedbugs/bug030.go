// errorcheck


package main

func main() {
	var x int;
	x := 0; // ERROR "declar|:="
	_ = x;
}
