// run


// Test using _ receiver.  Failed with gccgo.

package main

type S struct {}

func (_ S) F(i int) int {
	return i
}

func main() {
	s := S{}
	const c = 123
	i := s.F(c)
	if i != c {
		panic(i)
	}
}
