package a

type A[T any] = struct{ F T }

type B = struct{ F int }

func F() B {
	type a[T any] = struct{ F T }
	return a[int]{}
}
