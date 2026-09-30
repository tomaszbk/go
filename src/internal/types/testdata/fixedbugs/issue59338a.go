// -lang=go1.20


package p

func g[P any](P)      {}
func h[P, Q any](P) Q { panic(0) }

var _ func(int) = g /* ERROR "implicitly instantiated function in assignment requires go1.21 or later" */
var _ func(int) string = h[ /* ERROR "partially instantiated function in assignment requires go1.21 or later" */ int]

func f1(func(int))      {}
func f2(int, func(int)) {}

func _() {
	f1(g /* ERROR "implicitly instantiated function as argument requires go1.21 or later" */)
	f2(0, g /* ERROR "implicitly instantiated function as argument requires go1.21 or later" */)
}
