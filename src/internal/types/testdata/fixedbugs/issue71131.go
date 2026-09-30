package p

func _() {
	type Bool bool
	for range func /* ERROR "yield func returns user-defined boolean, not bool" */ (func() Bool) {} {
	}
	for range func /* ERROR "yield func returns user-defined boolean, not bool" */ (func(int) Bool) {} {
	}
	for range func /* ERROR "yield func returns user-defined boolean, not bool" */ (func(int, string) Bool) {} {
	}
}
