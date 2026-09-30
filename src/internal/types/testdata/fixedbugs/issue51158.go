package p

// Type checking the following code should not cause an infinite recursion.
func f[M map[K]int, K comparable](m M) {
        f(m)
}

// Equivalent code using mutual recursion.
func f1[M map[K]int, K comparable](m M) {
        f2(m)
}
func f2[M map[K]int, K comparable](m M) {
        f1(m)
}
