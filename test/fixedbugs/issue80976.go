// run


// Instantiated method type arguments might refer to a type instantiation that
// ends with the method name. Here, Set is instantiated using a type which
// refers to HashSet[int].

package main

type TypeMap struct{}

//go:noinline
func (tm *TypeMap) Set[T any](v T) {}

type HashSet[T any] struct{}

type res struct{ s HashSet[int] }

func main() { (&TypeMap{}).Set(res{}) }
