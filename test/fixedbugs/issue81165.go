// build


package main

type someType struct{}
type someTypeAlias = *someType

//go:noinline
func (tm someTypeAlias) FailedLink[T any](v T) {
	println(v)
}

//go:noinline
func (tm *someType) SuccessLink[T any](v T) {
	println(v)
}

func main() {
	var value someType
	value.SuccessLink(42)
	value.FailedLink(42)
}
