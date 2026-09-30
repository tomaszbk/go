package a

type S interface{
	F() T
}

type T struct {
	S
}

type U struct {
	error
}
