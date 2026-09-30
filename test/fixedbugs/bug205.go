// errorcheck


package main

var t []int
var s string;
var m map[string]int;

func main() {
	println(t["hi"]); // ERROR "non-integer slice index|must be integer|cannot convert"
	println(s["hi"]); // ERROR "non-integer string index|must be integer|cannot convert"
	println(m[0]);    // ERROR "cannot use.*as type string|cannot convert|cannot use"
}

