// Test that embedded interface types can have local methods.

package p

type T int

func (t T) m() {}

type I interface{ m() }
type J interface{ I }

func main() {
	var i I
	var j J
	var t T
	i = t
	j = t
	_ = i
	_ = j
	i = j
	_ = i
	j = i
	_ = j
}
