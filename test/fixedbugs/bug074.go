// errorcheck


package main

func main() {
	x := string{'a', 'b', '\n'};	// ERROR "composite"
	print(x);
}
