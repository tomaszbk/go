// compile

package p

type nat []int

var a, b nat = y()

func y() (nat, []int) {
	return nat{0}, nat{1}
}
