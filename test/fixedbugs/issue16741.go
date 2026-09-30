// compile

// Make sure CSE of multi-output opcodes works correctly
// with select0/1 operations.

package main

func div(d, r int64) int64 {
	if m := d % r; m > 0 {
		return d/r + 1
	}
	return d / r
}
