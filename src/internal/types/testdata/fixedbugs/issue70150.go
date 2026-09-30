package p

func _() {
	var values []int
	vf(values /* ERROR "(variable of type []int) as string value" */)
	vf(values...) /* ERROR "have ([]int...)\n\twant (string, ...int)" */
	vf("ab", "cd", values /* ERROR "have (string, string, []int...)\n\twant (string, ...int)" */ ...)
}

func vf(method string, values ...int) {
}
