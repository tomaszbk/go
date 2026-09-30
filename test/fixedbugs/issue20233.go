// errorcheck

// Issue 20233: panic while formatting an error message

package p

var f = func(...A) // ERROR "undefined: A"
