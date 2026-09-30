package a

func A() int {
	return p("count")
}

func p(which string, args ...string) int {
	switch which {
	case "count", "something":
		return 1
	default:
		return 2
	}
}
