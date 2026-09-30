package p

type S struct{}

func f[P S]() {}

var _ = f[S]
