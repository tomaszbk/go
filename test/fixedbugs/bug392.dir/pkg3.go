// Use the functions in pkg2.go so that the inlined
// forms get type-checked.

package pkg3

import "./pkg2"

var x = pkg2.F()
var v = pkg2.V
