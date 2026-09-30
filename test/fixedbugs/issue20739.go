// compile


package p

func F() {
	var x struct {
		x *int
		w [1e9][1e9][1e9][0]*int
		y *int
	}
	println(&x)
}
