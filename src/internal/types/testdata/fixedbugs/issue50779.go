package p

type AC interface {
	C
}

type ST []int

type R[S any, P any] struct{}

type SR = R[SS, ST]

type SS interface {
	NSR(any) *SR
}

type C interface {
	NSR(any) *SR
}
