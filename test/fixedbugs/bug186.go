// errorcheck


package main

const X = iota

func f(x int) { }

func main() {
	f(X);
	f(iota);	// ERROR "iota"
	f(X);
}
