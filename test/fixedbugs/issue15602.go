// compile


package p

func f(i interface{}) {
	i, _ = i.(error)
}
