package conditional

var global int

func reject(flag bool, values []int, index func() int) {
	if flag {
		values[index()] = 1
	} else {
		values[index()] = 2
	}
	if flag {
		global = 1
	} else {
		global = 2
	}
}

func comments(flag bool) int {
	if flag {
		// Preserve the user's explanation.
		return 1
	} else {
		return 2
	}
}

func initStatement(flag bool) int {
	if local := flag; local {
		return 1
	} else {
		return 2
	}
}

func multiple(flag bool) (int, error) {
	if flag {
		return 1, nil
	} else {
		return 2, nil
	}
}

func tuple(flag bool, f func() (int, int)) (int, int) {
	if flag {
		return f()
	} else {
		return f()
	}
}

func directNested(flag bool) int {
	if flag {
		return if flag { 1 } else { 2 }
	} else {
		return 3
	}
}

func preferCoalesce(p, fallback *int) *int {
	if p != nil {
		return p
	} else {
		return fallback
	}
}

func preferGuard(p *struct{ N int }) int {
	if p != nil {
		return p.N
	} else {
		return 0
	}
}

func mixedConstants(flag bool) any {
	if flag {
		return 1
	} else {
		return "other"
	}
}

func distinctDefaultTypes(flag bool) any {
	if flag {
		return 1
	} else {
		return 'a'
	}
}

func inferredBoxType(flag bool) any {
	if flag {
		return int8(1)
	} else {
		return 2
	}
}
