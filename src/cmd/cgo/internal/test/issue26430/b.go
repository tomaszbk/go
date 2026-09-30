package a

// typedef struct S ST;
// struct S { int f; };
import "C"

func F2(p *C.ST) {
	p.f = 1
}
