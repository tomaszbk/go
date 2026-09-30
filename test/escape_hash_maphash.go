// errorcheck -0 -m -l

// Test escape analysis for hash/maphash.

package escape

import (
	"hash/maphash"
)

func f() {
	var x maphash.Hash // should be stack allocatable
	x.WriteString("foo")
	x.Sum64()
}
