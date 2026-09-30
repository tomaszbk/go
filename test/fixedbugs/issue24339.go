// errorcheck


package p

// Use a different line number for each token so we can
// check that the error message appears at the correct
// position.
var _ = struct{}{ /*line :17:1*/foo /*line :18:1*/: /*line :19:1*/0 }







// ERROR "unknown field foo"
