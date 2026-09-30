// compile


// Test that when import gives multiple names
// to a single type, they still all refer to the same type.

package main

import _os_ "os"
import "os"
import . "os"

func f(e *os.File)

func main() {
	var _e_ *_os_.File
	var dot *File

	f(_e_)
	f(dot)
}
