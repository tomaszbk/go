// compile


package p

var (
	x  int
	xs []int
)

func a([]int) (int, error)

func b() (int, error) {
	return a(append(xs, x))
}

func c(int, error) (int, error)

func d() (int, error) {
	return c(b())
}
