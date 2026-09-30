// compile

package p

func f() {
	var x string
	func() [10][]bool {
		return [10][]bool{
			[]bool{bool(x < "")},
			[]bool{}, []bool{}, []bool{}, []bool{}}
	}()
}
