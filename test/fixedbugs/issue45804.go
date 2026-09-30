// errorcheck

package p

func g() int
func h(int)

var b bool

func f() {
	did := g()
	if !did && b { // ERROR "invalid operation"
		h(x) // ERROR "undefined"
	}
}
