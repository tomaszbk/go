// run

package main

type M map[int]int

var sink any

func main() {
	sink = M{}
	p := new([8]int)
	m := make(map[any]int)
	m[p] = 42
	if m[p] != 42 {
		panic("incorrect lookup")
	}
}
