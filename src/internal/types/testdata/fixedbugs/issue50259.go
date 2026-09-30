package p

var x T[B]

type T[_ any] struct{}
type A T[B]
type B = T[A]

// test case from issue

var v Box[Step]
type Box[T any] struct{}
type Step = Box[StepBox]
type StepBox Box[Step]
