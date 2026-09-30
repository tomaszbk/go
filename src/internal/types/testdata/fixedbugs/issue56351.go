// -lang=go1.20


package p

func _(s []int) {
	clear /* ERROR "clear requires go1.21 or later" */ (s)
}
