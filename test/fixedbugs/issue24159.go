// errorcheck

package p

type intAlias = int

func f() {
	switch interface{}(nil) {
	case uint8(0):
	case byte(0): // ERROR "duplicate case"
	case int32(0):
	case rune(0): // ERROR "duplicate case"
	case int(0):
	case intAlias(0): // ERROR "duplicate case"
	}
}
