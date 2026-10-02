package main

import "fmt"

func must(ok bool, message string) {
	if !ok {
		panic(message)
	}
}
func Map[T, U any](xs []T, f func(T) U) []U {
	out := make([]U, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}
func Reduce[T, A any](xs []T, a A, f func(A, T) A) A {
	for _, x := range xs {
		a = f(a, x)
	}
	return a
}
func Pair[T, U any](x T, y U, f func(T) U) U { return f(x) }
func Pick[T any](f func() T, x T) T          { return f() }
func Twice[T any, F ~func(T) T](x T, f F) T  { return f(f(x)) }

type User struct{ Name string }
type Callback func(int) int

func main() { scenarios(); fmt.Println("PASS") }

func noop()         {}
func failed() error { return fmt.Errorf("failure") }

func options(n int, fs ...func(int) int) int {
	for _, f := range fs {
		n = f(n)
	}
	return n
}
