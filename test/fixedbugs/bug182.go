// errorcheck


package main

func main() {
	x := 0;
	if x {	// ERROR "x.*int|bool"
	}
}
