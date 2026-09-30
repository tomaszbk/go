package issue62640

type E struct{}

// F should be hidden within S because of the S.F field.
func (E) F() {}

type S struct {
	E
	F int
}
