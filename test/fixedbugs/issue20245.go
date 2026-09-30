// errorcheck -d=panic

// Issue 20245: panic while formatting an error message

package p

var e = interface{ I1 } // ERROR "undefined: I1"
