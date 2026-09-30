// run


// Verify that a label name matching a constant name
// is permitted.

package main

const labelname = 1

func main() {
	goto labelname
labelname:
}
