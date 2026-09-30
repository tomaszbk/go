// This file contains tests for the unusedresult checker.

package unused

import "fmt"

func _() {
	fmt.Errorf("") // ERROR "result of fmt.Errorf call not used"
}
