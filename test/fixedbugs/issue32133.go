// errorcheck


package p

// errors for the //line-adjusted code below
// ERROR "newline in string"
// ERROR "newline in character literal|newline in rune literal"
// ERROR "newline in string"
// ERROR "string not terminated"

//line :7:1
import "foo

//line :16:1
func _() {
	0x // ERROR "hexadecimal literal has no digits"
}

func _() {
	0x1.0 // ERROR "hexadecimal mantissa requires a 'p' exponent"
}

func _() {
	0_i // ERROR "'_' must separate successive digits"
}

func _() {
//line :8:1
	'
}

func _() {
//line :9:1
	"
}

func _() {
//line :10:1
	`