package p

func Clip[S ~[]E, E any](s S) S {
	return s
}

var versions func()
var _ = Clip /* ERROR "in call to Clip, S (type func()) does not satisfy ~[]E" */ (versions)
