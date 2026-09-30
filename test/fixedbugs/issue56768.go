// compile


package p

type I interface {
	M()
}

type slice []any

func f() {
	ss := struct{ i I }{}

	_ = [...]struct {
		s slice
	}{
		{
			s: slice{ss.i},
		},
		{
			s: slice{ss.i},
		},
		{
			s: slice{ss.i},
		},
		{
			s: slice{ss.i},
		},
		{
			s: slice{ss.i},
		},
	}
}
