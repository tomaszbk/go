// build

// Issue 5002: 8g produces invalid CMPL $0, $0.
// Used to fail at link time.

package main

func main() {
	var y int64
	if y%1 == 0 {
	}
}
