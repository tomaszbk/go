// run


package main

import "math"

var doNotFold = 18446744073709549568.0

func main() {
	if math.Trunc(doNotFold) != doNotFold {
		panic("big (over 2**63-1) math.Trunc is incorrect")
	}
}
