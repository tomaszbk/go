// compile


package p

type Q[T any] struct {
	s []T
}

func (q *Q[T]) Push(v ...T) {
	q.s = append(q.s, v...)
}

func pushN(push func(*Q[int], ...int), n int) {
	var q Q[int]
	for i := 0; i < n; i++ {
		push(&q, i)
	}
}

func f() {
	pushN((*Q[int]).Push, 100)
}
