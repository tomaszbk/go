// run

package main

type Chan[T any] chan Chan[T]

func (ch Chan[T]) recv() Chan[T] {
	return <-ch
}

func main() {
	ch := Chan[int](make(chan Chan[int]))
	go func() {
		ch <- make(Chan[int])
	}()
	ch.recv()
}
