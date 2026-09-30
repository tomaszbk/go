// compile


package main

type Collector[T any] struct {
}

func (c *Collector[T]) Collect() {
}

func TestInOrderIntTree() {
	collector := Collector[int]{}
	_ = collector.Collect
}

func main() {
	TestInOrderIntTree()
}
