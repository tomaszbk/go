// skip


package main

var a [1<<31 - 1024]byte

func main() {
	if a[0] != 0 {
		panic("bad array")
	}
}
