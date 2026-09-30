package a

type A struct {
	New func() any
}

func NewA(i int) *A {
	return &A{
		New: func() any {
			_ = i
			return nil
		},
	}
}
