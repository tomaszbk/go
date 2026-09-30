// errorcheck

package p

const (
	f = 1.0
	c = 1.0i

	_ = ^f // ERROR "invalid operation|expected integer"
	_ = ^c // ERROR "invalid operation|expected integer"

	_ = f % f // ERROR "invalid operation|expected integer"
	_ = c % c // ERROR "invalid operation|expected integer"

	_ = f & f // ERROR "invalid operation|expected integer"
	_ = c & c // ERROR "invalid operation|expected integer"

	_ = f | f // ERROR "invalid operation|expected integer"
	_ = c | c // ERROR "invalid operation|expected integer"

	_ = f ^ f // ERROR "invalid operation|expected integer"
	_ = c ^ c // ERROR "invalid operation|expected integer"

	_ = f &^ f // ERROR "invalid operation|expected integer"
	_ = c &^ c // ERROR "invalid operation|expected integer"
)
