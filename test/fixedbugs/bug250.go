// compile

package bug250

type I1 interface {
	m() I2
}

type I2 interface {
	I1
}

var i1 I1 = i2
var i2 I2
var i2a I2 = i1
