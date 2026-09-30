// This file contains the test for untagged struct literals.

package composite

import "cmd/vet/testdata/composite/a"

// Testing is awkward because we need to reference things from a separate package
// to trigger the warnings.

var goodStructLiteral = a.Flag{
	Name:  "Name",
	Usage: "Usage",
}

var badStructLiteral = a.Flag{ // ERROR "unkeyed fields"
	"Name",
	"Usage",
	nil, // Value
	"DefValue",
}
