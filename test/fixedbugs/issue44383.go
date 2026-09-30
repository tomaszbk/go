// compile


// Issue 44383: gofrontend internal compiler error

package main

func main() {
	var b1, b2 byte
	f := func() int {
		var m map[byte]int
		return m[b1/b2]
	}
	f()
}
