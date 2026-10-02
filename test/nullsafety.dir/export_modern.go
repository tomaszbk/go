package lib

func Pick(p, q *int) *int       { return p ?? q }
func Generic[T ~*int](p, q T) T { return p ?? q }

type Node struct {
	Next *Node
	N    int
}

func Get(p *Node) int                    { return p?.Next?.N ?? 3 }
func Init(m map[int]*int, k int, v *int) { m[k] ??= v }
