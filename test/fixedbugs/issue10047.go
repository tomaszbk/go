// compile

// Issue 10047: gccgo failed to compile a type switch where the switch variable
// and the base type of a case share the same identifier.

package main

func main() {
	type t int
	var p interface{}
	switch t := p.(type) {
	case t:
		_ = t
	}
}
