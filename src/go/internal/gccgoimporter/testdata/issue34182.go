package issue34182

type T1 struct {
	f *T2
}

type T2 struct {
	f T3
}

type T3 struct {
	*T2
}
