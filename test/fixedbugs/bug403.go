// compile


// Crashed gccgo.

package p

type S struct {
	f interface{}
}

func F(p *S) bool {
	v := p.f
	switch a := v.(type) {
	case nil:
		_ = a
		return true
	}
	return true
}
