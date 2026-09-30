// errorcheck -0 -d=ssa/check_bce/debug=1


package x

func Found(x []string) string {
	switch len(x) {
	default:
		return x[0]
	case 0, 1:
		return ""
	}
}

func NotFound(x []string) string {
	switch len(x) {
	default:
		return x[0]
	case 0:
		return ""
	case 1:
		return ""
	}
}
