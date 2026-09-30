package p

func F(s string) {
	switch s[0] {
	case 'a':
		case s[2] { // ERROR unexpected {
		case 'b':
		}
	}
} // ERROR non-declaration statement
