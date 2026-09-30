package a

func Marshal[foobar any]() {
	_ = NewEncoder[foobar]()
}

func NewEncoder[foobar any]() *Encoder[foobar] {
	return nil
}

type Encoder[foobar any] struct {
}

func (e *Encoder[foobar]) EncodeToken(t Token[foobar]) {

}

type Token[foobar any] any
