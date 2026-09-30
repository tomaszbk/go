// compile


package p

type I interface {
	M()
}

type S struct{}

func (*S) M() {}

type slice []I

func f() {
	ss := struct {
		i I
	}{
		i: &S{},
	}

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
