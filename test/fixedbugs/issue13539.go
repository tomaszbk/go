// errorcheck


// Verify that a label named like a package is recognized
// as a label rather than a package and that the package
// remains unused.

package main

import "math" // ERROR "imported and not used"

func main() {
math:
	for {
		break math
	}
}
