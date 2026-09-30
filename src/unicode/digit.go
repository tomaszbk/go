package unicode

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	if r <= MaxLatin1 {
		return '0' <= r && r <= '9'
	}
	return isExcludingLatin(Digit, r)
}
