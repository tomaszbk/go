package a

// Type I is the first basic test for the issue, which relates to a type that is recursive
// via a type constraint.  (In this test, I -> IConstraint -> MyStruct -> I.)
type JsonRaw []byte

type MyStruct struct {
	x *I[JsonRaw]
}

type IConstraint interface {
	JsonRaw | MyStruct
}

type I[T IConstraint] struct {
}
