// run


// Issue 6902: confusing printing of large floating point constants

package main

import (
	"os"
)

var x = -1e-10000

func main() {
	if x != 0 {
		os.Exit(1)
	}
}
