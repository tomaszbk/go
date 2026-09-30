package main

func undefined()

func defined1() int {
	// To check multiple errors for a single symbol,
	// reference undefined more than once.
	undefined()
	undefined()
	return 0
}

func defined2() {
	undefined()
	undefined()
}

func init() {
	_ = defined1()
	defined2()
}

// The "main" function remains undeclared.
