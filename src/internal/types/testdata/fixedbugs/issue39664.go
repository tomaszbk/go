package p

type T[_ any] struct {}

func (T /* ERROR "instantiation" */ ) m()

func _() {
	var x interface { m() }
	x = T[int]{}
	_ = x
}
