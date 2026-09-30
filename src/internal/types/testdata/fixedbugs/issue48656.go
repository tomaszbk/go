package p

func f[P *Q, Q any](P, Q) {
	_ = f[P]
}

func f2[P /* ERROR "instantiation cycle" */ *Q, Q any](P, Q) {
	_ = f2[*P]
}
