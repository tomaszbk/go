// errorcheck

// Ensure that typed non-integer, negative and too large
// values are not accepted as size argument in make for
// channels.

package main

type T chan byte

var sink T

func main() {
	sink = make(T, -1)            // ERROR "negative buffer argument in make.*|must not be negative"
	sink = make(T, uint64(1<<63)) // ERROR "buffer argument too large in make.*|overflows int"

	sink = make(T, 0.5) // ERROR "constant 0.5 truncated to integer|truncated to int"
	sink = make(T, 1.0)
	sink = make(T, float32(1.0)) // ERROR "non-integer buffer argument in make.*|must be integer"
	sink = make(T, float64(1.0)) // ERROR "non-integer buffer argument in make.*|must be integer"
	sink = make(T, 1+0i)
	sink = make(T, complex64(1+0i))  // ERROR "non-integer buffer argument in make.*|must be integer"
	sink = make(T, complex128(1+0i)) // ERROR "non-integer buffer argument in make.*|must be integer"
}
