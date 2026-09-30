package p

type A[P any] = P // ERROR "cannot use type parameter declared in alias declaration as RHS"

func _[P any]() {
	type A[P any] = P // ERROR "cannot use type parameter declared in alias declaration as RHS"
	type B = P
	type C[Q any] = P
}
