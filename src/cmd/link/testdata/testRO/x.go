// Test that read-only data is indeed read-only. This
// program attempts to modify read-only data, and it
// should fail.

package main

import "unsafe"

var s = "hello"

func main() {
	println(s)
	*(*struct {
		p *byte
		l int
	})(unsafe.Pointer(&s)).p = 'H'
	println(s)
}
