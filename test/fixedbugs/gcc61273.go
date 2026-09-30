// compile


// PR61273: gccgo failed to compile a SendStmt in the PostStmt of a ForClause
// that involved predefined constants.

package main

func main() {
	c := make(chan bool, 1)
	for ; false; c <- false {
	}
}
