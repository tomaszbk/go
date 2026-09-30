// compile


package main

type Cache[E comparable] struct {
	adder func(...E)
}

func New[E comparable]() *Cache[E] {
	c := &Cache[E]{}

	c.adder = func(elements ...E) {
		for _, value := range elements {
			value := value
			go func() {
				println(value)
			}()
		}
	}

	return c
}

func main() {
	c := New[string]()
	c.adder("test")
}
