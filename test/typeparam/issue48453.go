// run

package main

//go:noinline
func CopyMap[M interface{ ~map[K]V }, K comparable, V any](m M) M {
	out := make(M, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func main() {
	var m map[*string]int
	CopyMap(m)
}
