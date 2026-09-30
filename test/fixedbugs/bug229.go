// errorcheck

package main

import "testing"

func main() {
	var t testing.T

	// make sure error mentions that
	// name is unexported, not just "name not found".

	t.common.name = nil // ERROR "unexported|undefined"

	println(testing.anyLowercaseName("asdf")) // ERROR "unexported|undefined"
}
