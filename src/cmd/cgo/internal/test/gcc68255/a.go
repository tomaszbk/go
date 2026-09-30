// Test that it's OK to have C code that does nothing other than
// initialize a global variable. This used to fail with gccgo.

package gcc68255

/*
#include "c.h"
*/
import "C"

func F() bool {
	return C.v != nil
}
