// compile


package p

func f() {
	for b := "" < join([]string{}, "") && true; ; {
		_ = b
	}
}

//go:noinline
func join(elems []string, sep string) string {
	return ""
}
