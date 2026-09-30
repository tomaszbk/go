// run

package main

import "unsafe"

const (
	_ = unsafe.Sizeof(func() int {
		const (
			_ = 1
			_
			_
		)
		return 0
	}())

	y = iota
)

func main() {
	if y != 1 {
		panic(y)
	}
}
