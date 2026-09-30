// run


package main

//go:noinline
func f(p *int, b bool) int {
	valid := *p >= 0
	if !b || !valid {
		return 5
	}
	return 6
}
func main() {
	defer func() {
		if e := recover(); e == nil {
			println("should have panicked")
		}
	}()
	f(nil, false)
}
