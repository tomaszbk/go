package stringintconv

func _(i int) {
	_ = string(i) // ERROR "conversion from int to string yields a string of one rune"
}
