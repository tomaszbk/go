package p

func f(...any)

func _(x int, s []int) {
	f(0, x /* ERROR "have (number, int...)\n\twant (...any)" */ ...)
	f(0, s /* ERROR "have (number, []int...)\n\twant (...any)" */ ...)
	f(0, 0 /* ERROR "have (number, number...)\n\twant (...any)" */ ...)
}
