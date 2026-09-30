// errorcheck


package p

// These error messages are for the invalid literals on lines 16 and 17:

// ERROR "newline in character literal|newline in rune literal"
// ERROR "invalid character literal \(missing closing '\)|rune literal not terminated"

const (
	_ = ''     // ERROR "empty character literal or unescaped ' in character literal|empty rune literal"
	_ = 'f'
	_ = 'foo'  // ERROR "invalid character literal \(more than one character\)|more than one character in rune literal"
//line issue15611.go:8
	_ = '
	_ = '