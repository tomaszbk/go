// errorcheck -d=panic


package main

// Issue 2623
var m = map[string]int{
	"abc": 1,
	1:     2, // ERROR "cannot use 1.*as type string in map key|incompatible type|cannot convert|cannot use"
}
