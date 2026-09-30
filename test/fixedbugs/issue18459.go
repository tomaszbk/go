// errorcheck


// Verify that we have a line number for this error.

package main

//go:nowritebarrier // ERROR "//go:nowritebarrier only allowed in runtime"
func main() {
}
