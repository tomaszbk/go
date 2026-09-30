// compile

// gccgo gave an invalid error ("floating point constant truncated to
// integer") compiling this.

package p

const C = 1<<63 - 1

func F(i int64) int64 {
	return i
}

var V = F(int64(C) / 1e6)
