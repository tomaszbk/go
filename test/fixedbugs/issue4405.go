// errorcheck


package p

const (
	_ = iota
	_ // ERROR "illegal character|invalid character"
	_  // ERROR "illegal character|invalid character"
	_  // ERROR "illegal character|invalid character"
	_  // ERROR "illegal character|invalid character"
)
