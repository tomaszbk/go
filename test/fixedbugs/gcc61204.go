// compile

// PR61204: Making temporaries for zero-sized types caused an ICE in gccgo.
// This is a reduction of a program reported by GoSmith.

package main

func main() {
	type t [0]int
	var v t
	v, _ = [0]int{}, 0
	_ = v
}
