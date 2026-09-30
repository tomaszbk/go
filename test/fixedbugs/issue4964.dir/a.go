package a

var global, global2 *int

type T struct {
	Pointer *int
}

//go:noinline
func Store(t *T) {
	global = t.Pointer
}

//go:noinline
func Store2(t *T) {
	global2 = t.Pointer
}

func Get() *int {
	return global
}
