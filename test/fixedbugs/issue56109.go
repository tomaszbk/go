// compile


package main

import "math"

func main() {
	f := func(p bool) {
		if p {
			println("hi")
		}
	}
	go f(true || math.Sqrt(2) > 1)
}
