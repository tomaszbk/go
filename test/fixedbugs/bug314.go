// run


// Used to call wrong methods; issue 1290.

package main

type S struct {
}
func (S) a() int{
	return 0
}
func (S) b() int{
	return 1
}

func main() {
	var i interface {
		b() int
		a() int
	} = S{}
	if i.a() != 0 {
		panic("wrong method called")
	}
	if i.b() != 1 {
		panic("wrong method called")
	}
}
