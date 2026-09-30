// compile

package p

type Any any
type IntOrBool interface{ int | bool }

type I interface{ Any | IntOrBool }

var (
	X I = 42
	Y I = "xxx"
	Z I = true
)

type A interface{ *B | int }
type B interface{ A | any }
