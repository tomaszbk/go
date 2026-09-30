package decl

func a() {
	var b, c bool // want `declared (and|but) not used`
	panic(c)

	if 1 == 1 {
		var s string // want `declared (and|but) not used`
	}
}

func b() {
	// b is a variable
	var b bool // want `declared (and|but) not used`
}

func c() {
	var (
		d string

		// some comment for c
		c bool // want `declared (and|but) not used`
	)

	panic(d)
}
