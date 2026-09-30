// run

package main

// issue 2337
// The program deadlocked.

import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	runtime.GC()
	runtime.GOMAXPROCS(1)
}
