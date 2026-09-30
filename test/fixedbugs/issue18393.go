// errorcheck


// Test that compiler directives are ignored if they
// don't start at the beginning of the line.

package p

//line issue18393.go:20
import 42 // error on line 20


/* //line not at start of line: ignored */ //line issue18393.go:30
var x     // error on line 24, not 30


// ERROR "import path must be a string"



// ERROR "syntax error: unexpected newline, expecting type|expected type"
