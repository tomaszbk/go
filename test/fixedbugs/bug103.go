// errorcheck


package main

func f() /* no return type */ {}

func main() {
	x := f();  // ERROR "mismatch|as value|no type"
	_ = x
}

