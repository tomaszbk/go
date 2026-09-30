// run


package main

type S []S

func main() {
	var s S
	s = append(s, s) // append a nil value to s
	if s[0] != nil {
		println("BUG: s[0] != nil")
	}
}
