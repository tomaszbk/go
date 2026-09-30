// This file contains the code to check canonical methods.

package method

import "fmt"

type MethodTest int

func (t *MethodTest) Scan(x fmt.ScanState, c byte) { // ERROR `should have signature Scan\(fmt\.ScanState, rune\) error`
}
