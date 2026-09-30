// errorcheck


package main

func x() {
}

func main() {
	if {  // ERROR "missing condition"
	}
	
	if x(); {  // ERROR "missing condition"
	}
}
