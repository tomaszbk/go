// run

package main

// This used to print 0, because x was incorrectly captured by value.

func f() (x int) {
	defer func() func() {
		return func() {
			println(x)
		}
	}()()
	return 42
}

func main() {
	f()
}
