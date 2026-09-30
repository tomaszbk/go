package main

import "testshared/iface_a"
import "testshared/iface_b"

func main() {
	if iface_a.F() != iface_b.F() {
		panic("empty interfaces not equal")
	}
	if iface_a.G() != iface_b.G() {
		panic("non-empty interfaces not equal")
	}
}
