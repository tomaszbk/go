// compile


// Issue 52841: gofrontend crashed writing export data

package p

func F() {
	x := ([17][1]interface {
		Method9()
		Method10()
	}{
		func() (V47 [1]interface {
			Method9()
			Method10()
		}) {
			return
		}(),
		func(V48 string) (V49 [1]interface {
			Method9()
			Method10()
		}) {
			return
		}("440"),
	})
	_ = x
}
