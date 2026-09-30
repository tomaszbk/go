// errorcheck

package p

type T struct {
	X int
}

func (t *T) X() {} // ERROR "type T has both field and method named X|redeclares struct field name|field and method with the same name"
