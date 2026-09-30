package cmem

import (
	"./aconfig"
	"./bresource"
)

type MemT *int

var G int

type memResource struct {
	x *int
}

func (m *memResource) initialize(*int) (res *int, err error) {
	return nil, nil
}

func (m *memResource) teardown() {
}

func NewResource(cfg *aconfig.Config) *bresource.Resource[*int] {
	res := &memResource{
		x: &G,
	}

	return bresource.New("Mem", res.initialize, bresource.ResConfig{
		// We always would want to retry the Memcache initialization.
		ShouldRetry: func(error) bool { return true },
		TearDown:    res.teardown,
	})
}
