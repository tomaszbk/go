// errorcheck -0 -l -d=wb

// Test write barrier for implicit assignments to result parameters
// that have escaped to the heap.

package issue13587

import "errors"

func escape(p *error)

func F() (err error) {
	escape(&err)
	return errors.New("error") // ERROR "write barrier"
}
