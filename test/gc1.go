// run

// A simple test of the garbage collector.

package main

func main() {
	for i := 0; i < 1e5; i++ {
		x := new([100]byte)
		_ = x
	}
}
