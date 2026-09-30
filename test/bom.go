// runoutput

// Test source file beginning with a byte order mark.

package main

import (
	"fmt"
	"strings"
)

func main() {
	prog = strings.Replace(prog, "BOM", "\uFEFF", -1)
	fmt.Print(prog)
}

var prog = `BOM
package main

func main() {
}
`
