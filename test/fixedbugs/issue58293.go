// compile

package p

var bar = f(13579)

func f(x uint16) uint16 {
	return x>>8 | x<<8
}
