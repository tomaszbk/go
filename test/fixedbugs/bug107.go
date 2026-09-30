// errorcheck


package main
import os "os"
type _ os.FileInfo
func f() (os int) {
	 // In the next line "os" should refer to the result variable, not
	 // to the package.
	 v := os.Open("", 0, 0);	// ERROR "undefined"
	 _ = v
	 return 0
}
