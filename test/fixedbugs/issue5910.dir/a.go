package a

type Package struct {
	name string
}

type Future struct {
	result chan struct {
		*Package
		error
	}
}

func (t *Future) Result() (*Package, error) {
	result := <-t.result
	t.result <- result
	return result.Package, result.error
}
