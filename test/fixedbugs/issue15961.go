// compile


package y

type symSet []int

//go:noinline
func (s symSet) len() (r int) {
	return 0
}

func f(m map[int]symSet) {
	var symSet []int
	for _, x := range symSet {
		m[x] = nil
	}
}
