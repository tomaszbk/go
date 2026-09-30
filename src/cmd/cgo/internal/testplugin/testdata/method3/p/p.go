package p

type T int

func (T) m() { println("m") }

type I interface{ m() }

func F() {
	i.m()
}

var i I = T(123)
