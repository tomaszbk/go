// errorcheck


package main
func f(a T) T { return a }	// ERROR "undefined"
func main() {
	x := f(0);
	_ = x;
}
