package a

func F1() {
L:
	goto L
}

func F2() {
L:
	for {
		break L
	}
}

func F3() {
L:
	for {
		continue L
	}
}

func F4() {
	switch {
	case true:
		fallthrough
	default:
	}
}

type T struct{}

func (T) M1() {
L:
	goto L
}

func (T) M2() {
L:
	for {
		break L
	}
}

func (T) M3() {
L:
	for {
		continue L
	}
}

func (T) M4() {
	switch {
	case true:
		fallthrough
	default:
	}
}
