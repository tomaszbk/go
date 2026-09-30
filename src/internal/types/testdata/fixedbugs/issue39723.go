package p

// A constraint must be an interface; it cannot
// be a type parameter, for instance.
func _[A interface{ ~int }, B A /* ERROR "cannot use a type parameter as constraint" */ ]() {}
