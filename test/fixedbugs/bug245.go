// compile


package main

type T []int
func (t T) m()

func main() {
	_ = T{}
}

// bug245.go:14: fatal error: method mismatch: T for T
