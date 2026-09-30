// errorcheck


package main

func main() {
	var _ string = nil // ERROR "illegal|invalid|incompatible|cannot"
}
