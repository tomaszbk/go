// compile

// Check that we don't crash due to "lost track of variable in
// liveness" errors against unused variables.

package p

import "strings"

// Minimized test case from github.com/mvdan/sh/syntax.
func F() {
	var _ = []string{
		strings.Repeat("\n\n\t\t        \n", 10) +
			"# " + strings.Repeat("foo bar ", 10) + "\n" +
			strings.Repeat("longlit_", 10) + "\n",
	}
}
