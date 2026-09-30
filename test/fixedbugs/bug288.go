// compile


// Used to run out of registers on 8g.  Issue 868.

package main

func main() {
	var r uint32
	var buf [4]byte
	a := buf[0:4]
	r = (((((uint32(a[3]) << 8) | uint32(a[2])) << 8) |
		uint32(a[1])) << 8) | uint32(a[0])
	_ = r
}
