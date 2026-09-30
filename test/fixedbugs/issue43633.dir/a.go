package a

func F() bool {
	{
		x := false
		_ = x
	}
	if false {
		_ = func(x bool) {}
	}
	x := true
	return x
}

func G() func() bool {
	x := true
	return func() bool {
		{
			x := false
			_ = x
		}
		return x
	}
}
