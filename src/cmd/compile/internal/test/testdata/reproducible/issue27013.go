package p

func A(arg interface{}) {
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
	_ = arg.(interface{ Func() int32 })
}
