package p

func f[P ~*T, T any]() {
	var p P
	var tp *T
	tp = p // this assignment is valid
	_ = tp
}
