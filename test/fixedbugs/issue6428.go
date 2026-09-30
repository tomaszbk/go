// errorcheck

package p

import . "testing" // ERROR "imported and not used"

type S struct {
	T int
}

var _ = S{T: 0}
