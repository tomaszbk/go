// Package predeclared is a go/doc test for handling of
// exported methods on locally-defined predeclared types.
// See issue 9860.
package predeclared

type error struct{}

// Must not be visible.
func (e error) Error() string {
	return ""
}

type bool int

// Must not be visible.
func (b bool) String() string {
	return ""
}
