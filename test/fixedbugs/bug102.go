// run

package main

func main() {
	var b [0]byte
	s := string(b[0:]) // out of bounds trap
	if s != "" {
		panic("bad convert")
	}
	var b1 = [5]byte{'h', 'e', 'l', 'l', 'o'}
	if string(b1[0:]) != "hello" {
		panic("bad convert 1")
	}
	var b2 = make([]byte, 5)
	for i := 0; i < 5; i++ {
		b2[i] = b1[i]
	}
	if string(b2) != "hello" {
		panic("bad convert 2")
	}
}
