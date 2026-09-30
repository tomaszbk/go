// run


// Issue 3899: 8g incorrectly thinks a variable is
// "set and not used" and elides an assignment, causing
// variables to end up with wrong data.
//
// The reason is a miscalculation of variable width.

package main

func bar(f func()) {
	f()
}

func foo() {
	f := func() {}
	if f == nil {
	}
	bar(f)
}

func main() {
	foo()
}
