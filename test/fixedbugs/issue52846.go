// compile


// Issue 52846: gofrontend crashed with alias as map key type

package p

type S struct {
	F string
}

type A = S

var M = map[A]int{A{""}: 0}
