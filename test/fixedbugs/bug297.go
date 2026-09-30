// errorcheck -d=panic


// Used to crash; issue 961.

package main

type ByteSize float64

const (
	_           = iota          // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * X) // ERROR "undefined"
)
