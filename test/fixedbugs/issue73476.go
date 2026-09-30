// run


package main

//go:noinline
func f(p *[4]int) {
	for i := range (*p) { // Note the parentheses! gofmt wants to remove them - don't let it!
		println(i)
	}
}
func main() {
	f(nil)
}
