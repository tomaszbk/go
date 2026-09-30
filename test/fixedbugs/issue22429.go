// compile


// Make sure SSA->assembly pass can handle SP as an index register.

package p

type T struct {
	a,b,c,d float32
}

func f(a *[8]T, i,j,k int) float32 {
	b := *a
	return b[i].a + b[j].b + b[k].c
}
