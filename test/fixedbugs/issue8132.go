// run


// issue 8132. stack walk handling of panic stack was confused
// about what was legal.

package main

import "runtime"

var p *int

func main() {
	func() {
		defer func() {
			runtime.GC()
			recover()
		}()
		var x [8192]byte
		func(x [8192]byte) {
			defer func() {
				if err := recover(); err != nil {
					println(*p)
				}
			}()
			println(*p)
		}(x)
	}()
}
