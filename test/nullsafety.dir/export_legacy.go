package lib

func Pick(p, q *int) *int {
	if p != nil {
		return p
	}
	return q
}
func Generic[T ~*int](p, q T) T {
	if p != nil {
		return p
	}
	return q
}

type Node struct {
	Next *Node
	N    int
}

func Get(p *Node) int {
	if p != nil && p.Next != nil {
		return p.Next.N
	}
	return 3
}
func Init(m map[int]*int, k int, v *int) {
	if m[k] == nil {
		m[k] = v
	}
}
