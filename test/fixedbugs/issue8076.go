// compile

// Issue 8076. nilwalkfwd walked forward forever
// on the instruction loop following the dereference.

package main

func main() {
	_ = *(*int)(nil)
L:
	_ = 0
	goto L
}
