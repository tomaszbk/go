// Use the functions in one.go so that the inlined
// forms get type-checked.

package two

import "./one"

func use() {
	var r one.T
	r.F()
}
