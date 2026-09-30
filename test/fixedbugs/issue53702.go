// run

package main

type Elem struct{}

func (*Elem) Wait(callback func()) {}

type Base struct {
	elem [8]*Elem
}

var g_val = 1

func (s *Base) Do() *int {
	resp := &g_val
	for _, e := range s.elem {
		e.Wait(func() {
			*resp = 0
		})
	}
	return resp
}

type Sub struct {
	*Base
}

func main() {
	a := Sub{new(Base)}
	resp := a.Do()
	if resp != nil && *resp != 1 {
		panic("FAIL")
	}
}
