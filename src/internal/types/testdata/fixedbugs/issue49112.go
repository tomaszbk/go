package p

func f[P int](P) {}

func _() {
	_ = f[int]
	_ = f[[ /* ERROR "[]int does not satisfy int ([]int missing in int)" */ ]int]

	f(0)
	f /* ERROR "P (type []int) does not satisfy int" */ ([]int{}) // TODO(gri) better error message
}
