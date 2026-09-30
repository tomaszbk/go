// errorcheck

// Verify that we have a line number for this error.

package main

//go:cgo_ldflag // ERROR "usage: //go:cgo_ldflag"
func main() {
}
