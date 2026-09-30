// compile


package p

type T struct{}

var (
	t = T{}
	u = t.New()
)

func x(T) (int, int) { return 0, 0 }

var _, _ = x(u)

func (T) New() T { return T{} }
