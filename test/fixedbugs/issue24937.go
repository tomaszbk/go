// run


package main

func main() {
	x := []byte{'a'}
	switch string(x) {
	case func() string { x[0] = 'b'; return "b" }():
		panic("FAIL")
	}
}
