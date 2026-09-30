// run

package main

type S struct {
	x int
}

func (t *S) M1() {
}

func F[T any](x T) any {
	return x
}

func main() {
	F(&S{}).(interface{ M1() }).M1()
}
