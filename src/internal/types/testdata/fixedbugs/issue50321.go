package p

func Ln[A A /* ERROR "cannot use a type parameter as constraint" */ ](p A) {
}
