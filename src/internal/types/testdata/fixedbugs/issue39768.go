package p

// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T[P any] P
// type A = T  // ERROR "cannot use generic type"
// var x A[int]
// var _ A
//
// type B = T[int]
// var y B = x
// var _ B /* ERROR "not a generic type" */ [int]

// test case from issue

type Vector[T any] []T
type VectorAlias = Vector // ERROR "cannot use generic type"
var v Vector[int]
