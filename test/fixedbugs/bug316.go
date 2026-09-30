// compile


// Issue 1369.

package main

const (
	c = complex(1, 2)
	r = real(c) // was: const initializer must be constant
	i = imag(c) // was: const initializer must be constant
)

func main() {}
