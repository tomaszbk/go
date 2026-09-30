// errorcheck


// Use //line to set the line number of the next line to 20.
//line fixedbugs/bug305.go:20

package p

// Introduce an error which should be reported on line 24.
var a int = "bogus"

// Line 15 of file.
// 16
// 17
// 18
// 19
// 20
// 21
// 22
// 23
// ERROR "cannot|incompatible"
