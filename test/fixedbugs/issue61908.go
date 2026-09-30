// compile


package p

func f(p []byte) int {
	switch "" < string(p) {
	case true:
		return 0
	default:
		return 1
	}
}
