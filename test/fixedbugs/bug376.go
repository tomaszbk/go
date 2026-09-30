// errorcheck


// issue 1951
package foo
import "unsafe"
var v = unsafe.Sizeof  // ERROR "not in function call|must be called"
