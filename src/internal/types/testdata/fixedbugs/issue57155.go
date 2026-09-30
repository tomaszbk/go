package p

func f[P *Q, Q any](p P, q Q) {
	func() {
		_ = f[P]
		f(p, q)
		f[P](p, q)
		f[P, Q](p, q)
	}()
}
