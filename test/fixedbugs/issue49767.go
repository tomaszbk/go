// errorcheck


package main

func main() {
	ch := make(chan struct{ v [65536]byte }) // ERROR "channel element type too large"
	close(ch)
}
