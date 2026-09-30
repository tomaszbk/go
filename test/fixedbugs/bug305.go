// errorcheck


// Use //line to set the line number of the next line to 17.
//line fixedbugs/bug305.go:17

package p

// Introduce an error which should be reported on line 21.
var a int = "bogus"

// Line 12 of file.
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// 20
// ERROR "cannot|incompatible"
