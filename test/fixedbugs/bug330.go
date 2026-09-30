// errorcheck


package main

func main() {
	x := ""
	x = +"hello"  // ERROR "invalid operation.*string|expected numeric"
	x = +x  // ERROR "invalid operation.*string|expected numeric"
}
