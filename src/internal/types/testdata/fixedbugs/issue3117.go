package p

type S struct {
	a [1]int
}

func _(m map[int]S, key int) {
	m /* ERROR "cannot assign to m[key].a[0] (neither addressable nor a map index expression)" */ [key].a[0] = 0
}
