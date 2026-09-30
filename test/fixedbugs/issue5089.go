// errorcheck


// issue 5089: gc allows methods on non-locals if symbol already exists

package p

import "bufio"

func (b *bufio.Reader) Buffered() int { // ERROR "non-local|redefinition"
	return -1
}
