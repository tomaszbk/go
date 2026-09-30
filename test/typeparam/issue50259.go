// compile

package p

var x T[B]

type T[_ any] struct{}
type A T[B]
type B = T[A]
