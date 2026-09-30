// run

package main

const (
	intSize = 32 << (^uint(0) >> 63)
	minInt  = -1 << (intSize - 1)
)

func main() {
	f()
}

func f() {
	for i := 0; true; i += minInt {
		if i < 0 {
			return
		}
	}
	panic("unreachable")
}
