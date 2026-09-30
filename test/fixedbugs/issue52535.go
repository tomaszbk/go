// compile


// gofrontend crashed converting unnamed bool type to any.

package p

func F() {
	m := make(map[int]int)
	var ok any
	_, ok = m[0]
	_ = ok
}
