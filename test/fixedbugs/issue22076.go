// compile

// Issue 22076: Couldn't use ":=" to declare names that refer to
// dot-imported symbols.

package p

import . "bytes"

var _ Reader // use "bytes" import

func f1() {
	Buffer := 0
	_ = Buffer
}

func f2() {
	for Buffer := range []int{} {
		_ = Buffer
	}
}
