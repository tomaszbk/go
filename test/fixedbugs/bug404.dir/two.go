// The gccgo compiler would fail on the import statement.
// two.go:10:13: error: use of undefined type ‘one.T2’

package two

import "./one"

var V one.T3
