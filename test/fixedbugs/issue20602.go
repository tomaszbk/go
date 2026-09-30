// errorcheck


// Verify that the correct (not implicitly dereferenced)
// type is reported in the error message.

package p

var p = &[1]complex128{0}
var _ = real(p)  // ERROR "type \*\[1\]complex128|argument must have complex type"
var _ = imag(p)	 // ERROR "type \*\[1\]complex128|argument must have complex type"
