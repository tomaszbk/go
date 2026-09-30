// errorcheck


// Test an internal compiler error on ? symbol in declaration
// following an empty import.

package a
var?      // ERROR "invalid character U\+003F '\?'|invalid character 0x3f in input file"

var x int // ERROR "unexpected keyword var|expected identifier|expected type"

func main() {
}
