// compile


// Returning an index into a conversion from string to slice caused a
// compilation error when using gccgo.

package p

func F1(s string) byte {
	return []byte(s)[0]
}

func F2(s string) rune {
	return []rune(s)[0]
}
