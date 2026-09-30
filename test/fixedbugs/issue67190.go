// run


package main

func main() {
	ch1 := make(chan struct{})
	var ch2 <-chan struct{} = ch1

	switch ch1 {
	case ch2:
	default:
		panic("bad narrow case")
	}

	switch ch2 {
	case ch1:
	default:
		panic("bad narrow switch")
	}
}
