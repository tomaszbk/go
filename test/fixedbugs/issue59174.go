// compile

package p

func p() {
	s := make([]int, copy([]byte{' '}, "")-1)
	_ = append([]int{}, make([]int, len(s))...)
}
