// errorcheck


// Verify that the array is reported in correct notation.

package p

var a [len(a)]int // ERROR "\[len\(a\)\]int|initialization cycle"
