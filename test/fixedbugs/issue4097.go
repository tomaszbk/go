// errorcheck


package foo

var s [][10]int
const m = len(s[len(s)-1]) // ERROR "is not a constant|is not constant" 

