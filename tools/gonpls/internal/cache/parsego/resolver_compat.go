// This file contains declarations needed for compatibility with resolver.go
// copied from GOROOT.

package parsego

import "go/token"

// assert panics with the given msg if cond is not true.
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// A bailout panic is raised to indicate early termination. pos and msg are
// only populated when bailing out of object resolution.
type bailout struct {
	pos token.Pos
	msg string
}
