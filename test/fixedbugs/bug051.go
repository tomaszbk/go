// errorcheck


package main

func f() int {
	return 0;
}

func main() {
	const n = f();	// ERROR "const"
}
