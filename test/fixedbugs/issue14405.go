// compile

// Mention of field with large offset in struct literal causes crash
package p

type T struct {
	Slice [1 << 20][]int
	Ptr   *int
}

func New(p *int) *T {
	return &T{Ptr: p}
}
