// errorcheck

// Use //line to set the line number of the next line to 16.
//line fixedbugs/bug305.go:16

package p

// Introduce an error which should be reported on line 20.
var a int = "bogus"

// Line 11 of file.
// 12
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// ERROR "cannot|incompatible"
