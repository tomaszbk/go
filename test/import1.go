// errorcheck


// Verify that import conflicts are detected by the compiler.
// Does not compile.

package main

import "bufio"	// ERROR "previous|not used"
import bufio "os"	// ERROR "redeclared|redefinition|incompatible" "imported and not used|imported as bufio and not used"

import (
	"fmt"	// ERROR "previous|not used"
	fmt "math"	// ERROR "redeclared|redefinition|incompatible" "imported and not used: \x22math\x22 as fmt|imported as fmt and not used"
	. "math"	// GC_ERROR "imported and not used: \x22math\x22$|imported and not used"
)
