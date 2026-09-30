// compile


// The gccgo compiler crashed while compiling this code.
// https://gcc.gnu.org/PR78763.

package p

import "unsafe"

func F() int {
	if unsafe.Sizeof(0) == 8 {
		return 8
	}
	return 0
}
