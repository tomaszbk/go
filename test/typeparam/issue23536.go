// run


// Test case where a slice of a user-defined byte type (not uint8 or byte) is
// converted to a string.  Same for slice of runes.

package main

type MyByte byte

type MyRune rune

func f[T []MyByte](x T) string {
	return string(x)
}

func g[T []MyRune](x T) string {
	return string(x)
}

func main() {
	var y []MyByte
	_ = f(y)
	_ = string(y)

	var z []MyRune
	_ = g(z)
	_ = string(z)
}
