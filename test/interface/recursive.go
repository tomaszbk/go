// compile


// Check mutually recursive interfaces

package recursive

type I1 interface {
	foo() I2
}

type I2 interface {
	bar() I1
}

type T int
func (t T) foo() I2 { return t }
func (t T) bar() I1 { return t }
