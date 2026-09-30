// errorcheck


package main

func main() {
	s = "bob" // ERROR "undefined.*s"
	_ = s // ERROR "undefined.*s"
}
