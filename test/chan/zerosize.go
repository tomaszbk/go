// run


// Test making channels of a zero-sized type.

package main

func main() {
	_ = make(chan [0]byte)
	_ = make(chan [0]byte, 1)
	_ = make(chan struct{})
	_ = make(chan struct{}, 1)
}
