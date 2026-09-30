// errorcheck

package main

func f[T interface{ ~[]P }, P any](t T) { // ERROR "instantiation cycle"
	if t == nil {
		return
	}
	f[[]T, T]([]T{t})
}

func main() {
	f[[]int](nil)
}
