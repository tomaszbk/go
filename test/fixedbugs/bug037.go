// errorcheck


package main

func main() {
	s := vlong(0);  // ERROR "undef"
	_ = s
}
