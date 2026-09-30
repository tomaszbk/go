package a

func Assert(msgAndArgs ...any) {
}

func Run() int {
	Assert("%v")
	return 0
}

func Run2() int {
	return Run()
}
