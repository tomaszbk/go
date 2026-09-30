// compile

// Make sure that the compiler can analyze non-reflect
// Type.{Method,MethodByName} calls.

package p

type I interface {
	MethodByName(string)
	Method(int)
}

type M struct{}

func (M) MethodByName(string) {}
func (M) Method(int)          {}

func f() {
	var m M
	I.MethodByName(m, "")
	I.Method(m, 42)
}
