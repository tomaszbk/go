// run


// https://golang.org/issue/843

package main

import "unsafe"

type T struct {
	X, Y uint8
}

func main() {
	var t T
	if unsafe.Offsetof(t.X) != 0 || unsafe.Offsetof(t.Y) != 1 {
		println("BUG", unsafe.Offsetof(t.X), unsafe.Offsetof(t.Y))
	}
}
