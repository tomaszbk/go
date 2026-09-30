// errorcheck


// Test an internal compiler error on ? symbol in declaration
// following an empty import.

package a
import""  // ERROR "import path is empty|invalid import path \(empty string\)"
