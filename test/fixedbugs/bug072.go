// errorcheck


package main

func main() {
	s := string(bug);  // ERROR "undef"
	_ = s
}
