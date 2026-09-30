// compile

package main

func g() bool

func f(y int) bool {
	b, ok := true, false
	if y > 1 {
		ok = g()
	}
	if !ok {
		ok = g()
		b = false
	}
	if !ok {
		return false
	}
	return b
}
