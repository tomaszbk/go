package p

func _() {
	var x = new(T)
	f[x /* ERROR "not a type" */ /* ERROR "use of .(type) outside type switch" */ .(type)]()
}

type T struct{}

func f[_ any]() {}
