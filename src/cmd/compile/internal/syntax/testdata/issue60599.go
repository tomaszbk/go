package p

func _(x, y, z int) {
	if x /* ERROR cannot use assignment x = y as value */ = y {}
	if x || y /* ERROR cannot use assignment \(x || y\) = z as value */ = z {}
	if x /* ERROR cannot use assignment x = \(y || z\) as value */ = y || z {}
}
