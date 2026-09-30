package p

func _() {
	var x interface{}
	switch t := x.(type) {
	case S /* ERROR "cannot use generic type" */ :
		t.m()
	}
}

type S[T any] struct {}

func (_ S[T]) m()
