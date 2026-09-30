// errorcheck


// Check that we don't ignore EOF.

package p

var f = func() { // ERROR "unexpected EOF|expected .*}.*"