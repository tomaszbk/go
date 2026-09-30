// compile

// Issue 16733: don't fold constant factors into a multiply
// beyond the capacity of a MULQ instruction (32 bits).

package p

func f(n int64) int64 {
	n *= 1000000
	n *= 1000000
	return n
}
