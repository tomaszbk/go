package issue25301

type (
	A = interface {
		M()
	}
	T interface {
		A
	}
	S struct{}
)

func (S) M() { println("m") }
