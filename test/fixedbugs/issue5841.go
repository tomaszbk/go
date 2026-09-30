// build


// Issue 5841: 8g produces invalid CMPL $0, $0.
// Similar to issue 5002, used to fail at link time.

package main

func main() {
	var y int
	if y%1 == 0 {
	}
}
