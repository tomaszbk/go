// run


// Simple test of the garbage collector.

package main

import "runtime"

func mk2() {
	b := new([10000]byte)
	_ = b
	//	println(b, "stored at", &b)
}

func mk1() { mk2() }

func main() {
	for i := 0; i < 10; i++ {
		mk1()
		runtime.GC()
	}
}
