// errorcheck


package main

func f(args ...int) {
	g(args)
}

func g(args ...interface{}) {
	f(args)	// ERROR "cannot use|incompatible"
}
