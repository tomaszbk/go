package conditional

func value(flag bool, a, b func() int) int {
	// want +1 "if/else value choice can use a conditional expression"
	if flag {
		return a()
	} else {
		return b()
	}
}

func assign(flag bool) (result int) {
	// want +1 "if/else value choice can use a conditional expression"
	if flag {
		result = 1
	} else {
		result = 2
	}
	return
}

func boxed(flag bool) any {
	// want +1 "if/else value choice can use a conditional expression"
	if flag {
		return int(1)
	} else {
		return string("other")
	}
}

func boxedNil(flag bool, p *int) any {
	// want +1 "if/else value choice can use a conditional expression"
	if flag {
		return p
	} else {
		return nil
	}
}

func sameConstantType(flag bool) any {
	// want +1 "if/else value choice can use a conditional expression"
	if flag {
		return 1
	} else {
		return 2
	}
}

func nested(flag bool) (func() int, error) {
	return func() int {
		// want +1 "if/else value choice can use a conditional expression"
		if flag {
			return 3
		} else {
			return 4
		}
	}, nil
}

func lambda(flag bool) func() int {
	return () => {
		// want +1 "if/else value choice can use a conditional expression"
		if flag {
			return 5
		} else {
			return 6
		}
	}
}
