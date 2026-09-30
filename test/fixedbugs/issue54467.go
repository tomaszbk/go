// run


package main

import "fmt"

func main() {
	var x [64]byte
	for i := range x {
		x[i] = byte(i)
	}
	y := x

	copy(x[4:36], x[2:34])
	*(*[32]byte)(y[4:36]) = *(*[32]byte)(y[2:34])

	for i := range x {
		if x[i] != y[i] {
			fmt.Printf("x[%v] = %v; y[%v] = %v\n", i, x[i], i, y[i])
		}
	}
}
