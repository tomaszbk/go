// errorcheck


// Tests correct reporting of line numbers for errors involving iota,
// Issue #8183.
package foo

const (
	ok = byte(iota + 253)
	bad
	barn
	bard // ERROR "constant 256 overflows byte|integer constant overflow|cannot convert"
)

const (
	c = len([1 - iota]int{})
	d
	e // ERROR "array bound must be non-negative|negative array bound|invalid array length"
	f // ERROR "array bound must be non-negative|negative array bound|invalid array length"
)
