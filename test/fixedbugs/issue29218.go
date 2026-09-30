// compile

package p

type T struct {
	b bool
	string
}

func f() {
	var b bool
	var t T
	for {
		switch &t.b {
		case &b:
			if b {
			}
		}
	}
}
