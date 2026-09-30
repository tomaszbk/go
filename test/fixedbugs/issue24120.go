// compile

package p

var F func(int)

func G() {
	if F(func() int { return 1 }()); false {
	}
}
