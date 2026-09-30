// errorcheck


package main

import . "unsafe"	// ERROR "not used"

func main() {
	var x int
	println(unsafe.Sizeof(x)) // ERROR "undefined"
}

/*
After a '.' import, "unsafe" shouldn't be defined as
an identifier. 6g complains correctly for imports other
than "unsafe".
*/
