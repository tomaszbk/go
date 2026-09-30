package bresource

type Resource[T any] struct {
	name        string
	initializer Initializer[T]
	cfg         ResConfig
	value       T
}

func Should[T any](r *Resource[T], e error) bool {
	return r.cfg.ShouldRetry(e)
}

type ResConfig struct {
	ShouldRetry func(error) bool
	TearDown    func()
}

type Initializer[T any] func(*int) (T, error)

func New[T any](name string, f Initializer[T], cfg ResConfig) *Resource[T] {
	return &Resource[T]{name: name, initializer: f, cfg: cfg}
}
