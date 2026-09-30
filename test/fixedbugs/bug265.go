// run

// Test case for https://golang.org/issue/700

package main

import "os"

func f() (e int) {
	_ = &e
	return 999
}

func main() {
	if f() != 999 {
		os.Exit(1)
	}
}
