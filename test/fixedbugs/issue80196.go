// run


package main

type S [5]*byte

//go:noinline
func f() S {
	return S{}
}

var sink *S

//go:noinline
func g() (s S) {
	sink = &s
	s = f()
	return
}

func main() {
	for i := range 1000000 {
		s := g()
		if s != (S{}) {
			println(s[0], s[1], s[2], s[3], s[4])
			panic(i)
		}
	}
}
