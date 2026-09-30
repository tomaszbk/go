// Mutually recursive type definitions imported and used by recursive1.go.

package p

type I1 interface {
	F() I2
}

type I2 interface {
	I1
}
