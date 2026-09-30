// run

// Bug in method values: escape analysis was off.

package main

import "sync"

var called = false

type T struct {
	once sync.Once
}

func (t *T) M() {
	called = true
}

func main() {
	var t T
	t.once.Do(t.M)
	if !called {
		panic("not called")
	}
}
