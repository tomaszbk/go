// run -tags=purego -gcflags=all=-d=checkptr

package main

import "crypto/subtle"

func main() {
	dst := make([]byte, 5)
	src := make([]byte, 5)
	for _, n := range []int{1024, 2048} { // just to make the size non-constant
		b := make([]byte, n)
		subtle.XORBytes(dst, src, b[n-5:])
	}
}
