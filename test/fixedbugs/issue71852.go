// compile


package main

import (
	"math"
)

func main() {
	test(2)
}

func test(i int) {
	if i <= 0 {
		return
	}

	_ = math.Pow10(i + 2)
}
