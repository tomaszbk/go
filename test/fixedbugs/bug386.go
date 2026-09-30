// errorcheck


// Issue 2451, 2452 
package foo

func f() error { return 0 } // ERROR "cannot use 0 (.type int.)?|has no methods"

func g() error { return -1 }  // ERROR "cannot use -1 (.type int.)?|has no methods"
