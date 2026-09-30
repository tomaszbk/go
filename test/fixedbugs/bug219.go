// compile


package bug219

func f(func()) int { return 0 }

// this doesn't work:
// bug219.go:16: syntax error near if
func g1() {
	if x := f(func() {
		if true {}
	}); true {
		_ = x;
	}
}

// this works
func g2() {
	if x := f(func() {
		//if true {}
	}); true {
		_ = x;
	}
}

// this works
func g3() {
	x := f(func() {
		if true {}
	});
	if true {
		_ = x;
	}
}
